package logging

import (
	"os"
	"fmt"
)

func Info(format string, args ...interface{}) {
    fmt.Printf("[INFO] "+format+"\n", args...)
}

func Warn(format string, args ...interface{}) {
    fmt.Printf("[WARN] "+format+"\n", args...)
}

func Error(format string, args ...interface{}) {
    fmt.Printf("[ERRO] "+format+"\n", args...)
}

func Debug(format string, args ...interface{}) {
    if os.Getenv("DEBUG") == "true" {
        fmt.Printf("[DEBU] "+format+"\n", args...)
    }
}

func Panic(format string, args ...interface{}) {
    fmt.Printf("[PANIC] "+format+"\n", args...)
	os.Exit(1)
}

