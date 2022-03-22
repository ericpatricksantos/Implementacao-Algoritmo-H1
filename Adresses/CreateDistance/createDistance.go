package main

import (
	"fmt"
	"main/Shared/Function"
)

func main() {
	ConnectionMongoDB := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"

	DataBaseAddr := "Adresses"
	// Collection do nivel dos endereços do Farao
	awaitingProcessingEnderecosEmAnalise := "awaitingProcessingEnderecosEmAnalise"
	processedEnderecosEmAnalise := "processedEnderecosEmAnalise"
	// Collection de outros niveis
	awaitingProcessingAddr := "awaitingProcessing"
	processingAddr := "processing"
	processedAddr := "processed"

	DataBaseDistancia := "Distancia"
	awaitingProcessing := "awaitingProcessingTerceiroNivel"
	processsedDistancia := "processed"

	tempo := 2

	confirm := Function.CreateDistance(ConnectionMongoDB,
		DataBaseAddr, awaitingProcessingEnderecosEmAnalise, processedEnderecosEmAnalise,
		awaitingProcessingAddr, processingAddr, processedAddr,
		DataBaseDistancia, awaitingProcessing, processsedDistancia, tempo)
	if confirm {
		fmt.Println()
		fmt.Println("Todos os endereços foram salvos")
	} else {
		fmt.Println()
		fmt.Println("Não foram salvos todos os endereços")
	}
}
