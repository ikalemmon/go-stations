package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mileusna/useragent"
)

func OSname(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userAgents := r.UserAgent()
		ua := useragent.Parse(userAgents)

		var ctx = r.Context()
		type OSname string
		k := OSname("OS")
		ctx = context.WithValue(ctx, k, ua.OS)
		r = r.Clone(ctx)

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

type Log struct {
	TimeStamp time.Time `json:"time_stamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

func acccessLog(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		var latency = (end.UnixNano() / int64(time.Millisecond)) - (start.UnixNano() / int64(time.Millisecond))
		var log = Log{}
		log.TimeStamp = time.Now()
		log.Latency = latency
		log.Path = r.URL.Path
		h = OSname(h)
		log.OS = r.Context().Value("OS").(string)
		fmt.Println(json.NewEncoder(w).Encode(log))
	}
	return http.HandlerFunc(fn)
}
