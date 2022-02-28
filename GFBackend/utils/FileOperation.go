package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

var DirBasePath = "./resources/userfiles/"

func IsDirExists(username string) bool {
	info, err := os.Stat(DirBasePath + username)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsFileExists(username, filename string) bool {
	info, err := os.Stat(DirBasePath + username + "/" + filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func CreateDir(username string) bool {
	err := os.Mkdir(DirBasePath+username, 755)
	if err != nil {
		return false
	}
	return true
}

func DeleteDir(username string) bool {
	err := os.RemoveAll(DirBasePath + username)
	if err != nil {
		return false
	}
	return true
}

func DeleteFile(username, filename string) bool {
	err := os.Remove(DirBasePath + username + "/" + filename)
	if err != nil {
		return false
	}
	return true
}

func DirSize(username string) (float64, error) {
	var size int64
	err := filepath.Walk(DirBasePath+username, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	float64Size := float64(size) / (1024 * 1024)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64Size), 64)
	return value, err
}

func GetFilesInDir(username string) ([]string, error) {
	return nil, nil
}
