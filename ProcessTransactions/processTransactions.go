package main

import (
	"fmt"
	"main/Shared/Controller"
	"main/Shared/Function"
	"strconv"
	"time"
)

var UrlAPI string = Controller.GetConfig().UrlAPI[0]
var RawTx string = Controller.GetConfig().RawTx

var FileLogHash string = Controller.GetConfig().FileLog[0]
var FileLogBlock string = Controller.GetConfig().FileLog[1]

func main() {
	tempo := 1

	ConnectionMongoDB := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	confirmContinue := false
	for {
		if confirmContinue {
			break
		}
		// Busca todos os Blocos que foram salvos na Collection LatestBlock
		allblock := Controller.GetAllLatestBlock(ConnectionMongoDB,
			"AnalyzedElement", "processing")

		//indiceInicial := Function.GetIndiceLogIndice(FileLogBlock)
		indiceInicial := 0
		indice := Function.GetIndiceLogIndice(FileLogHash) + 1

		if indice >= len(allblock[indiceInicial].TxIndexes) {
			fmt.Println("Todos os endereços foram salvos")

			Controller.SaveBlock(allblock[indiceInicial], ConnectionMongoDB,
				"AnalyzedElement", "processed")

			Controller.DeleteLatestBlock(allblock[indiceInicial].Hash, ConnectionMongoDB,
				"AnalyzedElement", "processing")

			block := Controller.GetBlock(ConnectionMongoDB, "AnalyzedElement", "awaitingProcessing")

			Controller.SaveBlock(block, ConnectionMongoDB,
				"AnalyzedElement", "processing")

			Controller.DeleteLatestBlock(block.Hash, ConnectionMongoDB,
				"AnalyzedElement", "awaitingProcessing")

			temp := []string{strconv.Itoa(0)}
			Function.EscreverTexto(temp, FileLogHash)
			fmt.Println("Indice do Tx Atualizado para 0")

		}

		//Salva todas as Transações dos blocos na Collection Txs
		for contador := indiceInicial; contador < len(allblock); contador++ {

			confirm := Controller.SaveTxs(allblock[contador].TxIndexes, UrlAPI, RawTx, ConnectionMongoDB,
				"Txs", "awaitingProcessing", FileLogHash, indice)

			if confirm {
				fmt.Println("Transação Salva")
				break
			} else {
				fmt.Println("Não foi salva a transação no MongoDb")
				confirmContinue = true
				break
			}

		}

		fmt.Println("Dormindo por ", tempo, "segundos")
		time.Sleep(time.Second * time.Duration(tempo))
	}

}
