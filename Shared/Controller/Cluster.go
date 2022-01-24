package Controller

import (
	"main/Shared/Service"
)

func CreateCluster(ConnectionMongoDB, DataBaseTx, CollectionTx, DataBaseCluster,
	CollectionCluster string) (bool, bool) {
	return Service.CreateCluster(ConnectionMongoDB, DataBaseTx, CollectionTx, DataBaseCluster,
		CollectionCluster)
}

func H1(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) bool {
	return Service.H1(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
}
