package main

import (
	"fmt"
	"main/Controller"
	"main/Function"
	"strconv"
)

func main() {
	// Salva o ultimo Bloco gerado na Blockchain na Collection LatestBlock
	Controller.SaveLatestBlock(UrlAPI, LatestBlock, ConnectionMongoDB, DataBase, CollectionLatestBlock)

	// Busca todos os Blocos que foram salvos na Collection LatestBlock
	allblock := Controller.GetAllLatestBlock(ConnectionMongoDB, DataBase, CollectionLatestBlock)
	indiceInicial := Function.GetIndiceLogIndice(FileLogBlock)

	//Salva todas as Transações dos blocos na Collection Txs
	for contador := indiceInicial; contador < len(allblock); contador++ {
		Controller.SaveTxs(allblock[contador].TxIndexes, UrlAPI, RawTx, ConnectionMongoDB, DataBase, CollectionTxs, FileLogHash)

		fmt.Println("Salvo todas as transações do Block")

		temp := []string{strconv.Itoa(contador)}
		Function.EscreverTexto(temp, FileLogBlock)

		fmt.Println("Indice do Bloco Atualizado")
	}

	//Recuperando as Transações da Collection Txs e cria Cluster
	Controller.CreateCluster()

	//Reorganiza os Cluster utilizando o algoritmo H1
	Controller.H1()
}
