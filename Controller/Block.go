package Controller

import (
	"fmt"
	"main/API"
	"main/Function"
	"main/Model"
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

func GetLatestBlock(UrlAPI string, LatestBlock string) Model.LatestBlock {
	return API.GetLatestBlock(UrlAPI, LatestBlock)
}
