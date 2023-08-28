package services

import (
	"os"

	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
)

// WebgoConsoleLog 代表控制台输出
type WebgoConsoleLog struct {
	WebgoLog
}

// NewWebgoConsoleLog 实例化WebgoConsoleLog
func NewWebgoConsoleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &WebgoConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
