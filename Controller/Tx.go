package Controller

import (
	"fmt"
	"main/API"
	"main/Function"
	"main/Model"
	"strconv"
)

// Salva todas as transações de um Block no MongoDB
func SaveTxs(Txs []int, urlAPI string, rawTx string, ConnectionMongoDB string, DataBaseMongo string, Collection string, FileLogHash string) bool {
	indiceInicial := Function.GetIndiceLogIndice(FileLogHash)
	for contador := indiceInicial; contador < len(Txs); contador++ {
		confirm := SaveTx(strconv.Itoa(Txs[contador]), urlAPI, rawTx, ConnectionMongoDB, DataBaseMongo, Collection)
		if !confirm {
			fmt.Println("Não foi salvo a transação ", Txs[contador])
			return false
		}
		fmt.Println("Salvo a Transação")
		temp := []string{strconv.Itoa(contador)}
		Function.EscreverTexto(temp, FileLogHash)
		fmt.Println("Indice Atualizado")
	}

	return true
}

// Salva as Transações no MongoDb
func SaveTx(hash string, urlAPI string, rawTx string, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	tx := GetTx(hash, urlAPI, rawTx)
	if len(tx.Hash) > 0 {
		resposta := Function.SaveTx(tx, ConnectionMongoDB, DataBaseMongo, Collection)
		if resposta {
			fmt.Println("Transacao Salva com Sucesso")
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
