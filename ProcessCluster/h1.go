package main

import (
	"fmt"
	"main/Shared/Controller"
)

func main() {
	execucoes := 1
	i := 0
	escolhaConexao := 1
	ConnectionMongoDB := []string{
		"mongodb+srv://ericpatrick:9858epJusd@cluster0.cieqi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
	}
	DataBaseCluster := "Cluster"
	ColClusterProcessed := "processed"
	for {
		if i >= execucoes {
			break
		}
		fmt.Println()
		fmt.Println("Aplicando Algoritmo H1")

		confirm := Controller.H1(ConnectionMongoDB[escolhaConexao], DataBaseCluster, ColClusterProcessed)

		if confirm {
			fmt.Println("Execução finalizada com Sucesso")
		} else {
			fmt.Println("Execução finalizada com erro")
		}
		fmt.Println()

		i++

	}
}
