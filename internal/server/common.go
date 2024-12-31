package server

import (
	"crypto/md5"
	"encoding/hex"
	"exporter/internal/models"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) Uploader(c *gin.Context) {
	file := &models.File{}
	err := c.ShouldBind(file)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of file",
		})
		return
	}
	s.db.FindFile(file)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename*=utf-8''"+file.Name+file.Suffix)
	c.File("../files/" + file.MD5 + file.Suffix)

	c.JSON(http.StatusOK, models.Message{
		RetMessage: file.Name,
	})
}

func (s *Server) SaveFile(c *gin.Context) (error, uint, string) {
	file, err := c.FormFile("file")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return err, 0, ""
		}
		defer src.Close()

		hash := md5.New()
		if _, err := io.Copy(hash, src); err != nil {
			return err, 0, ""
		}
		MD5 := hex.EncodeToString(hash.Sum(nil))
		FileName := file.Filename
		Suffix := filepath.Ext(file.Filename)
		File := &models.File{
			Name:   FileName,
			MD5:    MD5,
			Suffix: Suffix,
		}
		s.db.CreateFile(File)

		c.SaveUploadedFile(file, "../files/"+MD5+Suffix) // 以main文件所在目录为基准

		log.Printf(file.Filename + "\n")
		return nil, File.FileId, FileName
	}
	return nil, 0, ""
}

func (s *Server) Str2Uint(str string) uint {
	// 使用 strconv.ParseUint 将字符串解析为 uint64
	// 参数 10 表示十进制，参数 0 表示自动推断位数
	value, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		// 如果解析失败，返回错误
		return 0
	}
	// 将 uint64 转换为 uint 并返回
	return uint(value)
}
