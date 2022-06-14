package elasticsearch

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/kristiansantos/learning/src/config/initializer"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
)

type ElasticsearchSetup struct {
	Logger              logger.ILoggerProvider
	Config              initializer.Application
	ElasticsearchClient *elasticsearch.Client
}

func Connection(app initializer.Application, log logger.ILoggerProvider) error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			app.Elasticsearch.Url,
		},
		Username: app.Elasticsearch.User,
		Password: app.Elasticsearch.Pass,
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
