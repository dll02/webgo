package services

import (
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
	"io"
)

type WebgoCustomLog struct {
	WebgoLog
}

func NewWebgoCustomLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)
	output := params[4].(io.Writer)

	log := &WebgoConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.SetOutput(output)
	log.c = c
	return log, nil
}
