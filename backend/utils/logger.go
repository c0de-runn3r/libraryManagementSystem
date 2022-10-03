package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func GetFilenameDate() string {
	const layout = "02-01-2006"
	t := time.Now()
	return os.Getenv("LOG_PATH") + "/" + t.Format(layout) + ".log"
}

func checkLevel(level string) bool {
	switch os.Getenv("LOG_LEVEL") {
	case "error":
		if level == "error" {
			return true
		}
	case "warning":
		if level == "error" || level == "warning" {
			return true
		}
	case "info":
		if level == "error" || level == "warning" || level == "info" {
			return true
		}
	case "debug":
		if level == "error" || level == "warning" || level == "info" || level == "debug" {
			return true
		}
	case "disabled":
		return false
	default:
		fmt.Println("Inappropriate logger level!")
		return false
	}
	return false
}

func Log(level string, message string) {
	if len(message) == 0 {
		fmt.Println("Log message is null")
	}
	if checkLevel(level) == true {
		const layout = "15:04:05"
		StringToLog := fmt.Sprintf("(%v) [%s] %s\n", time.Now().Format(layout), strings.ToUpper(level), message)
		file, err := os.OpenFile(GetFilenameDate(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Problem occured when creating a file. %s\n", err)
		} else {
			file.WriteString(StringToLog)
			file.Close()
		}
		if os.Getenv("SILENT_LOG") == "false" {
			fmt.Printf(StringToLog)
		}
	}
}
