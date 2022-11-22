package main

import (
	"log"
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
	mux := router.NewRouter(todoDB)

<<<<<<< HEAD
	// TODO: ここから実装を行う
	//http.HandleFunc("/healthz", healthzHandler)
	mux.HandleFunc("/healthz", HealthzHandler)
	// handler := &healthzHandler{}
	// http.Handle("/healthz", handler)
	http.ListenAndServe(":8080",mux)//muxはハンドラ、すなわちリクエストを受けてレスポンスを返す処理を表す。
	//echoHandlerなど、いろいろな種類のHandlerがある。
	
	//hundlefuncは第二引数が関数、hundleはポインタ。

=======
	// TODO: サーバーをlistenする
>>>>>>> upstream/main

	return nil
}
