package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	Logger  = logrus.New()
	LogFile *os.File
)

func LoggerInit() {
	f, err := os.OpenFile("core_logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	LogFile = f
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	//Logger.SetOutput(f)
	Logger.SetOutput(os.Stdout)
}
