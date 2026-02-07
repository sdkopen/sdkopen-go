package restserver

import "net/http"

type IMiddleware interface {
	Apply(ctx ServerContext) error
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == Options.String() {
			return
		}

		next.ServeHTTP(w, r)
	})
}
