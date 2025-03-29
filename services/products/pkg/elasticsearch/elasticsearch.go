package elasticsearch

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

func NewClient(user string, password string, host string, port string) (*elasticsearch.Client, error) {
	url := fmt.Sprintf("http://%s:%s@%s:%s", user, password, host, port)

	cfg := elasticsearch.Config{
		Addresses: []string{url},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}
