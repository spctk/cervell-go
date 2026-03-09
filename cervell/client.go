package cervell

import "net/http"

// Client ...
type Client struct {
	APIKey     string
	httpClient http.Client
}
