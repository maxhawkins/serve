package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func openURL(url string) {
	try := []string{"xdg-open", "google-chrome", "open"}
	for _, bin := range try {
		if err := exec.Command(bin, url).Run(); err == nil {
			return
		}
	}
	log.Printf("Error opening URL in browser.")
}

func main() {
	var (
		root       = flag.String("root", ".", "Define the root filesystem path")
		shouldOpen = flag.Bool("open", false, "Whether to open a web browser to the running server")
	)
	flag.Parse()

	port := flag.Arg(0)
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	if *shouldOpen {
		go openURL("http://localhost" + addr)
	}

	fmt.Printf("Listening on %s...", addr)
	panic(http.ListenAndServe(addr, http.FileServer(http.Dir(*root))))
}
