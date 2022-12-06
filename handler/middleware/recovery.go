package middleware

import (
	"fmt"
	"net/http"
)

func Recovery(h http.Handler) http.Handler { //Handlerを受け取り、recover処理を追加する。
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if c := recover(); c != nil { //panicをrecoverで受け取り、なかったことにする。
				fmt.Println("Recovered :", c)
			}
		}()
		h.ServeHTTP(w, r) //Handlerの定義は「ServerHTTPを備えたインターフェース」なので、必ずServerHTTPを持つことが保証される。
	}
	return http.HandlerFunc(fn) //handlerfuncはfnを自動的にServerHTTPに変換し、Handlerを作成する。
}
