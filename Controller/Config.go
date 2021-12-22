package Controller

import (
	"main/Config"
	"main/Function"
	"main/Model"
	"main/Service"
)

func GetConfig() Model.Configuration {
	return Config.GetConfig()
}

func DeleteAll(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) {
	clusters := Function.GetAllCluster(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
	Service.DeleteListCluster(clusters, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
}
