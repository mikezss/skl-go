package models

import (
	"github.com/astaxie/beego/logs"
)

var ticketlog *logs.BeeLogger

func Getlog() *logs.BeeLogger {
	if ticketlog == nil {
		ticketlog = logs.NewLogger()
		ticketlog.SetLogFuncCallDepth(4)
		//ticketlog.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
		//ticketlog.SetLogger(logs.AdapterFile, `{"filename":"ticket.log","daily":true,"maxdays":60,"color":true}`)
		ticketlog.SetLogger(logs.AdapterMultiFile, `{"filename":"ticket.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
		//ticketlog.Async()
		//ticketlog.Async(1e3)
	}
	return ticketlog
}
