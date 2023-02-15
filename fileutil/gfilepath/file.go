package gfilepath

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func MkAllDir(dirPath string) error {
	err := os.MkdirAll(dirPath, 0666)
	if err != nil {
		return err
	}
	return nil
}

func MkFile(filePath string) {
	_, err := os.Create(filePath)
	if err != nil {
		return
	}
}

func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func ListDir(dirPath string) []string {
	files, _ := ioutil.ReadDir(dirPath)
	fileOrDirList := make([]string, len(files))
	for _, f := range files {
		fileName := f.Name()
		if !strings.HasPrefix(fileName, ".") {
			filePath := filepath.Join(dirPath, fileName)
			fileOrDirList = append(fileOrDirList, filePath)
		}
		fmt.Println(f.Name())
	}
	return fileOrDirList
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func GetFileSize(fpath string) (int64, error) {
	fi, err := os.Stat(fpath)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}
