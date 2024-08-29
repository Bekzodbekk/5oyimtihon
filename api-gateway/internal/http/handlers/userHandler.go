package handlers

import (
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/user"
	"github.com/gin-gonic/gin"
)

// @Summary Register a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserReq true "User registration information"
// @Success 200 {object} CreateUserResp
// @Failure 400 {object} Error
// @Router /users/register [post]
func (h *HandlerST) Register(ctx *gin.Context) {
	var req user.CreateUserReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.ClientRepository.Register(ctx, &req)
	if err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @Summary User login
// @Description Authenticate a user and return a token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body LoginReq true "User login credentials"
// @Success 200 {object} LoginResp
// @Failure 400 {object} Error
// @Router /users/login [post]
func (h *HandlerST) Login(ctx *gin.Context) {
	var req user.LoginReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.ClientRepository.Login(ctx, &req)
	if err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @Summary Verify user
// @Description Verify a user's account
// @Tags users
// @Accept json
// @Produce json
// @Param verificationData body VerifyReq true "User verification data"
// @Success 200 {object} VerifyResp
// @Failure 400 {object} Error
// @Router /users/verify [post]
func (h *HandlerST) Verify(ctx *gin.Context) {
	var req user.VerifyReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.ClientRepository.Verify(ctx, &req)
	if err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @Summary Update user
// @Description Update an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param user body UpdateUserReq true "Updated user information"
// @Success 200 {object} UpdateUserResp
// @Failure 400 {object} Error
// @Router /users/{user_id} [put]
func (h *HandlerST) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	var req user.UpdateUserReq
	req.UserId = userId
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.ClientRepository.UpdateUser(ctx, &req)
	if err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @Summary Delete user
// @Description Delete an existing user
// @Tags users
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} DeleteUserResp
// @Failure 400 {object} Error
// @Router /users/{user_id} [delete]
func (h *HandlerST) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	req := &user.DeleteUserReq{
		Id: userId,
	}

	resp, err := h.ClientRepository.DeleteUser(ctx, req)
	if err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @Summary Get user by ID
// @Description Retrieve a user's information by their ID
// @Tags users
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} GetUserByIdResp
// @Failure 400 {object} Error
// @Router /users/{user_id} [get]
func (h *HandlerST) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	req := &user.GetUserByIdReq{
		Id: userId,
	}

	resp, err := h.ClientRepository.GetUserById(ctx, req)
	if err != nil {
		ctx.JSON(400, Error{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}
