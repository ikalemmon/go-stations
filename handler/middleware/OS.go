package middleware

import (
	"context"
	"net/http"

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
