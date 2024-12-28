package server

import (
	"crypto/md5"
	"encoding/hex"
	"exporter/internal/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

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
	r.GET("/find/acct", s.FindAcctHandler)
	r.POST("/create/user", s.CreateUserHandler)
	r.POST("/auth", s.LoginHandler)
	r.POST("/create/acct", s.CreateAcctHandler)

	return r
}

func (s *Server) FindAcctHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	s.db.FindAcct(accts)
	c.JSON(http.StatusOK, models.Message{
		Acct: *accts,
	})
}

func (s *Server) CreateAcctHandler(c *gin.Context) {
	acct := &models.Acct{}
	err := c.ShouldBind(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of acct",
		})
		return
	}

	bankIndex := 0
	for {
		accNum := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].AccNum")
		accName := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].AccName")
		isUpload := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].IsUpload")
		currency := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].Currency")
		bankName := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].BankName")
		bankNum := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].BankNum")
		swiftCode := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].SwiftCode")
		bankAddr := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].BankAddr")
		notes := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].Notes")

		flag := false
		if isUpload == "true" {
			flag = true
		}

		if accNum == "" || accName == "" || isUpload == "" {
			break
		}
		acct.AcctBanks = append(acct.AcctBanks, models.AcctBank{
			AccName:   accName,
			AccNum:    accNum,
			IsUpload:  flag,
			Currency:  currency,
			BankName:  bankName,
			BankNum:   bankNum,
			SwiftCode: swiftCode,
			BankAddr:  bankAddr,
			Notes:     notes,
		})
		bankIndex++
	}

	now := 0
	allFiles, err := c.MultipartForm()
	attaches := allFiles.File["attach[]"]
	if err == nil {
		for index, attach := range attaches {
			log.Printf("%d index\n", index)
			src, err := attach.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Message{
					RetMessage: "error when opening file",
				})
				return
			}
			defer src.Close()

			hash := md5.New()
			if _, err := io.Copy(hash, src); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "文件读取失败",
					"details": err.Error(),
				})
				return
			}
			MD5 := hex.EncodeToString(hash.Sum(nil))
			FileName := attach.Filename
			Suffix := filepath.Ext(attach.Filename)
			File := &models.File{
				Name:   FileName,
				MD5:    MD5,
				Suffix: Suffix,
			}
			s.db.CreateFile(File)
			for ; now < bankIndex; now++ {
				if acct.AcctBanks[now].IsUpload {
					acct.AcctBanks[now].FileId = File.FileId
					now++
					break
				}
			}
			if now >= bankIndex && acct.IsUpload {
				acct.FileId = File.FileId
			}

			c.SaveUploadedFile(attach, "../files/"+MD5+Suffix) // 以main程序所在目录为基准

			log.Printf(attach.Filename + "\n")
		}
	}

	log.Printf("%v", acct)
	err = s.db.CreateAcct(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error when insert accounting information",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "insert accounting information successfully",
		})
	}
}

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
	auth.PasswordHash = ""
	c.JSON(http.StatusOK, models.Message{
		RetMessage: "login successfully",
		User:       *auth,
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
