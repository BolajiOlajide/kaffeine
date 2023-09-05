package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// URLs is a custom flag type that accumulates URL values.
type URLs []string

func (v *URLs) String() string {
	return strings.Join(*v, ", ")
}

func (v *URLs) Set(value string) error {
	*v = append(*v, value)
	return nil
}

func (v URLs) Validate() error {
	if len(v) == 0 {
		return errors.New("no urls provided")
	}

	for _, u := range v {
		parsedURL, err := url.Parse(u)
		if err != nil {
			return err
		}

		// Check if the URL has a valid scheme (http or https)
		if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
			return fmt.Errorf("invalid scheme %q in URL %q", parsedURL.Scheme, u)
		}

		// Check if the URL has a host
		if parsedURL.Host == "" {
			return fmt.Errorf("URL %q missing host", u)
		}
	}

	return nil
}

var (
	verbose = flag.Bool("v", true, "print verbose output")

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
	// Create a ticker that ticks every 5 minutes
	// ticker := time.NewTicker(5 * time.Minute)
	ticker := time.NewTicker(5 * time.Second)
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
			fmt.Println(infoStyle.Render("Program interrupted!"))
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
