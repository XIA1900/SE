package elasticsearch

import (
	"GFBackend/config"
	"GFBackend/logger"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"strconv"
)

var ES *elasticsearch.Client

func InitES() {
	appConfig := config.AppConfig

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://" + appConfig.ElasticSearch.IP + ":" + strconv.Itoa(appConfig.ElasticSearch.Port),
		},
		Username: appConfig.ElasticSearch.Username,
		Password: appConfig.ElasticSearch.Password,
	}
	newES, err := elasticsearch.NewClient(cfg)
	if err != nil {
		logger.AppLogger.Error(fmt.Sprintf("Create ElasticSearch Client Error: %s", err))
	}
	ES = newES
}
