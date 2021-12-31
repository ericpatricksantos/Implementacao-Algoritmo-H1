package Controller

import (
	"fmt"
	"main/Shared/API"
	"main/Shared/Function"
	"main/Shared/Model"
	"strconv"
)

// Salva todas as transações de um Block no MongoDB
func SaveTxs(Txs []int, urlAPI string, rawTx string, ConnectionMongoDB string, DataBaseMongo string, Collection string, FileLogHash string, indiceInicial int) bool {
	//indiceInicial := Function.GetIndiceLogIndice(FileLogHash) + 1
	for contador := indiceInicial; contador < len(Txs); contador++ {
		confirm := SaveTx(strconv.Itoa(Txs[contador]), urlAPI, rawTx, ConnectionMongoDB, DataBaseMongo, Collection)
		if !confirm {
			fmt.Println("Não foi salvo a transação ", Txs[contador])
			return false
		} else {
			fmt.Println("Salvo a Transação: Nº ", Txs[contador])
			fmt.Println("Indice Atualizado para ", contador)
			temp := []string{strconv.Itoa(contador)}
			Function.EscreverTexto(temp, FileLogHash)
			return true
		}
		//fmt.Println("Salvo a Transação: Nº ", Txs[contador])
		//temp := []string{strconv.Itoa(contador)}
		//Function.EscreverTexto(temp, FileLogHash)
		//fmt.Println("Indice Atualizado para ", contador)
		//
		//time.Sleep(time.Minute * time.Duration(tempo))
	}

	return true
}

// Salva as Transações no MongoDb
func SaveTx(hash string, urlAPI string, rawTx string, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	tx := GetTx(hash, urlAPI, rawTx)
	if len(tx.Hash) > 0 {
		resposta := Function.SaveTx(tx, ConnectionMongoDB, DataBaseMongo, Collection)
		if resposta {
			return true
		} else {
			return false
		}
	} else {
		fmt.Println("O campo Hash está vazio, por isso nao foi salvo")
		return false
	}
	return false
}

// Get Transação da API da Blockchain
func GetTx(hash string, urlAPI string, rawTx string) Model.Transaction {
	return API.GetTransaction(hash, urlAPI, rawTx)
}
