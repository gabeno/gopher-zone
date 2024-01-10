package store

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() bool {
	s.cancelled = true
	return s.cancelled
}

func TestServer(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		data := "Hello, world"
		svr := Server(&SpyStore{response: data, cancelled: false})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s want %s", response.Body.String(), data)
		}
	})

	t.Run("tell store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, cancelled: false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("store was not told to cancel")
		}
	})
}
