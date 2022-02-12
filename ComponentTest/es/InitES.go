package es

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var es *elasticsearch.Client

func InitES() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://167.71.166.120:9200",
		},
		Username: "elastic",
		Password: "gfes22",
	}
	newES, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	es = newES
}

func GetES() *elasticsearch.Client {
	return es
}
