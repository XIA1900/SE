package elasticsearch

import (
	"GFBackend/logger"
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
