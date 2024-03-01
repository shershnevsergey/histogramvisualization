package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	metricsprometheus "github.com/slok/go-http-metrics/metrics/prometheus"
	metricsmiddleware "github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"

	"github.com/shershnevsergey/histogramvisualization/utils/constants"
	"github.com/shershnevsergey/histogramvisualization/utils/rand"
	"github.com/shershnevsergey/histogramvisualization/utils/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go signal.HandleShutdown(cancel)

	if err := runServer(ctx); err != nil {
		fmt.Println(err)
	}

	fmt.Println("application was stopped")
}

func runServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/payload", payloadHandler) // Define payload endpoint.
	mux.Handle("/metrics", promhttp.Handler()) // Define prometheus metrics endpoint.

	server := http.Server{
		Addr: fmt.Sprintf(":%d", constants.Port),
		Handler: std.Handler(
			"",
			metricsmiddleware.New(metricsmiddleware.Config{
				Recorder: metricsprometheus.NewRecorder(metricsprometheus.Config{
					DurationBuckets: constants.HistogramBuckets,
				}),
			}),
			mux,
		),
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	fmt.Printf("HTTP server is running on %d port\n", constants.Port)

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		fmt.Println("shutting down the server...")

		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to shutdown the server: %w\n", err)
		}
	}

	return nil
}

func payloadHandler(response http.ResponseWriter, _ *http.Request) {
	now := time.Now()
	delay := rand.GetDelayDuration()
	time.Sleep(delay)

	fmt.Fprintf(response, "Request took %v with the delay %v", time.Since(now), delay)
}
