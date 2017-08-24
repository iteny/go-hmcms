package main

import (
	"go-hmcms/controllers/intendant"
	"go-hmcms/models/ini"
	_ "go-hmcms/models/sqlm"
	"go-hmcms/mymiddleware"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func main() {
	server := &server{Addr: "80", ReadTimeout: 10, WriteTimeout: 10}
	if servPort := ini.Value("servSet", "port"); servPort != "" {
		server.Addr = servPort
	}
	if servReadTimeout, _ := ini.Int("servSet", "ReadTimeout"); servReadTimeout != 0 {
		server.ReadTimeout = time.Duration(servReadTimeout)
	}
	if servWriteTimeout, _ := ini.Int("servSet", "WriteTimeout"); servWriteTimeout != 0 {
		server.WriteTimeout = time.Duration(servWriteTimeout)
	}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	//mount website back-stage routes
	r.Mount("/intendant", intendantRoutes())
	//Easily serve static files
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(r, "/static", http.Dir(filesDir))
	//http server
	s := &http.Server{
		Addr:           ":" + server.Addr,
		Handler:        r,
		ReadTimeout:    server.ReadTimeout * time.Second,
		WriteTimeout:   server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}
	fs := http.StripPrefix(path, http.FileServer(root))
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"
	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

//website back-stage routes
func intendantRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(mymiddleware.ArticleCtx)
	r.Get("/", intendant.LoginCtl.Login)
	r.Post("/login", intendant.LoginCtl.LoginGo)
	r.Get("/index", intendant.IndexCtl.Index)
	r.Get("/home", intendant.IndexCtl.Home)
	r.Post("/getLeftMenu", intendant.IndexCtl.GetLeftMenu)
	r.NotFound(intendant.LoginCtl.Error)
	// r.Get(pattern, handlerFn)
	//menu set routes
	r.Route("/site", func(r chi.Router) {
		r.Get("/menu", intendant.SiteCtl.Menu)
		r.Post("/sortmenu", intendant.SiteCtl.SortMenu)
		r.Route("/addEditMenu", func(r chi.Router) {
			r.Get("/", intendant.SiteCtl.AddEditMenuGet)
			r.Get("/?pid={articleID}", intendant.SiteCtl.AddEditMenuGet)
			r.Get("/?id={articleID}", intendant.SiteCtl.AddEditMenuGet)
			r.Post("/", intendant.SiteCtl.AddEditMenuPost)
		})
	})
	return r
}
