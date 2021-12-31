package Controller

import (
	"main/Shared/Config"
	"main/Shared/Function"
	"main/Shared/Model"
	"main/Shared/Service"
)

func GetConfig() Model.Configuration {
	return Config.GetConfig()
}

func DeleteAll(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) {
	clusters := Function.GetAllCluster(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
	Service.DeleteListCluster(clusters, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
}
