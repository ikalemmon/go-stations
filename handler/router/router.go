package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	//HandleFuncはアドレスに対してhandler関数を対応させる(handler.ServerHTTPがhandler関数、handler関数とhandlerは違う。)
	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP) // NewHealthzHandler→serverHandle。
	mux.HandleFunc("/todos", handler.NewTODOHandler(service.NewTODOService(todoDB)).ServeHTTP)
	mux.HandleFunc("/do-panic", handler.NewPanicHandler().ServeHTTP)
	mux.Handle("/do-panic-middleware", middleware.Recovery(handler.NewPanicHandler()))
	return mux
}
