package gracehttpserverrunner

import (
	"context"
	"net/http"
	"time"
)

// HTTPServerRunner runs a HTTP server.
type HTTPServerRunner struct {
	// Server is the HTTP server to serve.
	Server *http.Server
	// ShutdownTimeout is the maximum time to allow for a graceful shutdown.
	ShutdownTimeout time.Duration
	shutdownDoneCh  chan struct{}
}

// Run starts a http server.
func (r *HTTPServerRunner) Run(ctx context.Context) error {
	go r.handleCtxDone(ctx)

	err := r.Server.ListenAndServe()

	// ErrServerClosed is returned when we call Server.Close() from handleCtxDone
	// so we shouldn't treat it as an error.
	if err == http.ErrServerClosed {

		// ErrServerClosed is immediately returned when Shutdown is called on the Server, but we need to
		// wait until Shutdown has returned.
		if r.shutdownDoneCh != nil {
			<-r.shutdownDoneCh
			r.shutdownDoneCh = nil
		}

		return nil
	}

	return err
}

func (r *HTTPServerRunner) handleCtxDone(ctx context.Context) {
	<-ctx.Done()
	if r.Server != nil {
		if r.ShutdownTimeout == 0 {
			r.ShutdownTimeout = time.Second * 10
		}
		// We need to wait for Shutdown to return before we can return from Run, but ListenAndServe
		// returns immediately when we issue Shutdown.
		// This channel allows us to delay to return of Run.
		r.shutdownDoneCh = make(chan struct{})

		shutdownCtx, _ := context.WithTimeout(context.Background(), r.ShutdownTimeout)
		r.Server.Shutdown(shutdownCtx)

		close(r.shutdownDoneCh)
	}
}
