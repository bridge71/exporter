package server

import (
	"exporter/internal/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *Server) CreateUserHandler(c *gin.Context) {
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

	user.PasswordHash = ""
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "email may be illegal",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "register successfully",
			User:       *user,
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

func (s *Server) ModifyUserPassHandler(c *gin.Context) {
	user := &models.User{}
	idStr, err := c.Cookie("UserId")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to get cookie of UserId ",
		})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to transform string to number",
		})
		return
	}
	user.UserId = uint(id)

	password := c.PostForm("password")
	log.Printf("sss   %u\n", user.UserId)

	passwordHash, _ := EncryptPassword(password)
	s.db.ModifyUserPass(user, passwordHash)
}

func (s *Server) LoginHandler(c *gin.Context) {
	log.Printf("sss\n")
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	log.Printf("%v ss\n", *user)
	if err != nil {
		log.Printf("error in bind of user\n")
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
	auth.PasswordHash = ""
	c.JSON(http.StatusOK, models.Message{
		RetMessage: "login successfully",
		User:       *auth,
	})
}
