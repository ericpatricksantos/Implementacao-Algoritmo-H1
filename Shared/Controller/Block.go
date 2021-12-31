package Controller

import (
	"fmt"
	"main/Shared/API"
	"main/Shared/Function"
	"main/Shared/Model"
)

func GetAllLatestBlock(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) []Model.LatestBlock {
	return Function.GetAllLatestBlock(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
}

func SaveLatestBlock(UrlAPI string, LatestBlock string, ConnectionMongoDB string, DataBaseMongo string, Collection string) {
	ultimoBloco := GetLatestBlock(UrlAPI, LatestBlock)
	resposta := Function.SaveLatestBlock(ultimoBloco, ConnectionMongoDB, DataBaseMongo, Collection)
	if resposta {
		fmt.Println("Ultimo Bloco Salvo com Sucesso")
	}
}

func SaveBlock(latestBlock Model.LatestBlock, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	return Function.SaveLatestBlock(latestBlock, ConnectionMongoDB, DataBaseMongo, Collection)
}

func GetLatestBlock(UrlAPI string, LatestBlock string) Model.LatestBlock {
	return API.GetLatestBlock(UrlAPI, LatestBlock)
}

func DeleteLatestBlock(hash string, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) bool {
	return Function.DeleteLatestBlock(hash, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
}
