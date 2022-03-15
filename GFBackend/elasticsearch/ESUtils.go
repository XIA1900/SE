package elasticsearch

import (
	"GFBackend/controller"
	"GFBackend/logger"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
)

func IsIndexExisted(index string) bool {
	exists, err := ESClient.IndexExists(index).Do(ctx)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return false
	}
	return exists
}

func CreateIndex(index, mapping string) bool {
	createIndex, err := ESClient.CreateIndex(index).Body(mapping).Do(ctx)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return false
	}
	return createIndex.Acknowledged
}

func CreateDocument(articleOfES controller.ArticleOfES) bool {
	_, err := ESClient.Index().Index(indexName).
		Id(strconv.Itoa(articleOfES.ID)).BodyJson(articleOfES).Do(ctx)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return false
	}
	return true
}

func UpdateDocument(articleOfES controller.ArticleOfES) bool {
	_, err := ESClient.Update().Index(indexName).
		Id(strconv.Itoa(articleOfES.ID)).Doc(articleOfES).Do(ctx)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return false
	}
	return true
}

func DeleteDocument(articleOfES controller.ArticleOfES) bool {
	_, err := ESClient.Delete().Index(indexName).
		Id(strconv.Itoa(articleOfES.ID)).Do(ctx)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return false
	}
	return true
}

func MixSearchDocuments(searchingWord string, from, size int) []controller.ArticleOfES {
	multiMatchQuery := elastic.NewMultiMatchQuery(searchingWord, "Username", "Title", "Content")
	searchResult, err1 := ESClient.Search().
		Index(indexName).Query(multiMatchQuery).From(from).Size(size).Pretty(true).Do(ctx)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil
	}
	var articles []controller.ArticleOfES
	if searchResult.Hits.TotalHits.Value > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var article controller.ArticleOfES
			err2 := json.Unmarshal(hit.Source, &article)
			if err2 != nil {
				fmt.Println(err2.Error())
				return nil
			}
			articles = append(articles, article)
		}
	} else {
		return nil
	}
	return articles
}
