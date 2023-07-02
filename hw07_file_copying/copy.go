package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	progressbar "github.com/schollz/progressbar/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

const zeroLimit int64 = 0

func Copy(fromPath, toPath string, offset, limit int64) error {
	var err error

	fromFile, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("failed open from file %w", err)
	}

	defer func() {
		if deferErr := fromFile.Close(); deferErr != nil {
			log.Panicf("failed close from file %v", deferErr)
		}
	}()

	fileInfo, err := fromFile.Stat()
	if err != nil {
		return fmt.Errorf("failed get stat %w", err)
	}

	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit == zeroLimit {
		limit = fileInfo.Size()
	}

	if offset > 0 {
		_, err = fromFile.Seek(offset, io.SeekStart)
		if err != nil {
			return fmt.Errorf("set offset with seek %w", err)
		}
	}

	toFile, err := os.OpenFile(toPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0o755))
	if err != nil {
		return fmt.Errorf("failed open to file %w", err)
	}

	defer func() {
		if deferErr := toFile.Close(); deferErr != nil {
			log.Panicf("failed close to file %v", deferErr)
		}
	}()

	bar := progressbar.DefaultBytes(limit, "Copying: ")

	_, err = io.CopyN(io.MultiWriter(toFile, bar), fromFile, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("copy data: %w", err)
	}

	return nil
}
