package main

import (
	"os"

	cfg "github.com/cdutwhu/n3-frontend-test/config"
	"github.com/cdutwhu/n3-frontend-test/host"
	eg "github.com/cdutwhu/n3-util/n3errs"
)

func main() {
	failOnErrWhen(!cfg.InitEnvVarFromTOML("Cfg"), "%v: Config Init Error", eg.CFG_INIT_ERR)

	cfg := env2Struct("Cfg", &cfg.Config{}).(*cfg.Config)
	ws, logfile, servicename := cfg.WebService, cfg.LogFile, cfg.ServiceName

	os.Setenv("JAEGER_SERVICE_NAME", servicename)
	os.Setenv("JAEGER_SAMPLER_TYPE", "const")
	os.Setenv("JAEGER_SAMPLER_PARAM", "1")

	setLog(logfile)
	fPln(logWhen(true, "[%s] Hosting on: [%v:%d], version [%v]", servicename, localIP(), ws.Port, ws.Version))

	done := make(chan string)
	go host.HTTPAsync()
	<-done
}
