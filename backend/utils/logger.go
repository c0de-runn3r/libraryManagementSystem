package utils

import (
	"errors"
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

func checkLevel(level string) (bool, error) {
	switch os.Getenv("LOG_LEVEL") {
	case "error":
		if level == "error" {
			return true, nil
		}
	case "warning":
		if level == "error" || level == "warning" {
			return true, nil
		}
	case "info":
		if level == "error" || level == "warning" || level == "info" {
			return true, nil
		}
	case "debug":
		if level == "error" || level == "warning" || level == "info" || level == "debug" {
			return true, nil
		}
	case "disabled":
		return false, nil
	default:
		err := errors.New("Inappropriate environment variable 'LOG_LEVEL'")
		return false, err
	}
	return false, nil
}

func Log(level string, message string) {
	if level != "error" && level != "warning" && level != "info" && level != "debug" {
		fmt.Println(errors.New("Inappropriate argument in Log func. Incorrect log level."))
		os.Exit(1)
		return
	}
	if len(message) == 0 {
		fmt.Println("Log error. The message string is empty.")
		os.Exit(1)
		return
	}
	lvl, err := checkLevel(level)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	if lvl == true {
		const layout = "15:04:05"
		StringToLog := fmt.Sprintf("(%v) [%s] %s\n", time.Now().Format(layout), strings.ToUpper(level), message)
		file, err := os.OpenFile(GetFilenameDate(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Problem occured when creating a file. %s\n", err)
		} else {
			file.WriteString(StringToLog)
			file.Close()
		}
		silentLog := os.Getenv("SILENT_LOG")
		if silentLog == "false" || silentLog == "" {
			fmt.Printf(StringToLog)
		} else if os.Getenv("SILENT_LOG") != "true" {
			fmt.Println(errors.New("Inappropriate environment variable 'SILENT_LOG'"))
			os.Exit(1)
		}
	}
}
