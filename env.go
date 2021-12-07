package main

import "main/Controller"

// Variaveis Globais

var ConnectionMongoDB string = Controller.GetConfig().ConnectionMongoDB[0] //"connection string into your application code"
var DataBase string = Controller.GetConfig().DataBase[0]                   //blockchain

var UrlAPI string = Controller.GetConfig().UrlAPI[0] // "https://blockchain.info"

var LatestBlock string = Controller.GetConfig().LatestBlock
var RawTx string = Controller.GetConfig().RawTx

var CollectionLatestBlock string = Controller.GetConfig().Collection[0]
var CollectionTxs string = Controller.GetConfig().Collection[4]

var FileLogHash string = Controller.GetConfig().FileLog[0]
var FileLogBlock string = Controller.GetConfig().FileLog[1]
