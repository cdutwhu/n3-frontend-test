package host

import (
	"fmt"
	"strings"
	"sync"

	glb "github.com/cdutwhu/n3-frontend-test/global"
	cmn "github.com/cdutwhu/n3-util/common"
)

var (
	fPf         = fmt.Printf
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	sReplaceAll = strings.ReplaceAll
	sJoin       = strings.Join

	localIP       = cmn.LocalIP
	isXML         = cmn.IsXML
	isJSON        = cmn.IsJSON
	setLog        = cmn.SetLog
	log           = cmn.Log
	warnOnErr     = cmn.WarnOnErr
	failOnErr     = cmn.FailOnErr
	mustWriteFile = cmn.MustWriteFile
	mapFromStruct = cmn.MapFromStruct
	url1Value     = cmn.URL1Value
)

func initMutex() map[string]*sync.Mutex {
	mMtx := make(map[string]*sync.Mutex)
	for _, v := range mapFromStruct(glb.Cfg.Route) {
		mMtx[v.(string)] = &sync.Mutex{}
	}
	return mMtx
}

func initMapSrvIP() map[string]string {
	mSrvIP := make(map[string]string)
	for k, v := range mapFromStruct(glb.Cfg.ServiceIP) {
		mSrvIP[k] = v.(string)
	}
	return mSrvIP
}

type result struct {
	Data  *string `json:"data"`
	Info  string  `json:"info"`
	Error string  `json:"error"`
}
