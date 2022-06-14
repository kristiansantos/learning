package elasticsearch

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/kristiansantos/learning/src/config/initializers"
	"github.com/kristiansantos/learning/src/shared/providers/logger"
)

type ElasticsearchSetup struct {
	Logger              logger.ILoggerProvider
	Brand               string
	Config              initializers.Initializer
	ElasticsearchClient *elasticsearch.Client
}

func Connection(initializer initializers.Initializer, log logger.ILoggerProvider) error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			initializer.Elasticsearch.Url,
		},
		Username: initializer.Elasticsearch.User,
		Password: initializer.Elasticsearch.Pass,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	_, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Error("elasticsearch.Connection", fmt.Sprintf("Error creating the elasticsearch client: %v", err))
	}

	return nil
}
