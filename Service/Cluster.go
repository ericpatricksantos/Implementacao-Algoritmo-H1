package Service

import (
	"fmt"
	"main/Function"
	"main/Model"
)

/*
	Implementar o algoritmo H1
*/

func H1(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) {
	clusters := Function.GetAllCluster(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
	for _, item := range clusters {
		for _, item2 := range item.Input {
			resultSearch := SearchAddr(item2, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
			if len(resultSearch) > 1 {
				result := Function.RemoveCluster(item.Hash, resultSearch)
				DeleteConfirm := DeleteListCluster(result, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

				if !DeleteConfirm {
					fmt.Println("Não foram deletados todos os clusters")
				}

				clusterResultante, _ := Function.RemoveDuplicados(UnionCluster(result))

				SaveConfirm := Function.PutListCluster(item.Hash, clusterResultante, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

				if SaveConfirm {
					fmt.Println("Cluster Resultante Atualizado")
				} else {
					fmt.Println("Cluster Resultante não foi Atualizado")
				}
			}
		}
	}
}

func UnionCluster(clusters []Model.Cluster) (result []string) {
	for _, item := range clusters {
		result = append(result, item.Input...)
	}

	result, _ = Function.RemoveDuplicados(result)

	return result
}

func SearchAddr(addr string, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) []Model.Cluster {
	return Function.SearchAddr(addr, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
}

func DeleteListCluster(clusters []Model.Cluster, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) bool {
	for _, item := range clusters {
		confirm := Function.DeleteCluster(item.Hash, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

		if !confirm {
			return false
		}
	}
	return true
}

func CreateCluster(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string,
	CollectionSalvaDados string) {
	Txs := Function.GetAllTxs(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

	for i := 0; i < len(Txs); i++ {
		var Cluster Model.Cluster
		var inputs []string
		Cluster.Hash = Txs[i].Hash
		for j := 0; j < len(Txs[i].Inputs); j++ {
			if len(Txs[i].Inputs[j].Prev_out.Addr) > 0 {
				inputs = append(inputs, Txs[i].Inputs[j].Prev_out.Addr)
			}
		}
		Cluster.Input, _ = Function.RemoveDuplicados(inputs)
		confirm := Function.SaveCluster(Cluster, ConnectionMongoDB, DataBaseMongo, CollectionSalvaDados)

		if confirm {
			fmt.Println("Salvo com Sucesso")
		} else {
			fmt.Println("Não foi Salvo")
		}
	}
}
