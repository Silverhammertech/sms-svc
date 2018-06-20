package rest

import (
	"time"
	"runtime"
	"fmt"
	"strings"
	"net/http"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {

	// Always set content type and status code
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	var timestamp = time.Now().UTC().Format(time.RFC850)
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	message := []string{

		fmt.Sprintf("Received Ping on ... %s \n\n", timestamp),
		fmt.Sprintf("Alloc: %d  (bytes allocated and not yet freed)\n", mem.Alloc),
		fmt.Sprintf("Total Alloc: %d  (bytes allocated (even if freed))\n", mem.TotalAlloc),
		fmt.Sprintf("Sys: %d  (bytes obtained from system)\n", mem.Sys),
		fmt.Sprintf("Free Memory: %d  (Sys - Alloc)\n", mem.Sys-mem.Alloc),
		fmt.Sprintf("Lookups: %d  (number of pointer lookups)\n", mem.Lookups),
		fmt.Sprintf("Mallocs: %d  (number of mallocs)\n", mem.Mallocs),
		fmt.Sprintf("Frees: %d  (number of frees)\n", mem.Frees),
	}
	fmt.Fprintf(w, strings.Join(message, ""))
}
