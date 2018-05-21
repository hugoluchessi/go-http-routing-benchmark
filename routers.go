// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"

	// If you add new routers please:
	// - Keep the benchmark functions etc. alphabetically sorted
	// - Make a pull request (without benchmark results) at
	//   https://github.com/julienschmidt/go-http-routing-benchmark
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/go-playground/lars"
	// "github.com/daryl/zeus"
	"github.com/dimfeld/httptreemux"
	"github.com/gin-gonic/gin"
	"github.com/go-martini/martini"
	"github.com/gorilla/mux"
	"github.com/hugoluchessi/badger"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo"
	vulcan "github.com/mailgun/route"
	"github.com/mikespook/possum"
	possumrouter "github.com/mikespook/possum/router"
	possumview "github.com/mikespook/possum/view"
	"github.com/naoina/denco"
	_ "github.com/naoina/kocha-urlrouter/doublearray"
	"github.com/plimble/ace"
	"github.com/typepress/rivet"
	"github.com/ursiform/bear"
	"github.com/vanng822/r2router"
)

type route struct {
	method string
	path   string
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

var nullLogger *log.Logger

// flag indicating if the normal or the test handler should be loaded
var loadTestHandler = false

func init() {
	// beego sets it to runtime.NumCPU()
	// Currently none of the contesters does concurrent routing
	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)

	initGin()
	initMartini()
}

// Common
func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {}

func httpHandlerFuncTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

// Ace
func aceHandle(_ *ace.C) {}

func aceHandleWrite(c *ace.C) {
	io.WriteString(c.Writer, c.Param("name"))
}

func aceHandleTest(c *ace.C) {
	io.WriteString(c.Writer, c.Request.RequestURI)
}

func loadAce(routes []route) http.Handler {
	h := []ace.HandlerFunc{aceHandle}
	if loadTestHandler {
		h = []ace.HandlerFunc{aceHandleTest}
	}

	router := ace.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadAceSingle(method, path string, handle ace.HandlerFunc) http.Handler {
	router := ace.New()
	router.Handle(method, path, []ace.HandlerFunc{handle})
	return router
}

// Badger
func badgerHandle(_ http.ResponseWriter, _ *http.Request) {}

func badgerHandleWrite(rw http.ResponseWriter, req *http.Request) {
	rp := badger.GetRouteParamsFromRequest(req)
	value, _ := rp.GetString("name")
	io.WriteString(rw, value)
}

func badgerHandleTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

func loadBadger(routes []route) http.Handler {
	mux := badger.NewMux()
	router := mux.AddRouter("")

	h := http.HandlerFunc(badgerHandle)
	if loadTestHandler {
		h = http.HandlerFunc(badgerHandleTest)
	}

	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}

	return mux
}

func loadBadgerSingle(method, path string, handle http.Handler) http.Handler {
	mux := badger.NewMux()
	router := mux.AddRouter("")
	router.Handle(method, path, handle)
	return mux
}

// bear
func bearHandler(_ http.ResponseWriter, _ *http.Request, _ *bear.Context) {}

func bearHandlerWrite(w http.ResponseWriter, _ *http.Request, ctx *bear.Context) {
	io.WriteString(w, ctx.Params["name"])
}

func bearHandlerTest(w http.ResponseWriter, r *http.Request, _ *bear.Context) {
	io.WriteString(w, r.RequestURI)
}

