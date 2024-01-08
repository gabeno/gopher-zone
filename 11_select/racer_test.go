package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers returning url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("timeouts if server does not respond in 10 seconds", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Millisecond)

		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
