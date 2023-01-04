package helpers

import (
	"os"
)

func GetFileSize(fileName string) (int64, error) {
	f, err := os.Stat(fileName)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}
