package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/sse", sseHandler)
	http.ListenAndServe(":8080", nil)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Sending Server Sent Events")
	ctx := r.Context()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for i := range 10 {
		//https://web.dev/articles/eventsource-basics#event_stream_format
		// SSE Needs events to be of below format
		fmt.Fprintf(w, fmt.Sprintf("data: %d", i))
		time.Sleep(1 * time.Second)
		w.(http.Flusher).Flush()
	}
	select {
	case <-ctx.Done():
		return
	default:
	}
}
