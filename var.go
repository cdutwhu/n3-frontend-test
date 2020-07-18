package main

import (
	"fmt"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
)

var (
	fPln = fmt.Println

	setLog        = fn.SetLog
	logWhen       = fn.LoggerWhen
	failOnErr     = fn.FailOnErr
	failOnErrWhen = fn.FailOnErrWhen
	localIP       = net.LocalIP
	env2Struct    = rflx.Env2Struct
)
