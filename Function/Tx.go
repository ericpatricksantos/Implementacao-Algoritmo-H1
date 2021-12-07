package Function

import (
	"log"
	"main/Database"
	"main/Model"
)

func SaveTx(Tx Model.Transaction, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	if len(Tx.Inputs) > 0 {
		cliente, contexto, cancel, errou := Database.Connect(ConnectionMongoDB)
		if errou != nil {
			log.Fatal(errou)
		}

		Database.Ping(cliente, contexto)
		defer Database.Close(cliente, contexto, cancel)

		Database.ToDoc(Tx)

		_, err := Database.InsertOne(cliente, contexto, DataBaseMongo, Collection, Tx)

		// handle the error
		if err != nil {
			panic(err)
		} else {
			return true
		}

	} else {
		return false
	}
}
