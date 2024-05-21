package util

import (
	"apriori-backend/constant"
	"path/filepath"
	"strings"
	"time"
)

func GenerateImageName(nameUser string, filename string) string {
	ext := filepath.Ext(filename)
	currentTime := time.Now().UTC().Format("20060102150405.000000000")
	newCurrentTime := strings.ReplaceAll(currentTime, ".", "")
	return strings.ToLower(nameUser + "_" + newCurrentTime + ext)
}

func GetImageUrl(fileName string) string {
	return "https://alterra.sgp1.cdn.digitaloceanspaces.com/mini-project/" + fileName
}

func GetContentType(filename string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpeg":
		return "image/jpeg", nil
	case ".png":
		return "image/png", nil
	case ".jpg":
		return "image/jpg", nil
	default:
		return "", constant.ErrUnsupportedFileFormat
	}
}
