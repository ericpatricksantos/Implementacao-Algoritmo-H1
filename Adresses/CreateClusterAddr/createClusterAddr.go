package main

import (
	"fmt"
	"main/Shared/Service"
	"time"
)

func main() {
	countCreateClusters := 0
	escolhaConexao := 1
	newAddrAnalise := false
	tempo := 1
	ConnectionMongoDB := []string{
		"mongodb+srv://ericpatrick:9858epJusd@cluster0.cieqi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
	}
	DataBaseAddr := "Adresses"
	DataBaseCluster := "ClusterAdresses"

	processed := "processed"
	processedCluster := "processedCluster"
	processedAddrAnalise := "processedEnderecosEmAnalise"
	processedAddrAnaliseCluster := "processedEnderecosEmAnaliseCluster"

	CollectionCluster := "processed"

	encerraExecucao := false
	for {
		if encerraExecucao {
			break
		}
		fmt.Println()
		fmt.Println("*********************************** INICIO ***********************************************")
		confirmCluster, FinalizaExecucao := Service.CreateClustersAddr(
			ConnectionMongoDB[escolhaConexao], DataBaseCluster, CollectionCluster,
			DataBaseAddr, processed, processedCluster,
			processedAddrAnalise, processedAddrAnaliseCluster, newAddrAnalise)
		if confirmCluster {
			countCreateClusters++
			fmt.Println("Cluster criado com Sucesso")
		} else if !confirmCluster && !FinalizaExecucao {
			fmt.Println("Cluster foi criado anteriormente")
		} else {
			fmt.Println("Erro: Finalizando Execução")
			encerraExecucao = FinalizaExecucao
		}
		fmt.Println("****************************** FIM *******************************************************")
		fmt.Println("*********************** Clusters Criados ", countCreateClusters, " ********************************************")
		fmt.Println("*********************** Dormindo ", tempo, " segundos ********************************************")
		time.Sleep(time.Second * time.Duration(tempo))
	}
}
