package helpers

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SingleFileUpload(c *gin.Context, field_name string, path string, file_type []string, size int, require bool) (string, error) {
	file, err := c.FormFile(field_name)

	if err != nil {
		if require == true {
			return "", errors.New("field " + field_name + " is empty")
		} else {
			return "", nil
		}
	}

	if (file.Size / 1000) > int64(size) {
		return "", errors.New("file size to big")
	}
	if file_type != nil {
		if !checkFileType(file.Header["Content-Type"][0], file_type) {
			return "", errors.New("file type not allowed")
		}
	}

	filename := strconv.Itoa(rand.Intn(999999-111111+1)+111111) + "_" + file.Filename

	if err != nil {
		return "", err
	}

	if err := c.SaveUploadedFile(file, path+filename); err != nil {
		return "", err
	}

	return filename, nil
}

func checkFileType(current_type string, file_type []string) bool {
	for _, ft := range file_type {
		if current_type == ft {
			return true
		}
	}
	return false
}
