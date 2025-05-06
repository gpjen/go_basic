package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	LogFile *os.File
	Logger  *log.Logger
)

func InitLogger() {
	var err error
	LogFile, err = GetLogFile()
	if err != nil {
		log.Fatalf("Failed to initialize log file: %v", err)
	}

	var multiWriter io.Writer

	if os.Getenv("NODE_ENV") != "production" {
		multiWriter = io.MultiWriter(LogFile, os.Stdout)
	} else {
		multiWriter = LogFile
	}

	Logger = log.New(multiWriter, "", log.LstdFlags)
}

func GetLogFile() (*os.File, error) {
	fileName := fmt.Sprintf("%s-logfile.log", time.Now().Format("02-01-2006"))
	pathFile := fmt.Sprintf("./logs/%s", fileName)

	if err := os.MkdirAll("./logs", os.ModePerm); err != nil {
		return nil, err
	}

	if _, err := os.Stat(pathFile); os.IsNotExist(err) {
		file, err := os.Create(pathFile)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	file, err := os.OpenFile(pathFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
