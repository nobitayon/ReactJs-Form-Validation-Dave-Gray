package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"be-react-login/handler"
	"be-react-login/repository"
	"be-react-login/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func inject(d *dataSource) (*gin.Engine, error) {
	log.Println("injecting data sources")

	// repository layer
	userRepository := repository.NewUserRepository(d.DB)

	// service layer
	userService := service.NewUserService(&service.USConfig{
		UserRepository: userRepository,
	})

	secret := os.Getenv("SECRET")
	idTokenExp := os.Getenv("ID_TOKEN_EXP")
	idExp, err := strconv.ParseInt(idTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse ID_TOKEN_EXP as int: %w", err)
	}
	tokenService := service.NewTokenService(&service.TSConfig{
		Secret:           secret,
		IDExpirationSecs: idExp,
	})

	router := gin.Default()
	str_allowed_origins := os.Getenv("ALLOWED_ORIGINS")
	list_allowed_origins := strings.Split(str_allowed_origins, ",")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     list_allowed_origins,
		AllowHeaders:     []string{"X-Requested-with, Content-Type, Authorization, Access-Control-Allow-Origin"},
		AllowMethods:     []string{"POST, OPTIONS, GET, PUT, DELETE, OPTIONS"},
		AllowCredentials: true,
	}))
	baseURL := os.Getenv("API_URL")
	handler.NewHandler(&handler.Config{
		R:            router,
		UserService:  userService,
		TokenService: tokenService,
		BaseURL:      baseURL,
	})

	return router, nil

}
