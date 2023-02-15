package gjsonfile

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestReadJsonFile(t *testing.T) {
	rootPath, _ := os.Getwd()
	fmt.Println(rootPath)
	srcFilePath := filepath.Join(rootPath, "gjsonfile", "srcjson.json")
	readByte, err := ReadJsonFile(srcFilePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(readByte))
}
