package host

import (
	"io/ioutil"
	"net/http"
	"time"

	glb "github.com/cdutwhu/n3-frontend-test/global"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/middleware"
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

	cfg := glb.Cfg
	port := cfg.WebService.Port
	route := cfg.Route

	mMtx := initMutex()
	mSrvIP := initMapSrvIP()

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

		pvalues := c.QueryParams()
		service, url := "", "http://"
		if ok, v := url1Value(pvalues, 0, "service"); ok {
			if ip, ok := mSrvIP[v]; ok {
				service = v
				url += ip
			}
		}
		if service == "" {
			return c.JSON(http.StatusBadRequest, result{
				Data:  nil,
				Info:  "",
				Error: "Service only for [PRIVACY, SIF2JSON, CSV2JSON]",
			})
		}

		// fPln("------", url)

		// ---------------------------------- //
		time.Sleep(40 * time.Millisecond)
		sp := jaegertracing.CreateChildSpan(c, "Child span for additional processing")
		sp.LogEvent("Test log")
		sp.SetBaggageItem("Test baggage", "baggage")
		sp.SetTag("Test tag", "New tag")
		time.Sleep(40 * time.Millisecond)
		sp.Finish()
		// ---------------------------------- //

		// resp, err := http.Get(url)
		results := jaegertracing.TraceFunction(c, http.Get, url)
		resp := results[0].Interface().(*http.Response)

		if resp == nil {
			return c.JSON(http.StatusNotFound, result{
				Data:  nil,
				Info:  "Couldn't access: " + service,
				Error: "Couldn't access: " + url,
			})
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusNotFound, result{
				Data:  nil,
				Info:  "Accessed: " + service + " " + url,
				Error: fSf("resp Body read fatal: %v", err),
			})
		}

		Data := string(data)
		return c.JSON(http.StatusOK, result{
			Data:  &Data,
			Info:  "",
			Error: "",
		})
	})
}
