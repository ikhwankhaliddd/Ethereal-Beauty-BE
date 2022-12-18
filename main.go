package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"project_dwi/auth"
	"project_dwi/handler"
	"project_dwi/helper"
	"project_dwi/payment"
	"project_dwi/products"
	"project_dwi/transactions"
	"project_dwi/users"
	"strings"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	usersRepository := users.NewRepository(db)
	productRepository := products.NewRepository(db)
	transactionRepository := transactions.NewRepository(db)

	userService := users.NewService(usersRepository)
	productService := products.NewService(productRepository)
	paymentService := payment.NewService()
	transactionService := transactions.NewService(transactionRepository, productRepository, paymentService)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	productHandler := handler.NewProductHandler(productService)
	transactionHandler := handler.NewTransactionsHandler(transactionService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/checkEmail", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	api.GET("/products", productHandler.GetProducts)
	api.GET("/products/:id", productHandler.GetProduct)
	api.POST("/products", authMiddleware(authService, userService), productHandler.CreateProduct)
	api.PUT("/products/:id", authMiddleware(authService, userService), productHandler.UpdateProduct)
	api.POST("/product-images", authMiddleware(authService, userService), productHandler.UploadProductImage)

	api.GET("products/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetTransactionsByProductID)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateUserTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)
	router.Run()
}

func authMiddleware(authService auth.Service, userService users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(payload["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
