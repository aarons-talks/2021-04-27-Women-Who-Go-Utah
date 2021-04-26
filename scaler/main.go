package main

import (
	"fmt"
	"log"
	"net/http"

	"gcruaaron.dev/pkg/proxy"
)

const scaledTo = "1"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("scaler reporting %s replicas", scaledTo)
		w.Write([]byte(scaledTo))
	})
	addr := fmt.Sprintf("0.0.0.0:%d", proxy.ScalerPort)
	log.Printf("Serving the Scaler on %s with %s replicas", addr, scaledTo)
	log.Fatal(http.ListenAndServe(addr, mux))
}
