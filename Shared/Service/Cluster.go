package Service

import (
	"fmt"
	Function2 "main/Shared/Function"
	"main/Shared/Model"
)

/*
	Implementar o algoritmo H1
*/

func H1(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) {
	clusters := Function2.GetAllCluster(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
	for _, item := range clusters {
		for _, item2 := range item.Input {
			resultSearch := SearchAddr(item2, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
			if len(resultSearch) > 1 {
				result := Function2.RemoveCluster(item.Hash, resultSearch)
				DeleteConfirm := DeleteListCluster(result, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

				if !DeleteConfirm {
					fmt.Println("Não foram deletados todos os clusters")
				}

				clusterResultante, _ := Function2.RemoveDuplicados(UnionCluster(result))

				SaveConfirm := Function2.PutListCluster(item.Hash, clusterResultante, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

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

	result, _ = Function2.RemoveDuplicados(result)

	return result
}

func SearchAddr(addr string, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) []Model.Cluster {
	return Function2.SearchAddr(addr, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
}

func DeleteListCluster(clusters []Model.Cluster, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) bool {
	for _, item := range clusters {
		confirm := Function2.DeleteCluster(item.Hash, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

		if !confirm {
			return false
		}
	}
	return true
}

func CreateCluster(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string,
	CollectionSalvaDados string) {
	Txs := Function2.GetAllTxs(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

	for i := 0; i < len(Txs); i++ {
		var Cluster Model.Cluster
		var inputs []string
		Cluster.Hash = Txs[i].Hash
		for j := 0; j < len(Txs[i].Inputs); j++ {
			if len(Txs[i].Inputs[j].Prev_out.Addr) > 0 {
				inputs = append(inputs, Txs[i].Inputs[j].Prev_out.Addr)
			}
		}
		Cluster.Input, _ = Function2.RemoveDuplicados(inputs)
		confirm := Function2.SaveCluster(Cluster, ConnectionMongoDB, DataBaseMongo, CollectionSalvaDados)

		if confirm {
			fmt.Println("Salvo com Sucesso")
		} else {
			fmt.Println("Não foi Salvo")
		}
	}
}