func loadBear(routes []route) http.Handler {
	h := bearHandler
	if loadTestHandler {
		h = bearHandlerTest
	}

	router := bear.New()
	re := regexp.MustCompile(":([^/]*)")
	for _, route := range routes {
		switch route.method {
		case "GET", "POST", "PUT", "PATCH", "DELETE":
			router.On(route.method, re.ReplaceAllString(route.path, "{$1}"), h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return router
}

func loadBearSingle(method string, path string, handler bear.HandlerFunc) http.Handler {
	router := bear.New()
	switch method {
	case "GET", "POST", "PUT", "PATCH", "DELETE":
		router.On(method, path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return router
}

// Denco
func dencoHandler(w http.ResponseWriter, r *http.Request, params denco.Params) {}

func dencoHandlerWrite(w http.ResponseWriter, r *http.Request, params denco.Params) {
	io.WriteString(w, params.Get("name"))
}

func dencoHandlerTest(w http.ResponseWriter, r *http.Request, params denco.Params) {
	io.WriteString(w, r.RequestURI)
}

func loadDenco(routes []route) http.Handler {
	h := dencoHandler
	if loadTestHandler {
		h = dencoHandlerTest
	}

	mux := denco.NewMux()
	handlers := make([]denco.Handler, 0, len(routes))
	for _, route := range routes {
		handler := mux.Handler(route.method, route.path, h)
		handlers = append(handlers, handler)
	}
	handler, err := mux.Build(handlers)
	if err != nil {
		panic(err)
	}
	return handler
}

func loadDencoSingle(method, path string, h denco.HandlerFunc) http.Handler {
	mux := denco.NewMux()
	handler, err := mux.Build([]denco.Handler{mux.Handler(method, path, h)})
	if err != nil {
		panic(err)
	}
	return handler
}

// Echo
func echoHandler(c echo.Context) error {
	return nil
}

func echoHandlerWrite(c echo.Context) error {
	io.WriteString(c.Response(), c.Param("name"))
	return nil
}

func echoHandlerTest(c echo.Context) error {
	io.WriteString(c.Response(), c.Request().RequestURI)
	return nil
}

func loadEcho(routes []route) http.Handler {
	var h echo.HandlerFunc = echoHandler
	if loadTestHandler {
		h = echoHandlerTest
	}

	e := echo.New()
	for _, r := range routes {
		switch r.method {
		case "GET":
			e.GET(r.path, h)
		case "POST":
			e.POST(r.path, h)
		case "PUT":
			e.PUT(r.path, h)
		case "PATCH":
			e.PATCH(r.path, h)
		case "DELETE":
			e.DELETE(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return e
}

func loadEchoSingle(method, path string, h echo.HandlerFunc) http.Handler {
	e := echo.New()
	switch method {
	case "GET":
		e.GET(path, h)
	case "POST":
		e.POST(path, h)
	case "PUT":
		e.PUT(path, h)
	case "PATCH":
		e.PATCH(path, h)
	case "DELETE":
		e.DELETE(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return e
}

// Gin
func ginHandle(_ *gin.Context) {}

func ginHandleWrite(c *gin.Context) {
	io.WriteString(c.Writer, c.Params.ByName("name"))
}

func ginHandleTest(c *gin.Context) {
	io.WriteString(c.Writer, c.Request.RequestURI)
}

func initGin() {
	gin.SetMode(gin.ReleaseMode)
}

func loadGin(routes []route) http.Handler {
	h := ginHandle
	if loadTestHandler {
		h = ginHandleTest
	}

	router := gin.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadGinSingle(method, path string, handle gin.HandlerFunc) http.Handler {
	router := gin.New()
	router.Handle(method, path, handle)
	return router
}

// go-json-rest/rest
func goJsonRestHandler(w rest.ResponseWriter, req *rest.Request) {}

func goJsonRestHandlerWrite(w rest.ResponseWriter, req *rest.Request) {
	io.WriteString(w.(io.Writer), req.PathParam("name"))
}

func goJsonRestHandlerTest(w rest.ResponseWriter, req *rest.Request) {
	io.WriteString(w.(io.Writer), req.RequestURI)
}

func loadGoJsonRest(routes []route) http.Handler {
	h := goJsonRestHandler
	if loadTestHandler {
		h = goJsonRestHandlerTest
	}

	api := rest.NewApi()
	restRoutes := make([]*rest.Route, 0, len(routes))
	for _, route := range routes {
		restRoutes = append(restRoutes,
			&rest.Route{route.method, route.path, h},
		)
	}
	router, err := rest.MakeRouter(restRoutes...)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	return api.MakeHandler()
}

func loadGoJsonRestSingle(method, path string, hfunc rest.HandlerFunc) http.Handler {
	api := rest.NewApi()
	router, err := rest.MakeRouter(
		&rest.Route{method, path, hfunc},
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	return api.MakeHandler()
}

// gorilla/mux
func gorillaHandlerWrite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	io.WriteString(w, params["name"])
}

func loadGorillaMux(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")
	m := mux.NewRouter()
	for _, route := range routes {
		m.HandleFunc(
			re.ReplaceAllString(route.path, "{$1}"),
			h,
		).Methods(route.method)
	}
	return m
}

func loadGorillaMuxSingle(method, path string, handler http.HandlerFunc) http.Handler {
	m := mux.NewRouter()
	m.HandleFunc(path, handler).Methods(method)
	return m
}

// HttpRouter
func httpRouterHandle(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}

func httpRouterHandleWrite(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	io.WriteString(w, ps.ByName("name"))
}

func httpRouterHandleTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, r.RequestURI)
}

func loadHttpRouter(routes []route) http.Handler {
	h := httpRouterHandle
	if loadTestHandler {
		h = httpRouterHandleTest
	}

	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadHttpRouterSingle(method, path string, handle httprouter.Handle) http.Handler {
	router := httprouter.New()
	router.Handle(method, path, handle)
	return router
}

// httpTreeMux
func httpTreeMuxHandler(_ http.ResponseWriter, _ *http.Request, _ map[string]string) {}

func httpTreeMuxHandlerWrite(w http.ResponseWriter, _ *http.Request, vars map[string]string) {
	io.WriteString(w, vars["name"])
}

func httpTreeMuxHandlerTest(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	io.WriteString(w, r.RequestURI)
}

func loadHttpTreeMux(routes []route) http.Handler {
	h := httpTreeMuxHandler
	if loadTestHandler {
		h = httpTreeMuxHandlerTest
	}

	router := httptreemux.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadHttpTreeMuxSingle(method, path string, handler httptreemux.HandlerFunc) http.Handler {
	router := httptreemux.New()
	router.Handle(method, path, handler)
	return router
}

// LARS
func larsHandler(c lars.Context) {
}

func larsHandlerWrite(c lars.Context) {
	io.WriteString(c.Response(), c.Param("name"))
}

func larsHandlerTest(c lars.Context) {
	io.WriteString(c.Response(), c.Request().RequestURI)
}

func larsNativeHandlerTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

func loadLARS(routes []route) http.Handler {
	var h interface{} = larsHandler
	if loadTestHandler {
		h = larsHandlerTest
	}

	l := lars.New()

	for _, r := range routes {
		switch r.method {
		case "GET":
			l.Get(r.path, h)
		case "POST":
			l.Post(r.path, h)
		case "PUT":
			l.Put(r.path, h)
		case "PATCH":
			l.Patch(r.path, h)
		case "DELETE":
			l.Delete(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return l.Serve()
}

func loadLARSSingle(method, path string, h interface{}) http.Handler {
	l := lars.New()

	switch method {
	case "GET":
		l.Get(path, h)
	case "POST":
		l.Post(path, h)
	case "PUT":
		l.Put(path, h)
	case "PATCH":
		l.Patch(path, h)
	case "DELETE":
		l.Delete(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return l.Serve()
}

// Martini
func martiniHandler() {}

func martiniHandlerWrite(params martini.Params) string {
	return params["name"]
}

func initMartini() {
	martini.Env = martini.Prod
}

func loadMartini(routes []route) http.Handler {
	var h interface{} = martiniHandler
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	router := martini.NewRouter()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, h)
		case "POST":
			router.Post(route.path, h)
		case "PUT":
			router.Put(route.path, h)
		case "PATCH":
			router.Patch(route.path, h)
		case "DELETE":
			router.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	martini := martini.New()
	martini.Action(router.Handle)
	return martini
}

func loadMartiniSingle(method, path string, handler interface{}) http.Handler {
	router := martini.NewRouter()
	switch method {
	case "GET":
		router.Get(path, handler)
	case "POST":
		router.Post(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATCH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}

	martini := martini.New()
	martini.Action(router.Handle)
	return martini
}

// Possum
func possumHandler(c *possum.Context) error {
	return nil
}

func possumHandlerWrite(c *possum.Context) error {
	io.WriteString(c.Response, c.Request.URL.Query().Get("name"))
	return nil
}

func possumHandlerTest(c *possum.Context) error {
	io.WriteString(c.Response, c.Request.RequestURI)
	return nil
}

func loadPossum(routes []route) http.Handler {
	h := possumHandler
	if loadTestHandler {
		h = possumHandlerTest
	}

	router := possum.NewServerMux()
	for _, route := range routes {
		router.HandleFunc(possumrouter.Simple(route.path), h, possumview.Simple("text/html", "utf-8"))
	}
	return router
}

func loadPossumSingle(method, path string, handler possum.HandlerFunc) http.Handler {
	router := possum.NewServerMux()
	router.HandleFunc(possumrouter.Simple(path), handler, possumview.Simple("text/html", "utf-8"))
	return router
}

// R2router
func r2routerHandler(w http.ResponseWriter, req *http.Request, _ r2router.Params) {}

func r2routerHandleWrite(w http.ResponseWriter, req *http.Request, params r2router.Params) {
	io.WriteString(w, params.Get("name"))
}

func r2routerHandleTest(w http.ResponseWriter, req *http.Request, _ r2router.Params) {
	io.WriteString(w, req.RequestURI)
}

func loadR2router(routes []route) http.Handler {
	h := r2routerHandler
	if loadTestHandler {
		h = r2routerHandleTest
	}

	router := r2router.NewRouter()
	for _, r := range routes {
		router.AddHandler(r.method, r.path, h)
	}
	return router
}

func loadR2routerSingle(method, path string, handler r2router.HandlerFunc) http.Handler {
	router := r2router.NewRouter()
	router.AddHandler(method, path, handler)
	return router
}

// Rivet
func rivetHandler() {}

func rivetHandlerWrite(c *rivet.Context) {
	c.WriteString(c.Get("name"))
}

func rivetHandlerTest(c *rivet.Context) {
	c.WriteString(c.Req.RequestURI)
}

func loadRivet(routes []route) http.Handler {
	var h interface{} = rivetHandler
	if loadTestHandler {
		h = rivetHandlerTest
	}

	router := rivet.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadRivetSingle(method, path string, handler interface{}) http.Handler {
	router := rivet.New()

	router.Handle(method, path, handler)

	return router
}

// Mailgun Vulcan
func vulcanHandler(w http.ResponseWriter, r *http.Request) {}

func vulcanHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}

func loadVulcan(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")
	mux := vulcan.NewMux()
	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "<$1>")
		expr := fmt.Sprintf(`Method("%s") && Path("%s")`, route.method, path)
		if err := mux.HandleFunc(expr, h); err != nil {
			panic(err)
		}
	}
	return mux
}

func loadVulcanSingle(method, path string, handler http.HandlerFunc) http.Handler {
	re := regexp.MustCompile(":([^/]*)")
	mux := vulcan.NewMux()
	path = re.ReplaceAllString(path, "<$1>")
	expr := fmt.Sprintf(`Method("%s") && Path("%s")`, method, path)
	if err := mux.HandleFunc(expr, httpHandlerFunc); err != nil {
		panic(err)
	}
	return mux
}

// Zeus
// func zeusHandlerWrite(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, zeus.Var(r, "name"))
// }

// func loadZeus(routes []route) http.Handler {
// 	h := http.HandlerFunc(httpHandlerFunc)
// 	if loadTestHandler {
// 		h = http.HandlerFunc(httpHandlerFuncTest)
// 	}

// 	m := zeus.New()
// 	for _, route := range routes {
// 		switch route.method {
// 		case "GET":
// 			m.GET(route.path, h)
// 		case "POST":
// 			m.POST(route.path, h)
// 		case "PUT":
// 			m.PUT(route.path, h)
// 		case "DELETE":
// 			m.DELETE(route.path, h)
// 		default:
// 			panic("Unknow HTTP method: " + route.method)
// 		}
// 	}
// 	return m
// }

// func loadZeusSingle(method, path string, handler http.HandlerFunc) http.Handler {
// 	m := zeus.New()
// 	switch method {
// 	case "GET":
// 		m.GET(path, handler)
// 	case "POST":
// 		m.POST(path, handler)
// 	case "PUT":
// 		m.PUT(path, handler)
// 	case "DELETE":
// 		m.DELETE(path, handler)
// 	default:
// 		panic("Unknow HTTP method: " + method)
// 	}
// 	return m
// }

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
