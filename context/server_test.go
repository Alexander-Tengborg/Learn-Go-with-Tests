package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("simple uncancelled request that returns data from the store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		server := Server(store)
		
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		
		server.ServeHTTP(response, request)
		
		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}

		if store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		server := Server(store)
		
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(time.Millisecond * 5, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()
		
		server.ServeHTTP(response, request)
		
		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
}

type SpyStore struct {
	response string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(time.Millisecond * 100)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}
