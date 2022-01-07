package main

import (
	"fmt"
	"main/Shared/Controller"
	"strconv"
	"time"
)

var UrlAPI string = Controller.GetConfig().UrlAPI[0] // "https://blockchain.info"
var LatestBlock string = Controller.GetConfig().LatestBlock

func main() {
	ConnectionMongoDB := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	for {
		tempo := 10
		// Salva o ultimo Bloco gerado na Blockchain na Collection LatestBlock
		conf := Controller.SaveLatestBlock(UrlAPI, LatestBlock, ConnectionMongoDB,
			"AnalyzedElement", "awaitingProcessing")
		if !conf {
			tempo = tempo + 5
		} else {
			tempo = 15
		}
		fmt.Println("Dormindo por " + strconv.Itoa(tempo) + " minutos")
		time.Sleep(time.Minute * time.Duration(tempo))
	}
}
