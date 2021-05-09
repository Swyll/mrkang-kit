package main

import (
	"mrkang-kit/log"
	"mrkang-kit/log/zaplogger"
)

func main() {
	log.Init(zaplogger.NewLog())
	log.Info("wwwwww")
}
