package main

import (
	"os"

	"github.com/fxckcode/interactive-gemini-go-chat/utils"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func init() {
    log.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
        ForceColors:   true,
    })
    log.SetOutput(os.Stdout)
    log.SetLevel(logrus.InfoLevel)
}

func main() {
	utils.ClearConsole()
}
