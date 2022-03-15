package elasticsearch

import (
	"GFBackend/config"
	"context"
	"github.com/olivere/elastic/v7"
	"strconv"
)

var ESClient *elastic.Client
var ctx = context.Background()

func InitES() {
	appConfig := config.AppConfig
	url := "http://" + appConfig.ElasticSearch.IP + ":" + strconv.Itoa(appConfig.ElasticSearch.Port)
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(appConfig.ElasticSearch.Username, appConfig.ElasticSearch.Password),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err.Error())
	}
	ESClient = client

	DataInitialization()
}

func DataInitialization() {
	if !IsIndexExisted("article") {
		mapping := `
		{
			"settings": {},	
			"mappings": {
				"properties": {
					"ID": { 
						"type": "long" 
					},
					"Username": { 
						"type": "keyword" 
					},
					"Title": { 
						"type": "text" 
					},
					"Content": { 
						"type": "text" 
					}
				}	
			}
		}
		`
		if !CreateIndex("article", mapping) {
			panic("Create index \"article\" error")
		}
	}
}
