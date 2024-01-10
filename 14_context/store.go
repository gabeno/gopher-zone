package store

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel() bool
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}
