package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
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
