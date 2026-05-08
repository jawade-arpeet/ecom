package client

import (
	"ecom/internal/config"

	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

type ElasticSearchClient struct {
	es *elasticsearch.TypedClient
}

func newElasticSearchClient(cfg *config.ElasticSearchConfig) (*ElasticSearchClient, error) {
	esCfg := elasticsearch.Config{
		Addresses: []string{cfg.Address},
		Username:  cfg.Username,
		Password:  cfg.Password,
	}

	es, err := elasticsearch.NewTypedClient(esCfg)
	if err != nil {
		zap.L().Error(
			"failed to create elastic search client",
			zap.Error(err),
		)
		return nil, err
	}

	zap.L().Info("successfully connected to elastic search")

	return &ElasticSearchClient{es: es}, nil
}
