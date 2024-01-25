package healthcheck

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Live struct {
	IsLive bool `json:"islive"`
}

var (
	srv = &http.Server{
		Addr:    ":10101",
		Handler: newRouter(),
	}
)

func newRouter() *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/live", statusCheck())
	return mux
}

func statusCheck() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		live := Live{
			IsLive: true,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(live); err != nil {
			panic(err)
		}
	}
}

func ShutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("http server shutdown error %v", err)
	}
}

func StartServer() {
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("fatal http server failed to start: %v", err)
		}
	}
}
