package host

import (
	"net/http"

	cfg "github.com/cdutwhu/n3-frontend-test/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/middleware"
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

	Cfg := env2Struct("Cfg", &cfg.Config{}).(*cfg.Config)
	port := Cfg.WebService.Port
	route := Cfg.Route
	mMtx := initMutex(route)
	// serviceIP := Cfg.ServiceIP
	// mSrvIP := initMapSrvIP(serviceIP)

	// Server
	defer e.Start(fSf(":%d", port))

	// ------------------------------------------------------------------------------------ //

	path := route.PAGE
	e.File(path, "./www/service.html")
	e.Static(path, "./www/") //             "/" is html - ele - <src>'s path

	// ------------------------------------------------------------------------------------ //

	path = route.HELP
	e.GET(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

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

		str, err := cltS2J.DO(
			"cfg-clt-sif2json.toml",
			"HELP",
			cltS2J.Args{
				File:      "",
				Ver:       "3.4.5",
				WholeDump: false,
				ToNATS:    false,
			},
		)
		if err == nil {
			return c.JSON(http.StatusOK, result{
				Data:  &str,
				Info:  "",
				Error: "",
			})
		}
		return c.JSON(http.StatusInternalServerError, result{
			Data:  &str,
			Info:  "",
			Error: err.Error(),
		})
	})
}
