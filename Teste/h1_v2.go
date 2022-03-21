package main

import (
	"fmt"
	"main/Shared/Service"
)

func main() {
	IgnoraCluster := 100000

	encerraExecucao := false
	escolhaConexao := 1
	ConnectionMongoDB := []string{
		"mongodb+srv://ericpatrick:9858epJusd@cluster0.cieqi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
	}
	DataBaseCluster := "teste"
	ColClusterProcessed := "processed"
	for {

		if encerraExecucao {
			break
		}

		fmt.Println("Aplicando Algoritmo H1")

		confirm, erro, executeAll, pausaExecucao :=
			Service.H1_V2(ConnectionMongoDB[escolhaConexao],
				DataBaseCluster, ColClusterProcessed,
				IgnoraCluster, true, 30000)

		if pausaExecucao {
			encerraExecucao = pausaExecucao
			fmt.Println("Pausa Execucao")
		} else if confirm && executeAll {
			fmt.Println("Execução finalizada com Sucesso")
			encerraExecucao = executeAll
		} else if !confirm && erro {
			encerraExecucao = erro
			fmt.Println("Execução com erro")
		} else if !confirm && !erro && !pausaExecucao {
			fmt.Println("Nao encerra a execução")
			fmt.Println("Será executado o algoritmo novamente" +
				" para executar nos dados atualizados")
		} else if confirm {
			fmt.Println("Continua execucao")
		}
		fmt.Println()

	}
}
