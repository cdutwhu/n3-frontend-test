package main

import (
	eg "github.com/cdutwhu/json-util/n3errs"
	g "github.com/cdutwhu/n3-frontend-test/global"
	"github.com/cdutwhu/n3-frontend-test/host"
)

func main() {
	failOnErrWhen(!g.Init(), "%v: Global Config Init Error", eg.CFG_INIT_ERR)

	cfg := g.Cfg
	ws, logfile, servicename := cfg.WebService, cfg.LogFile, cfg.ServiceName

	setLog(logfile)
	fPln(logWhen(true, "[%s] Hosting on: [%v:%d], version [%v]", servicename, localIP(), ws.Port, ws.Version))

	done := make(chan string)
	go host.HTTPAsync()
	<-done
}
