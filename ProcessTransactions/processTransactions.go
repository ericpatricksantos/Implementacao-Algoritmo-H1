package main

import (
	"fmt"
	"main/Shared/Controller"
	"main/Shared/Function"
	"time"
)

var UrlAPI string = Controller.GetConfig().UrlAPI[0]
var RawTx string = Controller.GetConfig().RawTx

var FileLogHash string = Controller.GetConfig().FileLog[0]
var FileLogBlock string = Controller.GetConfig().FileLog[1]

func main() {
	tempo := 2

	ConnectionMongoDB := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	for {
		// Busca todos os Blocos que foram salvos na Collection LatestBlock
		allblock := Controller.GetAllLatestBlock(ConnectionMongoDB,
			"AnalyzedElement", "processing")

		//indiceInicial := Function.GetIndiceLogIndice(FileLogBlock)
		indiceInicial := 0
		indice := Function.GetIndiceLogIndice(FileLogHash) + 1

		if indice > len(allblock[indiceInicial].TxIndexes) {
			fmt.Println("Todos os endereços foram salvos")

			//Controller.SaveBlock(allblock[contador], ConnectionMongoDB,
			//	"AnalyzedElement", "processed")
			//
			//Controller.DeleteLatestBlock(allblock[contador].Hash, ConnectionMongoDB,
			//	"AnalyzedElement", "processing")

			//temp := []string{strconv.Itoa(contador)}
			//Function.EscreverTexto(temp, FileLogBlock)
			//fmt.Println("Indice do Bloco Atualizado")
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
				break
			}

		}

		fmt.Println("Dormindo por ", tempo, "minutos")
		time.Sleep(time.Minute * time.Duration(tempo))
	}

}
