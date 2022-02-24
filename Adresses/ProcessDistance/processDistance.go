package main

import (
	"fmt"
	"main/Shared/Function"
	"time"
)

func main() {
	ConnectionMongoDB := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	DataBaseDistancia := "Distancia"
	DataBaseAddr := "Adresses"
	awaitingProcessing := "awaitingProcessing"
	processedAddr := "processed"
	processsedDistance := "processed"
	processingDistance := "processing"
	urlAPI := "https://blockchain.info"
	RawAddr := "/rawaddr/"
	MultiAddr := "/multiaddr?active="
	EncerraExecucao := false
	tempo := 6
	tempoAux := 6
	count := 0

	for {
		if EncerraExecucao || tempo > 200 {
			break
		}
		fmt.Println("Verificando se existem distancias em processing")
		distancia := Function.GetDistanciaMongoDB(ConnectionMongoDB, DataBaseDistancia, processingDistance)

		if len(distancia.AddressInput) < 1 {
			fmt.Println("Colocando distancia para serem processadas")
			distancia = Function.GetDistanciaMongoDB(ConnectionMongoDB, DataBaseDistancia, awaitingProcessing)

			if len(distancia.AddressInput) < 1 {
				fmt.Println()
				fmt.Println("Não tem nenhum distância aguardando processamento")
				break
			}

			salvo, _, existente := Function.SalveDistanciaMongoDB(distancia, ConnectionMongoDB, DataBaseDistancia, processingDistance)

			if !salvo && !existente {
				fmt.Println()
				fmt.Println("Falha na mudança de status:")
				fmt.Println("Não foi salvo a distancia na collection processing")
				fmt.Println()
				EncerraExecucao = true
				break
			} else if !salvo && existente {
				fmt.Println("Não foi salva a distancia, porque a distancia ja esta salva")
			}
			deleteDistancia := Function.DeleteDistanciaMongo(distancia.AddressInput, ConnectionMongoDB, DataBaseDistancia, awaitingProcessing)

			if deleteDistancia {
				fmt.Println("Deletado com sucesso da distancia na collection awaitingProcessing")
				fmt.Println("Mudança de status concluida awaitingProcessing -> processing")
			} else {
				fmt.Println("Falha na deleção da distancia na collection awaitingProcessing")
				EncerraExecucao = true
				break
			}

			fmt.Println()
		} else {

			confirm, finalizaExecucao, erro := Function.ProcessDistance(ConnectionMongoDB, DataBaseAddr, awaitingProcessing, processedAddr,
				DataBaseDistancia, processingDistance, processsedDistance, urlAPI, RawAddr, MultiAddr)

			if confirm && !finalizaExecucao {
				fmt.Println("Distancia Processada")
				count++
			} else if !confirm && !finalizaExecucao && !erro {
				fmt.Println("O endereço foi salvo anteriormente")
				EncerraExecucao = finalizaExecucao
			} else if !erro {
				fmt.Println("Distancia nao foi processada")
				EncerraExecucao = true
				break
			}
			if erro {
				tempo = tempo + 30
			} else {
				tempo = tempoAux
			}
			fmt.Println()
			fmt.Println("Enderecos Salvos: ", count)
			fmt.Println("Dormindo ", tempo, " segundos")
			time.Sleep(time.Second * time.Duration(tempo))
			fmt.Println()
		}

	}
}
