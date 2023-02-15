package gjsonfile

import (
	"github.com/tidwall/gjson"
	"io"
	"os"
)

func WriteToJson(srcJsonFile, dstJsonPath string) (int, error) {
	dstFile, err := os.OpenFile(dstJsonPath, os.O_RDONLY, 0666)
	defer dstFile.Close()
	if err != nil {
		return 0, nil
	}

	srcFileByte, err := ReadJsonFile(srcJsonFile)
	if err != nil {
		return 0, nil
	}
	writeCount := 0

	gjson.ForEachLine(string(srcFileByte), func(line gjson.Result) bool {
		wNum, err := io.WriteString(dstFile, line.String()+"\n")
		if err != nil {
			return false
		}
		writeCount += wNum
		return true
	})

	return writeCount, nil
}
