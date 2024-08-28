package handlers

import (
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/user"
	"github.com/gin-gonic/gin"
)

func (h *HandlerST) Register(ctx *gin.Context) {
	var req user.CreateUserReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	resp, err := h.ClientRepository.Register(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) Login(ctx *gin.Context) {
	var req user.LoginReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	resp, err := h.ClientRepository.Login(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) Verify(ctx *gin.Context) {
	var req user.VerifyReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	resp, err := h.ClientRepository.Verify(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	var req user.UpdateUserReq
	req.UserId = userId
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	resp, err := h.ClientRepository.UpdateUser(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	req := &user.DeleteUserReq{
		Id: userId,
	}

	resp, err := h.ClientRepository.DeleteUser(ctx, req)
	if err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	req := &user.GetUserByIdReq{
		Id: userId,
	}

	resp, err := h.ClientRepository.GetUserById(ctx, req)
	if err != nil {
		ctx.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}
