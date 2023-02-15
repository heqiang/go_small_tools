package gjsonfile

import (
	"errors"
	"go_small_tools/fileutil/gfilepath"
	"io"
	"io/ioutil"
	"os"
)

var (
	FileSize    int64 = 2 << 12
	FileNoExist       = errors.New("file is no exist,please check the file path")
)

func ReadJsonFile(jsonFile string) ([]byte, error) {
	if gfilepath.IsFileExist(jsonFile) {
		fileSize, err := gfilepath.GetFileSize(jsonFile)
		if err != nil {
			return nil, err
		}
		var chunk []byte

		f, err := os.OpenFile(jsonFile, os.O_RDONLY, 0666)
		if err != nil {
			return nil, err
		}
		//把file读取到缓冲区中
		defer f.Close()

		if fileSize > FileSize {
			buf := make([]byte, 1024)
			for {
				//从file读取到buf中
				n, err := f.Read(buf)
				if err != nil && err != io.EOF {
					return nil, err
				}
				//说明读取结束
				if n == 0 {
					break
				}
				//读取到最终的缓冲区中
				chunk = append(chunk, buf[:n]...)
			}

		} else {
			return ioutil.ReadFile(jsonFile)
		}
	}

	return nil, FileNoExist

}
