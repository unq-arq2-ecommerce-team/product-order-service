package http

import (
	"github.com/hashicorp/go-cleanhttp"
	"net/http"
)

func NewClient() *http.Client {
	return cleanhttp.DefaultPooledClient()
}
