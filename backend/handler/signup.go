package handler

import (
	"log"
	"net/http"

	"be-react-login/handler/model"
	"be-react-login/handler/model/apperrors"

	"github.com/gin-gonic/gin"
)

type SignUpReq struct {
	Username string `json:"username" binding:"required,gte=4,lte=24"`
	Password string `json:"password" binding:"required,gte=8,lte=24"`
}

func (h *Handler) Signup(c *gin.Context) {
	var req SignUpReq

	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Username: req.Username,
		Password: req.Password,
	}
	ctx := c.Request.Context()
	err := h.UserService.Signup(ctx, u)
	if err != nil {
		log.Printf("failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := h.TokenService.NewTokenFromUser(ctx, u, "")
	if err != nil {
		log.Printf("failed to create token for user: %v\n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
