package main

import (
	"fmt"
	"main/Controller"
	"main/Function"
	"strconv"
)

var ConnectionMongoDB string = Controller.GetConfig().ConnectionMongoDB[0] //"connection string into your application code"
var DataBase string = Controller.GetConfig().DataBase[0]                   //blockchain

var UrlAPI string = Controller.GetConfig().UrlAPI[0] // "https://blockchain.info"

var LatestBlock string = Controller.GetConfig().LatestBlock
var RawTx string = Controller.GetConfig().RawTx

var CollectionLatestBlock string = Controller.GetConfig().Collection[0]
var CollectionTesteTxs string = Controller.GetConfig().Collection[1]
var CollectionTesteCluster string = Controller.GetConfig().Collection[2]
var CollectionCluster string = Controller.GetConfig().Collection[3]
var CollectionTxs string = Controller.GetConfig().Collection[4]

var FileLogHash string = Controller.GetConfig().FileLog[0]
var FileLogBlock string = Controller.GetConfig().FileLog[1]

func main() {
	// Salva o ultimo Bloco gerado na Blockchain na Collection LatestBlock
	Controller.SaveLatestBlock(UrlAPI, LatestBlock, ConnectionMongoDB, DataBase, CollectionLatestBlock)

	// Busca todos os Blocos que foram salvos na Collection LatestBlock
	allblock := Controller.GetAllLatestBlock(ConnectionMongoDB, DataBase, CollectionLatestBlock)
	indiceInicial := Function.GetIndiceLogIndice(FileLogBlock)

	//Salva todas as Transações dos blocos na Collection Txs
	for contador := indiceInicial; contador < len(allblock); contador++ {

		confirm := Controller.SaveTxs(allblock[contador].TxIndexes, UrlAPI, RawTx, ConnectionMongoDB, DataBase, CollectionTxs, FileLogHash)

		if confirm {
			fmt.Println("Todas as transações do Block foram salvas no MongoDb")

			temp := []string{strconv.Itoa(contador)}
			Function.EscreverTexto(temp, FileLogBlock)

			fmt.Println("Indice do Bloco Atualizado")
		} else {
			fmt.Println("Não foram salvas todas as transações no MongoDb")
			break
		}

	}

	//Recuperando as Transações da Collection Txs e cria Cluster
	Controller.CreateCluster(ConnectionMongoDB, DataBase, CollectionTxs, CollectionCluster)

	////Reorganiza os Cluster utilizando o algoritmo H1
	Controller.H1(ConnectionMongoDB, DataBase, CollectionCluster)

}
