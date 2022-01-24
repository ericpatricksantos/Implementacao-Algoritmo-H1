package main

import (
	"fmt"
	"main/Shared/Function"
	"time"
)

func main() {
	ConnectionMongoDB := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	DatabaseEnderecosAnalise := "EnderecosEmAnalise"
	DatabaseAdresses := "Adresses"
	AwaitingProcessing := "awaitingProcessing"
	AwaitingProcessingEnderecosEmAnalise := "awaitingProcessingEnderecosEmAnalise"
	CollectionAnaliseProcessed := "processed"
	urlAPI := "https://blockchain.info"
	RawAddr := "/rawaddr/"
	for {
		confirm := Function.ProcessAdressesAnalysis(ConnectionMongoDB, DatabaseEnderecosAnalise, AwaitingProcessing, CollectionAnaliseProcessed,
			DatabaseAdresses, AwaitingProcessingEnderecosEmAnalise, urlAPI, RawAddr)

		if confirm {
			fmt.Println("Endereco Salvo com Sucesso")
		} else {
			fmt.Println("Falha ao Salvar Endereco")
			break
		}
		time.Sleep(time.Second * time.Duration(60))
		fmt.Println("Dormindo 60 segundos")
	}
}
