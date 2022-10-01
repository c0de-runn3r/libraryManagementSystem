package utils

import (
	"fmt"
	"os"
	"time"
)

const (
	LOG_LEVEL  string = ""
	SILENT_LOG bool   = false
	LOG_PATH   string = "/Users/bodya-pc/go/libraryManagementSystem/backend/logs"
)

func GetFilenameDate() string {
	const layout = "02-01-2006"
	t := time.Now()
	return t.Format(layout) + ".log"
}

func Log(level string, message string) {
	if LOG_LEVEL != "disabled" {
		const layout = "15:04:05"
		StringToLog := fmt.Sprintf("(%v) [%s] %s", time.Now().Format(layout), level, message)
		os.Chdir(LOG_PATH)
		file, err := os.OpenFile(GetFilenameDate(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			Log("ERROR", "Problem occured when creating a file.")
		}
		if SILENT_LOG == false {
			println(StringToLog)
		}
		file.WriteString(StringToLog)
		file.Close()
	}
}
