package media

import (
	"encoding/base64"
	"os"
)

func LoadEncodedMediaString(filepath string) string {
	file, _ := os.Open(filepath)
	defer file.Close()

	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	return base64.RawStdEncoding.EncodeToString(data)
}
