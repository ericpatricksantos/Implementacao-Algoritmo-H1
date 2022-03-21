package main

import (
	"fmt"
	"main/Shared/Service"
)

func a() {
	NoCheckNextAddr := false
	IgnoraCluster := 100

	encerraExecucao := false
	escolhaConexao := 1
	ConnectionMongoDB := []string{
		"mongodb+srv://ericpatrick:9858epJusd@cluster0.cieqi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
	}
	DataBaseCluster := "teste"
	ColClusterProcessed := "algoritmo"
	for {

		if encerraExecucao {
			break
		}
		fmt.Println()
		fmt.Println("Aplicando Algoritmo H1")

		confirm, encerra, executeAll := Service.H1(ConnectionMongoDB[escolhaConexao], DataBaseCluster, ColClusterProcessed, IgnoraCluster, NoCheckNextAddr)

		if confirm && executeAll {
			fmt.Println("Execução finalizada com Sucesso")
			encerraExecucao = executeAll
		} else if !confirm && encerra {
			encerraExecucao = encerra
			fmt.Println("Execução finalizada com erro")
		} else if !confirm && !encerra {
			fmt.Println("Nao encerra a execução")
			fmt.Println("Será executado o algoritmo novamente para executar nos dados atualizados")
		}
		fmt.Println()

	}
}
