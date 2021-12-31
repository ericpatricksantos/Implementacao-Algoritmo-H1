package main

import "main/Shared/Controller"

var ConnectionMongoDB string = Controller.GetConfig().ConnectionMongoDB[0] //"connection string into your application code"
var DataBase string = Controller.GetConfig().DataBase[0]                   //blockchain

var UrlAPI string = Controller.GetConfig().UrlAPI[0] // "https://blockchain.info"

var LatestBlock string = Controller.GetConfig().LatestBlock
var RawTx string = Controller.GetConfig().RawTx

var CollectionLatestBlock string = Controller.GetConfig().Collection[0]
var CollectionTesteTxs string = Controller.GetConfig().Collection[1]
var CollectionTesteCluster string = Controller.GetConfig().Collection[2]
var CollectionCluster string = Controller.GetConfig().Collection[3]
var CollectionTxs string = Controller.GetConfig().Collection[4]

var FileLogHash string = Controller.GetConfig().FileLog[0]
var FileLogBlock string = Controller.GetConfig().FileLog[1]

func main() {

	////Reorganiza os Cluster utilizando o algoritmo H1
	Controller.H1(ConnectionMongoDB, DataBase, CollectionTesteCluster)
}
