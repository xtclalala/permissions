package utils

import (
	"mime/multipart"
	"net/http"
)

func GetFileType(out multipart.File) (fileType string, err error) {
	buffer := make([]byte, 512)
	if _, err = out.Read(buffer); err != nil {
		return "", err
	}
	fileType = http.DetectContentType(buffer)
	return
}
