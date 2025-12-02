package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			// TODO: log error
			http.Error(w, "Store fetch timedout", http.StatusNotFound)
			return
		}

		fmt.Fprint(w, data)
	}
}
