package query

import (
	"net/http"

	es "github.com/signalfx/signalfx-agent/pkg/monitors/elasticsearch/client"
)

type ESQueryHTTPClient struct {
	esClient *es.ESClient
}

// NewESQueryClient creates a new ESQueryHTTPClient
func NewESQueryClient(host string, port string, scheme string, client *http.Client) ESQueryHTTPClient {
	return ESQueryHTTPClient{
		esClient: &es.ESClient{
			Scheme:     scheme,
			Host:       host,
			Port:       port,
			HTTPClient: client,
		},
	}
}
