package storage

import (
	"fmt"
	"os"

	"github.com/datagenx/license-generator/internal/generate"
)

type FileStorage struct {
	FilePath string
}

func (fs *FileStorage) Save(sl generate.Slic, data string) error {

	fh, err := os.OpenFile(fs.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		return err
	}
	defer fh.Close()

	value := fmt.Sprintf("%s|%s|%s|%s|%s|%d|%d|%s|%s\n", sl.Id, sl.Customer,
		sl.ValidFrom, sl.Expiry, sl.HardExpiry, sl.Seats, sl.HardSeats, sl.Type, data)

	_, err = fh.WriteString(value)

	return err
}

func (fs *FileStorage) ReadAll() ([]string, error) {
	content, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return nil, err
	}
	return []string{string(content)}, nil
}
