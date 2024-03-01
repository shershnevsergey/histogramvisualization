package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/shershnevsergey/histogramvisualization/utils/constants"
	"github.com/shershnevsergey/histogramvisualization/utils/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go signal.HandleShutdown(cancel)

	fmt.Printf("Sending HTTP requests to %s in an infinite loop...\n", constants.TargetURL)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("generator was stopped")
			return
		default:
			sendHTTPRequest()
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// Function for sending HTTP GET request.
func sendHTTPRequest() {
	resp, err := http.Get(constants.TargetURL)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("HTTP GET request sent to %s. Status Code: %d\n", constants.TargetURL, resp.StatusCode)
}
