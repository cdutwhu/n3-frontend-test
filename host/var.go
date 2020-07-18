package host

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/n3-util/rest"
)

var (
	fPf         = fmt.Printf
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	sReplaceAll = strings.ReplaceAll
	sJoin       = strings.Join
	sToUpper    = strings.ToUpper

	localIP       = net.LocalIP
	isXML         = judge.IsXML
	isJSON        = judge.IsJSON
	setLog        = fn.SetLog
	log           = fn.Logger
	warnOnErr     = fn.WarnOnErr
	failOnErr     = fn.FailOnErr
	mustWriteFile = io.MustWriteFile
	url1Value     = rest.URL1Value
	env2Struct    = rflx.Env2Struct
	struct2Env    = rflx.Struct2Env
	struct2Map    = rflx.Struct2Map
)

func initMutex(route interface{}) map[string]*sync.Mutex {
	mMtx := make(map[string]*sync.Mutex)
	for _, v := range struct2Map(route) {
		mMtx[v.(string)] = &sync.Mutex{}
	}
	return mMtx
}

func initMapSrvIP(serviceIP interface{}) map[string]string {
	mSrvIP := make(map[string]string)
	for k, v := range struct2Map(serviceIP) {
		mSrvIP[k] = v.(string)
	}
	return mSrvIP
}

type result struct {
	Data  string `json:"data"`
	Info  string `json:"info"`
	Empty bool   `json:"empty"`
	Error string `json:"error"`
}
