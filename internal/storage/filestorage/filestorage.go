package filestorage

import (
	"bufio"
	"encoding/gob"
	"errors"
	"io"
	"os"
)

var ErrFileStorageNotClose = errors.New("file storage has not close")

// Write data to path
func Write(path string, data interface{}) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	// Handle for file close
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(ErrFileStorageNotClose)
		}
	}(f)

	// logger.Info("Write data to storage", zap.Reflect("data", data))
	// Convert to gob
	buffer := bufio.NewWriter(f)
	ge := gob.NewEncoder(buffer)
	// encode
	if err := ge.Encode(data); err != nil {
		return err
	}
	_ = buffer.Flush()
	return nil
}

// Read data from path to data variable
func Read(path string, data interface{}) error {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	// handle for file close
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(ErrFileStorageNotClose)
		}
	}(f)
	gd := gob.NewDecoder(f)
	if err := gd.Decode(data); err != nil {
		if err != io.EOF {
			return err
		}
	}
	// logger.Info("Read data from storage: "+path, zap.Reflect("data", data))
	return nil
}
