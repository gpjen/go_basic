package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func InitializeChronJobs() (*cron.Cron, error) {
	c := cron.New()

	// tugas wajib setiap server dijalankan
	deleteLoggerFile()

	// hapus log file setiap 1 bulan
	_, err := c.AddFunc("0 0 1 * *", deleteLoggerFile)
	if err != nil {
		Logger.Println("Failed to add cron job :", err)
		return nil, fmt.Errorf("error adding cron job: %w", err)
	}

	c.Start()

	return c, nil
}

func deleteLoggerFile() {

	daysAgo := time.Now().AddDate(0, 0, -30)

	Logger.Printf("delete log files before date ( %s )", daysAgo.Format("02-01-2006"))

	logDir := "./logs"
	files, err := os.ReadDir(logDir)
	if err != nil {
		Logger.Fatalf("Failed to read log directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			if strings.HasSuffix(fileName, "-logfile.log") {
				datePart := strings.TrimSuffix(fileName, "-logfile.log")
				fileDate, err := time.Parse("02-01-2006", datePart)
				if err != nil {
					Logger.Printf("Failed to parse date from file name %s: %v", fileName, err)
					continue
				}

				if fileDate.Before(daysAgo) {
					filePath := filepath.Join(logDir, fileName)
					err := os.Remove(filePath)
					if err != nil {
						Logger.Printf("Failed to delete log file %s: %v", fileName, err)
					} else {
						Logger.Printf("Deleted old log file: %s", fileName)
					}
				}
			}
		}
	}
}
