package main

import (
	"github.com/Swyll/mrkang-kit/log"
	"github.com/Swyll/mrkang-kit/log/zaplogger"
)

func main() {
	log.Init(zaplogger.NewLog())
	log.Info("wwwwww")
}
