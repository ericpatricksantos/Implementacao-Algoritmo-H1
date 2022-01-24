package Service

import (
	"fmt"
	"main/Shared/API"
	"main/Shared/Config"
	Function2 "main/Shared/Function"
	"main/Shared/Model"
)

/*
	Implementar o algoritmo H1
*/

func H1(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) bool {
	clusters := Function2.GetAllCluster(ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
	if len(clusters) < 1 {
		fmt.Println("Não tem nenhum cluster")
	}
	for _, item := range clusters {
		for _, item2 := range item.Input {
			resultSearch := SearchAddr(item2, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
			if len(resultSearch) > 1 {
				result := Function2.RemoveCluster(item.Hash, resultSearch)
				DeleteConfirm := DeleteListCluster(result, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

				if !DeleteConfirm {
					fmt.Println("Não foram deletados todos os clusters")
					return false
				}

				clusterResultante, _ := Function2.RemoveDuplicados(UnionCluster(result))

				SaveConfirm := Function2.PutListCluster(item.Hash, clusterResultante, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)

				if SaveConfirm {
					fmt.Println("Cluster Resultante Atualizado")
				} else {
					fmt.Println("Cluster Resultante não foi Atualizado")
					return false
				}
			}
		}
	}
	return true
}

func UnionCluster(clusters []Model.Cluster) (result []string) {
	for _, item := range clusters {
		result = append(result, item.Input...)
	}

	result, _ = Function2.RemoveDuplicados(result)

	return result
}

func SearchAddr(addr string, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) []Model.Cluster {
	return Function2.SearchClusters(addr, ConnectionMongoDB, DataBaseMongo, CollectionRecuperaDados)
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

func CreateCluster(ConnectionMongoDB, DataBaseTx, CollectionTx, DataBaseCluster,
	CollectionCluster string) (confirmCluster bool, FinalizaExecucao bool) {

	Txs := Function2.GetAllTxs(ConnectionMongoDB, DataBaseTx, CollectionTx)
	QtdEnderecosVazios := 0
	Tentativas := 2

	Hash := Txs[0].Hash

	if len(Txs) < 1 {
		fmt.Println("******** Não têm Transações para serem processadas **")
		return false, true
	} else if len(Txs[0].Inputs) < 1 {
		fmt.Println("******** A lista de Inputs está vazio **")
		fmt.Println("******** A Transação será salva em processed **")
		// processed := Config.GetConfig().Collection[2]
		// Teste
		processed := "Txprocessed"
		confirmSalveTx := Function2.SalveTxMongoDB(Txs[0], ConnectionMongoDB, DataBaseTx, processed)
		if confirmSalveTx {
			fmt.Println("******** Salva com Sucesso: ", Hash, " **")
			fmt.Println("******** A Transação será excluída de processing **")
			confirmDelete := Function2.DeleteTxMongo(Hash, ConnectionMongoDB, DataBaseTx, CollectionTx)
			if confirmDelete {
				fmt.Println("******** Transação deletada da Collection processing **")
			} else {
				fmt.Println("******** Transação nao foi deletada da Collection processing **")
				return false, true
			}
		} else {
			fmt.Println("******** Não foi salvo: ", Hash, " **")
			return false, true
		}

		return false, false
	} else {
		for k := 0; k < Tentativas; k++ {
			if QtdEnderecosVazios > 0 {
				fmt.Println("******** Tentando recuperar os endereços que estão na lista de inputs **")

				urlAPI := Config.GetConfig().UrlAPI[0]
				rawTx := Config.GetConfig().RawTx

				txTemp := API.GetTransaction(Hash, urlAPI, rawTx)
				Txs = []Model.Transaction{txTemp}

				QtdEnderecosVazios = 0
			}

			for i := 0; i < len(Txs); i++ {
				var Cluster Model.Cluster
				var inputs []string
				Cluster.Hash = Txs[i].Hash
				for j := 0; j < len(Txs[i].Inputs); j++ {
					if len(Txs[i].Inputs[j].Prev_Out.Addr) > 0 {
						inputs = append(inputs, Txs[i].Inputs[j].Prev_Out.Addr)
					} else {
						QtdEnderecosVazios++
					}
				}
				lenInputs := len(Txs[i].Inputs)
				if QtdEnderecosVazios == lenInputs && k == 0 {
					fmt.Println("******** Na Hash: ", Txs[i].Hash)
					fmt.Println("******** Os enderecos da lista de inputs estão vazios **")
					fmt.Println("******** Tamanho da lista de Inputs", lenInputs, " **")
					break
				} else if QtdEnderecosVazios > 0 && QtdEnderecosVazios < lenInputs && k == 0 {
					fmt.Println("******** Na Hash: ", Txs[i].Hash, " **")
					fmt.Println("******** Existem ", QtdEnderecosVazios, " vazios dentro da lista de inputs **")
					fmt.Println("******** Tamanho da lista de Inputs", lenInputs, " **")
					break
				} else if QtdEnderecosVazios == lenInputs && k == 1 {
					fmt.Println("******** Tentantiva de recuperar os endereços do input foi falha, pois todos permaneceram vazios **")

					fmt.Println("******** Salvando a Transação em processed **")
					processed := Config.GetConfig().Collection[2]
					confirmSalveTx := Function2.SalveTxMongoDB(Txs[i], ConnectionMongoDB, DataBaseTx, processed)
					if confirmSalveTx {
						fmt.Println("********  Transação salva com Sucesso: ", Hash, " **")

						fmt.Println("********  Excluindo a Transação processing **")
						confirmDelete := Function2.DeleteTxMongo(Hash, ConnectionMongoDB, DataBaseTx, CollectionTx)
						if confirmDelete {
							fmt.Println("******** Transação excluida com sucesso **")
						} else {
							fmt.Println("******** Transação nao foi excluida **")
							return false, true
						}
					} else {
						fmt.Println("******** Não foi salvo: ", Txs[i].Hash, " **")
						return false, true
					}

					fmt.Println("******** O hash dessa Transação será salvo em um arquivo chamado AddrInputEmpty.txt **")
					Function2.EscreverTextoSemApagar([]string{Hash}, "..\\Tcc\\AddrInputEmpty.txt")
					break
				} else if QtdEnderecosVazios > 0 && QtdEnderecosVazios < lenInputs && k == 1 {
					fmt.Println("******** Na Hash: ", Txs[i].Hash)
					fmt.Println("******** Existem ", QtdEnderecosVazios, " vazios dentro da lista de inputs **")
					fmt.Println("******** Tamanho Inputs: ", lenInputs)
					fmt.Println("******** Será criado o Cluster com os endereços que estão preenchidos na lista de inputs **")
				}

				Cluster.Input, _ = Function2.RemoveDuplicados(inputs)
				confirm := Function2.SaveCluster(Cluster, ConnectionMongoDB, DataBaseCluster, CollectionCluster)

				if confirm {
					return true, false
				} else {
					return false, true
				}
			}
		}
	}
	return false, true
}
