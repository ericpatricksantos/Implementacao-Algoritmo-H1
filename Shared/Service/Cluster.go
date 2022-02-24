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
		processed := Config.GetConfig().Collection[2]
		mudou := Function2.MudancaStatusTx(Txs[0], ConnectionMongoDB, DataBaseTx, CollectionTx, processed)

		if mudou {
			fmt.Println("Mudança de status concluida ", CollectionTx, " >> ", processed)
		} else {
			fmt.Println("Falha na mudança de status ", CollectionTx, " >> ", processed)
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

					mudou := Function2.MudancaStatusTx(Txs[i], ConnectionMongoDB, DataBaseTx, CollectionTx, processed)

					if mudou {
						fmt.Println("Mudança de status concluida ", CollectionTx, " >> ", processed)
					} else {
						fmt.Println("Falha na mudança de status ", CollectionTx, " >> ", processed)
						return false, true
					}
					fmt.Println("******** O hash dessa Transação será salvo em um arquivo chamado AddrInputEmpty.txt **")
					Function2.EscreverTextoSemApagar([]string{Hash}, "..\\Tcc\\AddrInputEmpty.txt")
					return false, false

				} else if QtdEnderecosVazios > 0 && QtdEnderecosVazios < lenInputs && k == 1 {
					fmt.Println("******** Na Hash: ", Txs[i].Hash)
					fmt.Println("******** Existem ", QtdEnderecosVazios, " vazios dentro da lista de inputs **")
					fmt.Println("******** Tamanho Inputs: ", lenInputs)
					fmt.Println("******** Será criado o Cluster com os endereços que estão preenchidos na lista de inputs **")
				}

				Cluster.Input, _ = Function2.RemoveDuplicados(inputs)
				confirm, existente := Function2.SaveClusterMongo(Cluster, ConnectionMongoDB, DataBaseCluster, CollectionCluster)

				if confirm {
					return true, false
				} else if !confirm && existente {
					processed := Config.GetConfig().Collection[2]
					mudou := Function2.MudancaStatusTx(Txs[i], ConnectionMongoDB, DataBaseTx, CollectionTx, processed)

					if mudou {
						fmt.Println("Mudança de status concluida ", CollectionTx, " >> ", processed)
					} else {
						fmt.Println("Falha na mudança de status ", CollectionTx, " >> ", processed)
						return false, true
					}
					return false, false
				} else {
					return false, true
				}
			}
		}
	}
	return false, true
}

