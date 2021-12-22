package Controller

import "main/Service"

func CreateCluster(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string,
	CollectionSalvaDados string) {
	Service.CreateCluster(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados,
		CollectionSalvaDados)
}

func H1(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) {
	Service.H1(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
}
