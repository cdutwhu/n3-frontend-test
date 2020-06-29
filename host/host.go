package host

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	cfg "github.com/cdutwhu/n3-frontend-test/config"
	eg "github.com/cdutwhu/n3-util/n3errs"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	cltC2J "github.com/nsip/n3-csv2json/Server/go-client"
	cltPRI "github.com/nsip/n3-privacy/Server/go-client"
	cltS2J "github.com/nsip/n3-sif2json/Server/go-client"
)

// HTTPAsync : Host a HTTP Server
func HTTPAsync() {

	e := echo.New()
	defer e.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2G"))

	// Add Jaeger Tracer into Middleware
	c := jaegertracing.New(e, nil)
	defer c.Close()

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	ICfg, err := env2Struct("Cfg", &cfg.Config{})
	failOnErr("%v", err)
	Cfg := ICfg.(*cfg.Config)
	port := Cfg.WebService.Port
	route := Cfg.Route
	mMtx := initMutex(&route)
	// serviceIP := Cfg.ServiceIP
	// mSrvIP := initMapSrvIP(&serviceIP)

	// Server
	defer e.Start(fSf(":%d", port))

	// ------------------------------------------------------------------------------------ //

	path := route.PAGE
	e.File(path, "./www/service.html")
	e.Static(path, "./www/")

	// ------------------------------------------------------------------------------------ //

	path = route.HELP
	e.GET(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		var (
			RetStat   = http.StatusOK
			RetStr    string
			RetInfo   string
			RetErr    error
			RetErrStr string
		)

		var (
			pvalues = c.QueryParams()
			ctx     = c.Request().Context()
		)

		_, service := url1Value(pvalues, 0, "service")

		// dialog.Alert(service)

		// pvalues := c.QueryParams()
		// service, url := "", "http://"
		// if ok, v := url1Value(pvalues, 0, "service"); ok {
		// 	if ip, ok := mSrvIP[v]; ok {
		// 		service = v
		// 		url += ip
		// 	}
		// }
		// if service == "" {
		// 	return c.JSON(http.StatusBadRequest, result{
		// 		Data:  nil,
		// 		Info:  "",
		// 		Error: "Service only for [PRIVACY, SIF2JSON, CSV2JSON]",
		// 	})
		// }

		// ------------------------------------------------------------------------ //
		// time.Sleep(40 * time.Millisecond)
		// sp := jaegertracing.CreateChildSpan(c, "Child span for additional processing")
		// sp.LogEvent("Test log")
		// sp.SetBaggageItem("Test baggage", "baggage")
		// sp.SetTag("Test tag", "New tag")
		// time.Sleep(40 * time.Millisecond)
		// sp.Finish()
		// ------------------------------------------------------------------------ //

		// resp, err := http.Get(url)

		// results := jaegertracing.TraceFunction(c, http.Get, url)
		// resp := results[0].Interface().(*http.Response)

		switch service {
		case "privacy", "PRIVACY", "Privacy":
			RetStr, RetErr = cltPRI.DOwithTrace(
				ctx,
				"cfg-clt-privacy.toml",
				"HELP",
				nil,
			)

		case "sif2json", "SIF2JSON":
			RetStr, RetErr = cltS2J.DOwithTrace(
				ctx,
				"cfg-clt-sif2json.toml",
				"HELP",
				nil,
			)

		case "csv2json", "CSV2JSON":
			RetStr, RetErr = cltC2J.DOwithTrace(
				ctx,
				"cfg-clt-csv2json.toml",
				"HELP",
				nil,
			)

		default:
			RetErr = warnOnErr("%v: %s", eg.PARAM_NOT_SUPPORTED, service)
			RetStat = http.StatusBadRequest
			goto RET
		}

	RET:
		if RetErr != nil {
			RetErrStr = RetErr.Error()
		}
		return c.JSON(RetStat, result{
			Data:  RetStr,
			Info:  RetInfo,
			Error: RetErrStr,
		})
	})

	// --------------------------------------------- //

	path = route.UPLOAD
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		var (
			RetStat   = http.StatusOK
			RetStr    string
			RetInfo   string
			RetEmpty  bool
			RetErr    error
			RetErrStr string
			m         = make(map[string]interface{})
		)

		var (
			Service string
			Fn      string
			ToNATS  bool
			Data    []byte
			User    string
			Ctx     string
			RW      string
			Object  string
			SIFVer  string
		)

		var (
			pvalues = c.QueryParams()
			ctx     = c.Request().Context()
		)

		if ok, s := url1Value(pvalues, 0, "service"); ok {
			Service = s
		}
		if ok, fn := url1Value(pvalues, 0, "fn"); ok {
			Fn = fn
		}
		if ok, n := url1Value(pvalues, 0, "tonats"); ok && n == "true" {
			ToNATS = true
		}
		if ok, user := url1Value(pvalues, 0, "user"); ok {
			User = user
		}
		if ok, ctx := url1Value(pvalues, 0, "ctx"); ok {
			Ctx = ctx
		}
		if ok, rw := url1Value(pvalues, 0, "rw"); ok {
			RW = rw
		}
		if ok, obj := url1Value(pvalues, 0, "object"); ok {
			Object = obj
		}
		if ok, sv := url1Value(pvalues, 0, "sv"); ok {
			SIFVer = sv
		}
		if Data, RetErr = ioutil.ReadAll(c.Request().Body); RetErr != nil {
			warnOnErr("%v: %s", RetErr, RetStr)
			RetStat = http.StatusInternalServerError
			goto RET
		}

		// fPln(Service, ToNATS, User, Ctx, RW, Object)
		// dialog.Alert(Service)

		switch Service {
		case "privacy", "PRIVACY", "Privacy":
			RetStr, RetErr = cltPRI.DOwithTrace(
				ctx,
				"cfg-clt-privacy.toml",
				Fn,
				&cltPRI.Args{
					Policy: Data,
					User:   User,
					Ctx:    Ctx,
					RW:     RW,
					Object: Object,
					Data:   Data,
				},
			)

		case "sif2json", "SIF2JSON":
			RetStr, RetErr = cltS2J.DOwithTrace(
				ctx,
				"cfg-clt-sif2json.toml",
				Fn,
				&cltS2J.Args{
					Data:   Data,
					Ver:    SIFVer,
					ToNATS: ToNATS,
				},
			)

		case "csv2json", "CSV2JSON":
			RetStr, RetErr = cltC2J.DOwithTrace(
				ctx,
				"cfg-clt-csv2json.toml",
				Fn,
				&cltC2J.Args{
					Data:   Data,
					ToNATS: ToNATS,
				},
			)

		default:
			RetErr = warnOnErr("%v: %s", eg.PARAM_NOT_SUPPORTED, Service)
			RetStat = http.StatusBadRequest
			goto RET
		}

		if RetErr != nil {
			warnOnErr("%v: %s", RetErr, RetStr)
			RetStat = http.StatusInternalServerError
			goto RET
		}

		failOnErr("json.Unmarshal ... %v: %s", json.Unmarshal([]byte(RetStr), &m), RetStr)

		if retStr, ok := m["data"]; ok && retStr != nil {
			RetStr = m["data"].(string)
		}
		if retInfo, ok := m["info"]; ok && retInfo != nil {
			RetInfo = m["info"].(string)
		}
		if retEmpty, ok := m["empty"]; ok && retEmpty != nil {
			RetEmpty = m["empty"].(bool)
		}
		if retErrStr, ok := m["error"]; ok && retErrStr != "" {
			RetErrStr = retErrStr.(string)
			RetStat = http.StatusInternalServerError
		}

	RET:
		if RetErr != nil {
			RetErrStr = RetErr.Error()
		}
		return c.JSON(RetStat, result{
			Data:  RetStr,
			Info:  RetInfo,
			Empty: RetEmpty,
			Error: RetErrStr,
		})
	})
}
