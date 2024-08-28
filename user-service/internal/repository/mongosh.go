package repository

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/pkg/load"
	"user-service/token"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/user"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
	"gopkg.in/gomail.v2"
)

type UserRepo struct {
	coll *mongo.Collection
	rdb  *redis.Client
	conf load.Config
}

func NewUserRepo(coll *mongo.Collection, rdb *redis.Client, conf load.Config) UserRepository {
	return &UserRepo{
		coll: coll,
		rdb:  rdb,
		conf: conf,
	}
}

// Create (Register)
func (u *UserRepo) Register(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	expiredPassword := RandomNumber()

	userData := map[string]interface{}{
		"email":           req.Email,
		"username":        req.Username,
		"password":        passwordHash,
		"expiredPassword": expiredPassword,
	}
	err = u.rdb.HSet(ctx, req.Email, userData).Err()
	if err != nil {
		return nil, err
	}
	err = u.rdb.Expire(ctx, req.Email, 60*time.Second).Err()
	if err != nil {
		return nil, err
	}

	err = u.SendEmail(req.Email, expiredPassword)
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResp{
		Status:  true,
		Message: "Data saved in redis!",
	}, nil
}

// Read (GetUserById)
func (u *UserRepo) GetUserById(ctx context.Context, req *user.GetUserByIdReq) (*user.GetUserByIdResp, error) {
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	var userDoc bson.M
	err = u.coll.FindOne(ctx, bson.M{"_id": objectID, "deleted_at": 0}).Decode(&userDoc)
	if err != nil {
		return nil, err
	}

	var deletedAt int64
	if v, ok := userDoc["deleted_at"].(int32); ok {
		deletedAt = int64(v)
	} else if v, ok := userDoc["deleted_at"].(int64); ok {
		deletedAt = v
	}

	userResp := &user.User{
		UserId:   req.Id,
		Email:    userDoc["email"].(string),
		Username: userDoc["username"].(string),
		Cud: &user.CUDUser{
			CreatedAt: userDoc["created_at"].(string),
			UpdatedAt: userDoc["updated_at"].(string),
			DeletedAt: deletedAt,
		},
	}

	return &user.GetUserByIdResp{
		Status:  true,
		Message: "User found",
		User:    userResp,
	}, nil
}

// Update
func (u *UserRepo) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	update := bson.M{
		"$set": bson.M{
			"email":      req.Email,
			"username":   req.Username,
			"updated_at": time.Now().String(),
		},
	}

	if req.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		update["$set"].(bson.M)["password"] = string(passwordHash)
	}

	_, err := u.coll.UpdateOne(ctx, bson.M{"email": req.Email}, update)
	if err != nil {
		return nil, err
	}

	return &user.UpdateUserResp{
		Status:  true,
		Message: "User updated successfully",
	}, nil
}

// Delete
func (u *UserRepo) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (*user.DeleteUserResp, error) {
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().Unix(),
		},
	}

	_, err = u.coll.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return &user.DeleteUserResp{
		Status:  true,
		Message: "User deleted successfully",
	}, nil
}

// Login
func (u *UserRepo) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	conf, err := load.LOAD("./config/config.yaml")
	if err != nil {
		return nil, err
	}

	var userDoc bson.M
	err = u.coll.FindOne(ctx, bson.M{"username": req.Username}).Decode(&userDoc)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDoc["password"].(string)), []byte(req.Password))
	if err != nil {
		return &user.LoginResp{
			Status:  false,
			Message: "Invalid credentials",
		}, nil
	}
	idObject := userDoc["_id"].(primitive.ObjectID)
	id := idObject.Hex()

	token, err := token.GenerateToken(*conf, id)
	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		Status:  true,
		Message: "Login successful",
		Token:   token,
	}, nil
}

// Verify
func (u *UserRepo) Verify(ctx context.Context, req *user.VerifyReq) (*user.VerifyResp, error) {
	expPass, err := u.rdb.HGet(ctx, req.Email, "expiredPassword").Result()
	if err != nil {
		return nil, err
	}
	if expPass != req.Password {
		return &user.VerifyResp{
			Status:  false,
			Message: "Password Wrong",
		}, nil
	}

	userData, err := u.rdb.HGetAll(ctx, req.Email).Result()
	if err != nil {
		return nil, err
	}

	userInsert := bson.M{
		"email":      userData["email"],
		"username":   userData["username"],
		"password":   userData["password"],
		"created_at": time.Now().String(),
		"updated_at": time.Now().String(),
		"deleted_at": 0,
	}
	_, err = u.coll.InsertOne(ctx, userInsert)
	if err != nil {
		return nil, err
	}

	return &user.VerifyResp{
		Status:  true,
		Message: "User Created Successfully",
	}, nil
}

func RandomNumber() int {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Intn(9000) + 1000
}

func (u *UserRepo) SendEmail(to string, code int) error {
	subject := "----Welcome buddy----"

	body := fmt.Sprintf(`
	  <!DOCTYPE html>
	  <html>
	  <head>
		<style>
		  .container {
			font-family: Arial, sans-serif;
			background-color: #f4f4f4;
			padding: 20px;
			border-radius: 10px;
			width: 80%%;
			margin: 0 auto;
			color: #333;
		  }
		  .header {
			background-color: #4CAF50;
			color: white;
			padding: 10px;
			border-radius: 10px 10px 0 0;
			text-align: center;
		  }
		  .content {
			padding: 20px;
			background-color: white;
			border-radius: 0 0 10px 10px;
		  }
		  .code {
			font-size: 24px;
			font-weight: bold;
			color: #4CAF50;
			text-align: center;
			margin: 20px 0;
		  }
		  .footer {
			text-align: center;
			padding: 10px;
			font-size: 12px;
			color: #777;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <div class="header">
			<h1>Welcome to Our Service!</h1>
		  </div>
		  <div class="content">
			<p>Dear user,</p>
			<p>Thank you for signing up. To complete your registration, please use the following confirmation code:</p>
			<div class="code">%d</div>
			<p>If you didn't sign up, please ignore this email.</p>
		  </div>
		  <div class="footer">
			<p>This is an automated message, please do not reply.</p>
		  </div>
		</div>
	  </body>
	  </html>
	`, code)

	m := gomail.NewMessage()
	m.SetHeader("From", "bekzodnematov709@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	fmt.Println(u.conf.Email)
	fmt.Println(u.conf.PassKey)
	d := gomail.NewDialer("smtp.gmail.com", 587, u.conf.Email, u.conf.PassKey)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
