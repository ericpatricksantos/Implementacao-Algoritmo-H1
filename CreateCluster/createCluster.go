package main

import (
	"fmt"
	"main/Shared/Function"
	"main/Shared/Service"
	"time"
)

func main() {
	countCreateClusters := 0
	escolhaConexao := 1
	ConnectionMongoDB := []string{
		"mongodb+srv://ericpatrick:9858epJusd@cluster0.cieqi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
	}
	DataBaseTx := "Txs"
	DataBaseCluster := "Cluster"

	ColTxProcessed := "processed"
	ColTxProcessing := "processing"
	ColTxAwaitingProcessing := "awaitingProcessing"

	ColClusterProcessed := "processed"

	EncerraExecucao := false
	tempo := 1
	for {
		if EncerraExecucao {
			break
		}
		fmt.Println("** Buscando Transações que estão em processamento **")
		TxsProcessing := Function.GetTxMongoDB(ConnectionMongoDB[escolhaConexao], DataBaseTx, ColTxProcessing)

		if len(TxsProcessing.Hash) < 1 && len(TxsProcessing.Inputs) < 1 {
			fmt.Println("---------------------- Não existem Transações que estão em processamento no momento ------")
			fmt.Println("---------------------- Buscando uma Transação que esta Aguardando processamento ----------")
			fmt.Println("------------------------------------------------------------------------------------------")
			fmt.Println()
			TxsAwaitingProcessing := Function.GetTxMongoDB(ConnectionMongoDB[escolhaConexao], DataBaseTx, ColTxAwaitingProcessing)

			if len(TxsAwaitingProcessing.Hash) < 1 && len(TxsAwaitingProcessing.Inputs) < 1 {
				fmt.Println("---------------------- Não tem Transação aguardando processamento ------------------------")
				fmt.Println("****************************** FIM *******************************************************")
				fmt.Println("*********************** Clusters Criados ", countCreateClusters, " ***********************************************")
				EncerraExecucao = true
				break
			}

			mudou := Function.MudancaStatusTx(TxsAwaitingProcessing, ConnectionMongoDB[escolhaConexao], DataBaseTx, ColTxAwaitingProcessing, ColTxProcessing)

			if mudou {
				fmt.Println("Mudança de status concluida ", ColTxAwaitingProcessing, " >> ", ColTxProcessing)
			} else {
				fmt.Println("Falha na mudança de status ", ColTxAwaitingProcessing, " >> ", ColTxProcessing)
				EncerraExecucao = true
				break
			}

			fmt.Println()
		} else {
			fmt.Println()
			fmt.Println("*********************************** INICIO ***********************************************")
			for {
				fmt.Println("** Existem Transações em Processamento **")
				fmt.Println("** Criando Clusters **")
				confirmCluster, FinalizaExecucao := Service.CreateCluster(ConnectionMongoDB[escolhaConexao], DataBaseTx, ColTxProcessing, DataBaseCluster, ColClusterProcessed)
				if confirmCluster {
					countCreateClusters++
					fmt.Println("---------------------- Clusters criados com Sucesso --------------------------------------")
					fmt.Println("---------------------- Mudança de status: -processing- --> -processed- -------------------")
					mudou := Function.MudancaStatusTx(TxsProcessing, ConnectionMongoDB[escolhaConexao], DataBaseTx, ColTxProcessing, ColTxProcessed)

					if mudou {
						fmt.Println("Mudança de status concluida ", ColTxProcessing, " >> ", ColTxProcessed)
					} else {
						fmt.Println("Falha na mudança de status ", ColTxProcessing, " >> ", ColTxProcessed)
						EncerraExecucao = true
						break
					}

					fmt.Println()
					break
				} else if !FinalizaExecucao {
					fmt.Println("Clusters nao foram criados")
					break
				} else {
					fmt.Println("---------------------- Clusters nao foram criados ----------------------------------------")
					EncerraExecucao = FinalizaExecucao
					fmt.Println()
					break
				}

			}
			fmt.Println("****************************** FIM *******************************************************")
			fmt.Println("*********************** Clusters Criados ", countCreateClusters, " ***********************************************")
		}
		fmt.Println("Dormindo ", tempo, " segundos")
		time.Sleep(time.Second * time.Duration(tempo))
	}
}
