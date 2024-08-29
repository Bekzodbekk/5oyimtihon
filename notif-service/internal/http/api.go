package http

import (
	"net/http"
	"notif-service/internal/repository"
	"time"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/notif"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	rdb  *redis.Client
	repo repository.NotifRepository
}

type NotifS struct {
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) GetNotification(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "WebSocket aloqasini ochishda xato",
			"error":   err.Error(),
		})
		return
	}
	defer conn.Close()

	email := ctx.Param("email")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			result, err := h.rdb.Get(ctx, email).Result()
			if err != nil {
				if err == redis.Nil {
					continue
				}
				conn.WriteMessage(websocket.TextMessage, []byte("Xabarlarni olishda xato: "+err.Error()))
				time.Sleep(5 * time.Second)
				continue
			}

			if err := conn.WriteJSON(NotifS{
				Message: result,
			}); err != nil {
				return
			}
			_, err = h.repo.Notification(ctx, &notif.NotificationReq{
				Status:  true,
				Message: result,
				Email:   email,
			})
			if err != nil {
				ctx.JSON(400, "Email yuborilmadi")
			}

			if err := h.rdb.Del(ctx, email).Err(); err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Redisdan ma'lumotni o'chirishda xato: "+err.Error()))
			}

			time.Sleep(1 * time.Second)
		}
	}

}

func NewGin(rdb *redis.Client) *gin.Engine {
	r := gin.Default()
	notifRepo := repository.NewNotifRepo()
	hnd := Handler{
		rdb:  rdb,
		repo: notifRepo,
	}
	r.GET("/notification/:email", hnd.GetNotification)
	return r
}
