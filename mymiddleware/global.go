package mymiddleware

import (
	"context"
	"log"
	"net/http"
)

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)

		// if err != nil {
		// 	http.Error(w, http.StatusText(404), 404)
		// 	return
		// }
		ctx := context.WithValue(r.Context(), "article", "111")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
