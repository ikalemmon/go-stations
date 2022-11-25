package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TechBowl-japan/go-stations/db"
	"github.com/TechBowl-japan/go-stations/handler/router"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func realMain() error {
	// config values
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		return err
	}
	defer todoDB.Close()

	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする

	// TODO: ここから実装を行う
	//(muxが帰ってきているが、NewRouterの内部で行っているのはhttp:NewServeMux()で、muxの型は*http.ServeMux.)
	mux := router.NewRouter(todoDB)

	// TODO: サーバーをlistenする
	//addressとhandlerを持つserverを返す(第二引数はhandler、nilにすることが多い)
	http.ListenAndServe(port, mux) //

	//ServeMux→マルチプレクサ。アドレスとhandlerの対応を保持するもの。
	//Hundler：ServeHTTPメソッドをもつもの。

	return nil
}
