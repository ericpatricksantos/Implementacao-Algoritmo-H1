package main

import (
	"fmt"
	"main/Shared/Function"
	"main/Shared/Service"
)

func main() {
	countCreateClusters := 0
	escolhaConexao := 0
	ConnectionMongoDB := []string{
		"mongodb+srv://ericpatrick:9858epJusd@cluster0.cieqi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
	}
	DataBaseTx := "blockchain"
	DataBaseCluster := "blockchain"

	ColTxProcessed := "Txprocessed"
	ColTxProcessing := "processing"
	ColTxAwaitingProcessing := "awaitingProcessing"

	ColClusterProcessed := "processed"

	EncerraExecucao := false

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
			fmt.Println("---------------------- Salvando Transação na Collection Processing -----------------------")
			confirm := Function.SalveTxMongoDB(TxsAwaitingProcessing, ConnectionMongoDB[escolhaConexao], DataBaseTx, ColTxProcessing)
			if confirm {
				fmt.Println("---------------------- Transação salva com sucesso na Collection Processing --------------")

				con := Function.DeleteTxMongo(TxsAwaitingProcessing.Hash, ConnectionMongoDB[escolhaConexao], DataBaseTx, ColTxAwaitingProcessing)
				if con {
					fmt.Println("---------------------- Transação deletada da Collection Awaiting Processing --------------")
					fmt.Println("---------------------- Transação Pronta para ser Processada em Cluster -------------------")
				} else {
					fmt.Println("---------------------- Transação nao foi deletada da Collection Awaiting Processing ------")
					EncerraExecucao = true
					break
				}
			} else {
				fmt.Println("---------------------- Transação nao foi salva na Collection Processing ------------------")
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
					fmt.Println("---------------------- Salvando a transação na Collection processed ----------------------")
					conf := Function.SalveTxMongoDB(TxsProcessing, ConnectionMongoDB[escolhaConexao], DataBaseTx, ColTxProcessed)
					if conf {
						fmt.Println("---------------------- Salvo com Sucesso -------------------------------------------------")
						fmt.Println("---------------------- Deletando a transação da Collection processing --------------------")
						con := Function.DeleteTxMongo(TxsProcessing.Hash, ConnectionMongoDB[escolhaConexao], DataBaseTx, ColTxProcessing)
						if con {
							fmt.Println("---------------------- Deletado com sucesso ----------------------------------------------")
							fmt.Println("---------------------- Mudanda de Status realizada com Sucesso ---------------------------")
						} else {
							fmt.Println("---------------------- Nao foi deletado com sucesso --------------------------------------")
							EncerraExecucao = true
						}
					} else {
						fmt.Println("---------------------- Não foi salvo com sucesso -----------------------------------------")
						EncerraExecucao = true
					}
					fmt.Println()
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
	}
}
