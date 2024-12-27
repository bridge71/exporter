package server

import (
	"exporter/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)
	r.POST("/create/user", s.CreaterUserHandler)
	r.POST("/auth", s.LoginHandler)

	return r
}

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *Server) CreaterUserHandler(c *gin.Context) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of user",
		})
		return
	}
	fmt.Printf("%v\n", *user)

	isLong, message := CheckStringLen(*user, false)
	if isLong {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: message,
		})
		return
	}
	user.PasswordHash, err = EncryptPassword(user.PasswordHash)

	err = s.db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "email may be illegal",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "register successfully",
		})
	}
}

func CheckStringLen(user models.User, flag bool) (bool, string) {
	if len(user.PasswordHash) > 36 || len(user.PasswordHash) == 0 {
		return true, "password length is illegal"
	}
	if len(user.Email) > 36 || len(user.Email) == 0 {
		return true, "email length is illegal"
	}
	if flag {
		return false, ""
	}
	if len(user.UserName) > 57 || len(user.UserName) == 0 {
		return true, "userName length is illegal"
	}
	return false, ""
}

func (s *Server) LoginHandler(c *gin.Context) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of user",
		})
		return
	}
	isLong, message := CheckStringLen(*user, true)
	if isLong {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: message,
		})
		return
	}

	auth := &models.User{}
	s.db.FindUser(auth, user.Email)
	err = bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to login",
		})
	}
	c.JSON(http.StatusOK, models.Message{
		RetMessage: "login successfully",
	})
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
