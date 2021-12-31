package Function

import (
	"log"
	"main/Shared/Database"
	"main/Shared/Model"

	"gopkg.in/mgo.v2/bson"
)

/*
	Esse arquivo foi criado para armazenar todas as funções que são utilizadas frequentemente.

*/

func GetAllLatestBlock(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) (blocks []Model.LatestBlock) {

	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		panic(err)
	}

	// Free the resource when mainn dunction is  returned
	defer Database.Close(client, ctx, cancel)

	// create a filter an option of type interface,
	// that stores bjson objects.
	var filter, option interface{}

	// filter  gets all document,
	// with maths field greater that 70
	filter = bson.M{}

	//  option remove id field from all documents
	option = bson.M{}

	// call the query method with client, context,
	// database name, collection  name, filter and option
	// This method returns momngo.cursor and error if any.
	cursor, err := Database.Query(client, ctx, DataBaseMongo,
		CollectionRecuperaDados, filter, option)
	// handle the errors.
	if err != nil {
		panic(err)
	}

	// le os documentos em partes, testei com 1000 documentos e deu certo
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var bloco Model.LatestBlock

		if err := cursor.Decode(&bloco); err != nil {
			log.Fatal(err)
		}

		blocks = append(blocks, bloco)

	}

	return blocks
}

func SaveLatestBlock(latestBlock Model.LatestBlock,
	ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	if len(latestBlock.TxIndexes) > 0 {
		cliente, contexto, cancel, errou := Database.Connect(ConnectionMongoDB)
		if errou != nil {
			log.Fatal(errou)
		}

		Database.Ping(cliente, contexto)
		defer Database.Close(cliente, contexto, cancel)

		Database.ToDoc(latestBlock)

		insert, err := Database.InsertOne(cliente, contexto, DataBaseMongo, Collection, latestBlock)

		// handle the error
		if err != nil {
			panic(err)
		}

		if insert.InsertedID != nil {
			return true
		} else {
			return false
		}

	} else {
		return false
	}
}

func DeleteLatestBlock(hash string, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) bool {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		panic(err)
	}

	// Free the resource when mainn dunction is  returned
	defer Database.Close(client, ctx, cancel)

	// create a filter an option of type interface,
	// that stores bjson objects.
	var filter interface{}

	// filter  gets all document,
	// with maths field greater that 70
	filter = bson.M{
		"hash": hash,
	}

	cursor, err := Database.DeleteOne(client, ctx, DataBaseMongo, CollectionRecuperaDados, filter)

	if err != nil {
		panic(err)
	}
	// verifica a quantidade de linhas afetadas
	if cursor.DeletedCount > 0 {
		return true
	} else {
		return false
	}
}
