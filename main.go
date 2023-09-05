package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	verbose  = flag.Bool("v", true, "print verbose output")
	interval = flag.Int("interval", 60, "check interval (in seconds)")

	urls URLs
)

func main() {
	flag.Var(&urls, "url", "Specify URL. This flag can be used multiple times.")
	flag.Parse()

	maybePrint(infoStyle.Render("validating urls..."))
	if err := urls.Validate(); err != nil {
		maybePrint(errorStyle.Render(err.Error()))
		return
	}

	maybePrint(infoStyle.Render("initializing kaffeine..."))

	// Create a ticker
	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	defer ticker.Stop()

	// Create a channel to listen for interrupt signals (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the loop
	for {
		select {
		case <-ticker.C:
			ping()
		case <-sigChan:
			fmt.Println(errorStyle.Render("Program interrupted!"))
			return
		}
	}
}

func maybePrint(s string) {
	if *verbose {
		fmt.Println(s)
	}
}

func ping() {
	for _, u := range urls {
		go func(u string) {
			fmt.Println(genericStyle.Render(fmt.Sprintf("Pinging %q", u)))
			_, err := http.Get(u)
			if err != nil {
				maybePrint(errorStyle.Render(fmt.Sprintf("error pinging %q: %s", u, err)))
				return
			}
			maybePrint(successStyle.Render(fmt.Sprintf("pinged %q successfully", u)))
		}(u)
	}
}