func CreateClustersAddr(ConnectionMongoDB, DataBaseCluster, CollectionCluster,
	DataBaseAddr, processed, processedCluster,
	processedAddrAnalise, processedAddrAnaliseCluster string, NewAddrAnalise bool) (createClusterSucess bool, FinalizaExecucao bool) {
	var Cluster Model.Cluster
	tamanhoTxAddrOutrosNiveis := 0
	tamanhoAddrOutrosNiveis := 0
	tamanhoTxAddrAnalise := 0
	tamanhoAddrAnalise := 0
	enderecosEmAnalise := Model.UnicoEndereco{}
	enderecosOutrosNiveis := Model.UnicoEndereco{}
	if NewAddrAnalise {
		enderecosEmAnalise = Function2.GetAddrMongoDB(ConnectionMongoDB, DataBaseAddr, processedAddrAnalise)
		tamanhoTxAddrAnalise = len(enderecosEmAnalise.Txs)
		tamanhoAddrAnalise = len(enderecosEmAnalise.Address)
	}

	if tamanhoAddrAnalise > 0 && tamanhoTxAddrAnalise > 0 {
		inputs := Function2.GetAllInputs(enderecosEmAnalise)
		if tamanhoTxAddrAnalise > 0 && tamanhoAddrAnalise > 0 {
			Cluster.Hash = enderecosEmAnalise.Address
			Cluster.Input = inputs
		} else {
			fmt.Println("Não foi criado o cluster do endereço ", enderecosEmAnalise.Address)
			return false, true
		}

	} else {
		enderecosOutrosNiveis = Function2.GetAddrMongoDB(ConnectionMongoDB, DataBaseAddr, processed)
		tamanhoTxAddrOutrosNiveis = len(enderecosOutrosNiveis.Txs)
		tamanhoAddrOutrosNiveis = len(enderecosOutrosNiveis.Address)
		if tamanhoAddrOutrosNiveis > 0 && tamanhoTxAddrOutrosNiveis > 0 {
			inputs := Function2.GetAllInputs(enderecosOutrosNiveis)
			Cluster.Hash = enderecosOutrosNiveis.Address
			Cluster.Input = inputs
		} else {
			fmt.Println("Não foi criado o cluster do endereço ", enderecosOutrosNiveis.Address)
			return false, true
		}
	}
	fmt.Println()
	fmt.Println("Criando clusters com o Address: ", Cluster.Hash)
	fmt.Println()
	if len(Cluster.Hash) > 0 && len(Cluster.Input) > 0 {
		confirm, existente := Function2.SaveClusterMongo(Cluster, ConnectionMongoDB, DataBaseCluster, CollectionCluster)

		if confirm {
			if tamanhoAddrAnalise > 0 && tamanhoTxAddrAnalise > 0 {
				mudou, _ := Function2.MudancaStatusAddr(enderecosEmAnalise, ConnectionMongoDB, DataBaseAddr, processedAddrAnalise, processedAddrAnaliseCluster)
				if mudou {
					fmt.Println("Mudado o status de ", processedAddrAnalise, " >> ", processedAddrAnaliseCluster)
					return true, false
				} else {
					fmt.Println("Não foi mudado o status de ", processedAddrAnalise, " >> ", processedAddrAnaliseCluster)
					return true, true
				}
			} else if tamanhoTxAddrOutrosNiveis > 0 && tamanhoAddrOutrosNiveis > 0 {
				mudou, _ := Function2.MudancaStatusAddr(enderecosOutrosNiveis, ConnectionMongoDB, DataBaseAddr, processed, processedCluster)
				if mudou {
					fmt.Println("Mudado o status de ", processed, " >> ", processedCluster)
					return true, false
				} else {
					fmt.Println("Não foi mudado o status de ", processed, " >> ", processedCluster)
					return true, true
				}
			}
			fmt.Println("Cluster salvo, mas nao mudou o status")
			fmt.Println("Address: ", Cluster.Hash, " para ser analisado")
			return true, true
		} else if !confirm && existente {
			if tamanhoAddrAnalise > 0 && tamanhoTxAddrAnalise > 0 {
				mudou, _ := Function2.MudancaStatusAddr(enderecosEmAnalise, ConnectionMongoDB, DataBaseAddr, processedAddrAnalise, processedAddrAnaliseCluster)
				if mudou {
					fmt.Println("Mudado o status de ", processedAddrAnalise, " >> ", processedAddrAnaliseCluster)
					return false, false
				} else {
					fmt.Println("Não foi mudado o status de ", processedAddrAnalise, " >> ", processedAddrAnaliseCluster)
					return false, true
				}
			} else if tamanhoTxAddrOutrosNiveis > 0 && tamanhoAddrOutrosNiveis > 0 {
				mudou, _ := Function2.MudancaStatusAddr(enderecosOutrosNiveis, ConnectionMongoDB, DataBaseAddr, processed, processedCluster)
				if mudou {
					fmt.Println("Mudado o status de ", processed, " >> ", processedCluster)
					return true, false
				} else {
					fmt.Println("Não foi mudado o status de ", processed, " >> ", processedCluster)
					return false, true
				}
			}
			return false, true
		} else {
			return false, true
		}
	} else {
		fmt.Println("Valores do Hash e inputs do Cluster vazios")
		return false, true
	}
}
