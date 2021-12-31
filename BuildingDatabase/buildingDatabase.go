package main

import (
	"fmt"
	"main/Shared/Controller"
	"time"
)

var UrlAPI string = Controller.GetConfig().UrlAPI[0] // "https://blockchain.info"
var LatestBlock string = Controller.GetConfig().LatestBlock

func main() {
	ConnectionMongoDB := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	for {
		// Salva o ultimo Bloco gerado na Blockchain na Collection LatestBlock
		Controller.SaveLatestBlock(UrlAPI, LatestBlock, ConnectionMongoDB,
			"AnalyzedElement", "awaitingProcessing")

		fmt.Println("Dormindo por 30 minutos")
		time.Sleep(time.Minute * 30)
	}
}
