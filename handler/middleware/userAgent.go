package middleware

import (
	"context"
	"net/http"

	"github.com/mileusna/useragent"
)

func userAgent(r *http.Request) *http.Request {
	type OSname string
	userAgents := r.UserAgent()
	ua := useragent.Parse(userAgents)

	var ctx = r.Context()
	k := OSname("OS")
	ctx = context.WithValue(ctx, k, ua.OS)
	return r.Clone(ctx)
}
