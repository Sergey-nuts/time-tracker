package mwpagination

import (
	"context"
	"net/http"
	"strconv"
)

type CustomKey string

const PageId CustomKey = "page"

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		PageID := r.URL.Query().Get(string(PageId))
		intPageID := 0
		var err error
		if PageID != "" {
			intPageID, err = strconv.Atoi(PageID)
			if err != nil {
				http.Error(w, "bad page pequest", http.StatusBadRequest)

				return
			}
		}
		ctx := context.WithValue(r.Context(), PageId, intPageID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
