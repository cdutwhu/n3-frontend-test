package host

import (
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

	htmlPath := "/"
	e.File(htmlPath, "./www/service.html")
	e.Static(htmlPath, "./www/") //             "/" is html - ele - <src>'s path

	// // Maybe Auth middleware dislike long request body ? manually check
	// e.POST(g.Cfg.Route.FilePub, postFileToNode)

	// // Group
	// api := e.Group(g.Cfg.Group.API)
	// api.Use(middleware.Logger())
	// api.Use(middleware.Recover())
	// api.Use(middleware.BodyLimit("2G"))

	// uname := ""
	// // BasicAuth ( Big Body has ERR_CONNECTION_RESET in this )
	// api.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	fPln("---------------------in basicAuth-----------------------------", username, password)
	// 	if g.CurCtx = ctxFromCredential(username, password); g.CurCtx == "" {
	// 		return false, c.String(http.StatusUnauthorized, "wrong username or password")
	// 	}
	// 	uname = username
	// 	return true, nil
	// }))

	// // api Route
	// // api.GET("/filetest", func(c echo.Context) error { return c.File("/home/qing/Desktop/index.html") })
	// api.GET(g.Cfg.Route.Greeting, func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, "Hello, "+uname+". n3client is running @ "+time.Now().Format("2006-01-02 15:04:05.000"))
	// })
	// api.GET(g.Cfg.Route.ID, getIDList)
	// api.GET(g.Cfg.Route.Obj, getObject)
	// api.GET(g.Cfg.Route.Scm, getSchema)
	// api.POST(g.Cfg.Route.Pub, postToNode)
	// api.POST(g.Cfg.Route.GQL, postQueryGQL)
	// api.DELETE(g.Cfg.Route.Del, delFromNode)

	// Server
	e.Start(fSf(":%d", 1320))
}
