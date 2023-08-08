package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ankitshah86/jsoniz/core"
	jsonHelper "github.com/ankitshah86/jsoniz/internal/helpers"
)

// StartServer starts the server
func StartServer() {
	// Create a new server
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/", handleQuery)

	// Create a channel to listen for OS signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine so it doesn't block
	go func() {
		fmt.Println("Starting server on port 8080")
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	// Block until we receive OS signal
	if <-signalChan == syscall.SIGINT {
		fmt.Println("Received SIGINT")
	} else {
		fmt.Println("Received SIGTERM")
	}

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline
	srv.Shutdown(ctx)

	fmt.Println("Shutting down")
	os.Exit(0)

}

func handleQuery(w http.ResponseWriter, r *http.Request) {

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// convert body to string
	req := string(body)

	if !jsonHelper.ValidateJson(req) {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// handle the incoming request
	_, err = core.ParseRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Request received", req)

	w.WriteHeader(http.StatusOK)
}
