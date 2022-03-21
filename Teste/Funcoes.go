package main

import (
	"fmt"
	Function2 "main/Shared/Function"
)

/* Usado para testar funções*/
func mnnain() {
	//usar esse endereço para agrupar o resto
	j := Function2.GetClusterByIdenficador("37PkQ921NSAQs6b4rxAFrEg9qnSpiY1Y7d",
		"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
		"teste", "processed")

	for index, item := range j.Input {
		if item == "1JawWE56G5NmnB5iuYbFikbdETs88Fxkwo" {
			fmt.Println("Posicao: ", index)
			fmt.Println("achou")
		}
	}

	//resultSearch := Function2.SearchClustersLimit( 100000000000000 , "1JawWE56G5NmnB5iuYbFikbdETs88Fxkwo",
	//"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
	//"teste", "processed")
	//
	//fmt.Println(len(resultSearch))

}

func TesteInserir() {
	v, _, _ := Function2.AddListToList("",
		[]string{},
		"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
		"teste", "processed")
	fmt.Println(v)
}

func QuantidadeTotalEnderecos() {
	// limit 2 offset 0
	// limit 2 offset 2
	// limit 2 offset 4

	// limit 100000 offset 0
	// limit 100000 offset 100000
	// limit 100000 offset 200000
	clusters := Function2.GetAllClusterLimitOffset(2, 4, "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
		"teste", "algoritmo")
	contador := 0
	fmt.Println("QUantidade buscados: ", len(clusters))
	for _, cluster := range clusters {
		contador = contador + len(cluster.Input)
	}
	fmt.Println("Quantidade total de endereços: ", contador)
}
