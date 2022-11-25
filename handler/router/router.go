package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	//HandleFuncはアドレスに対してhandler関数を対応させる(handler.ServerHTTPがhandler関数、handler関数とhandlerは違う。)
	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP) // NewHealthzHandler→serverHandle。
	return mux
}
