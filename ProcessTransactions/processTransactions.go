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
	contadorTxsSalvas := 0
	tempo := 1

	ConnectionMongoDB := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	confirmContinue := false
	for {
		if confirmContinue {
			break
		}
		// Busca todos os Blocos que foram salvos no Database AnalyzedElement na Collection processing
		allblock := Controller.GetAllLatestBlock(ConnectionMongoDB,
			"AnalyzedElement", "processing")
		numeroBlocos := len(allblock)
		if numeroBlocos < 1 {
			fmt.Println("Todos os blocos foram Processadas")
			break
		}
		//indiceInicial := Function.GetIndiceLogIndice(FileLogBlock)
		indiceInicial := 0
		indice := Function.GetIndiceLogIndice(FileLogHash) + 1

		if indice >= len(allblock[indiceInicial].TxIndexes) {
			fmt.Println()
			fmt.Println()
			fmt.Println("-------------------------------------------------------------------")
			fmt.Println("-------------------------------------------------------------------")

			fmt.Println("Todos os endereços foram salvos")
			fmt.Println()
			fmt.Println("Mudando status do bloco de -processing- para -processed-.")
			fmt.Println("Salvando o Bloco: " + allblock[indiceInicial].Hash + " em -processed-.")
			Controller.SaveBlock(allblock[indiceInicial], ConnectionMongoDB,
				"AnalyzedElement", "processed")

			fmt.Println()

			fmt.Println("Deletando o Bloco: " + allblock[indiceInicial].Hash + " de -processing-.")
			Controller.DeleteLatestBlock(allblock[indiceInicial].Hash, ConnectionMongoDB,
				"AnalyzedElement", "processing")

			fmt.Println()

			fmt.Println("Buscando o próximo Bloco de -awaitingProcessing- ")
			block := Controller.GetBlock(ConnectionMongoDB, "AnalyzedElement", "awaitingProcessing")

			fmt.Println()

			fmt.Println("Salvando o Bloco: " + block.Hash + " em -processing-.")
			Controller.SaveBlock(block, ConnectionMongoDB,
				"AnalyzedElement", "processing")

			fmt.Println()

			fmt.Println("Deletando o Bloco: " + block.Hash + " de -awaitingProcessing-.")
			Controller.DeleteLatestBlock(block.Hash, ConnectionMongoDB,
				"AnalyzedElement", "awaitingProcessing")

			fmt.Println()

			temp := []string{strconv.Itoa(0)}
			Function.EscreverTexto(temp, FileLogHash)
			fmt.Println("O indice salvo no indiceTx.txt Atualizado para 0")

			indice = 0

			fmt.Println("-------------------------------------------------------------------")
			fmt.Println("-------------------------------------------------------------------")
			fmt.Println()
			fmt.Println()
			fmt.Println("Dormindo 60 segundos")
			time.Sleep(time.Second * time.Duration(60))
		}

		//Salva todas as Transações dos blocos na Collection Txs
		for contador := indiceInicial; contador < numeroBlocos; contador++ {

			confirm, FinalizaExecucao := Controller.SaveTxs(allblock[contador].TxIndexes, UrlAPI, RawTx, ConnectionMongoDB,
				"Txs", "awaitingProcessing", FileLogHash, indice)

			if confirm {
				fmt.Println("Transação Salva")
				contadorTxsSalvas++
				fmt.Println()
				break
			} else {
				fmt.Println("Não foi salva a transação no MongoDb")
				fmt.Println()
				confirmContinue = FinalizaExecucao
				break
			}

		}

		fmt.Println("Dormindo por ", tempo, "segundos")
		fmt.Println()
		fmt.Println("Txs Salvas ", contadorTxsSalvas)
		fmt.Println()
		time.Sleep(time.Second * time.Duration(tempo))
	}

}
