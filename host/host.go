package host

import (
	glb "github.com/cdutwhu/n3-frontend-test/global"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/middleware"
)

// HTTPAsync : Host a HTTP Server
func HTTPAsync() {

	e := echo.New()

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

	// Server
	defer e.Start(fSf(":%d", port))

	htmlPath := "/"
	e.File(htmlPath, "./www/service.html")
	e.Static(htmlPath, "./www/") //             "/" is html - ele - <src>'s path
}
