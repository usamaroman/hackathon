package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/usamaroman/hackathon/pkg/jwt"
)

var WrongEnteredPasswordError = errors.New("неверно веден пароль")

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.Handle(http.MethodPost, "/auth/registration", h.Registration)
	router.Handle(http.MethodPost, "/auth/login", h.Login)
	router.Handle(http.MethodGet, "/auth/logout", h.Logout)
	router.Handle(http.MethodPost, "/auth/refresh", h.RefreshToken)
}

func (h *handler) Registration(ctx *gin.Context) {
	var body RegistrationDto
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println("error during binding json", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ValidateForEmptyPassword(body.Password)
	if err != nil {
		log.Println("error during validating password and full name", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Registration(ctx, body)
	if err != nil {
		log.Println("error during registration process", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "успешная регистрация"})
}

func (h *handler) Login(ctx *gin.Context) {
	var body LoginDto
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("wrong entered data").Error()})
		return
	}

	user, token, err := h.service.Login(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res LoginResDto
	res.User = user
	res.AccessToken = token

	refreshToken, err := jwt.GenerateRefreshToken(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	res.RefreshToken = refreshToken

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) RefreshToken(ctx *gin.Context) {
	var reqDto struct {
		RefreshToken string `json:"refresh_token"`
	}

	err := ctx.ShouldBindJSON(&reqDto)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "не авторизован",
		})
		return
	}
	log.Println(reqDto)

	if reqDto.RefreshToken == "" {
		fmt.Println("empty")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "не авторизован",
		})
		return
	}

	token, err := jwt.ParseRefreshTokenToken(reqDto.RefreshToken)
	if err != nil {
		fmt.Println("hz")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "не авторизован",
		})
		return
	}

	id := token["id"]
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	u, err := h.service.repository.GetOneUserById(ctx, fmt.Sprintf("%s", id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	role, err := h.service.repository.GetRole(ctx, u.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	generateAccessToken, err := jwt.GenerateAccessToken(u.Id, u.Email, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	generateRefreshToken, err := jwt.GenerateRefreshToken(u.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  generateAccessToken,
		"refresh_token": generateRefreshToken,
	})
}

func (h *handler) Logout(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "127.0.01", false, false)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "log out",
	})
}

func ValidateForEmptyPassword(password string) error {
	if strings.Contains(password, " ") {
		return WrongEnteredPasswordError
	}

	password = strings.Trim(password, " ")

	return nil
}
