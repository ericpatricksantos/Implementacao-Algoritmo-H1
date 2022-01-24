package Function

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"log"
	"main/Shared/Database"
	"main/Shared/Model"
)

func GetAllCluster(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) (Clusters []Model.Cluster) {

	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {

		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função GetAllCluster - {Function/Cluster.go}")
		fmt.Println()

		fmt.Println()
		fmt.Println(err.Error())
		fmt.Println()

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

		fmt.Println()
		fmt.Println("Erro na resposta da função Query - {Database/Mongo.go} que esta sendo chamada na Função GetAllCluster - {Function/Cluster.go}")
		fmt.Println()

		fmt.Println()
		fmt.Println(err.Error())
		fmt.Println()

		panic(err)
	}

	// le os documentos em partes, testei com 1000 documentos e deu certo
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var cluster Model.Cluster

		if err := cursor.Decode(&cluster); err != nil {

			fmt.Println()
			fmt.Println("Erro na resposta da função Decode que esta sendo chamada na Função GetAllCluster - {Function/Cluster.go}")
			fmt.Println()

			fmt.Println()
			fmt.Println(err.Error())
			fmt.Println()

			log.Fatal(err)
		}

		Clusters = append(Clusters, cluster)

	}

	return Clusters
}

func SearchClusters(addr string, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) (result []Model.Cluster) {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {

		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função SearchClusters - {Function/Cluster.go}")
		fmt.Println()

		fmt.Println()
		fmt.Println(err.Error())
		fmt.Println()

		panic(err)
	}

	// Free the resource when mainn dunction is  returned
	defer Database.Close(client, ctx, cancel)

	// create a filter an option of type interface,
	// that stores bjson objects.
	var filter, option interface{}

	// filter  gets all document,
	// with maths field greater that 70
	filter = bson.M{
		"input": addr,
	}

	option = bson.M{}

	cursor, err := Database.Query(client, ctx, DataBaseMongo, CollectionRecuperaDados, filter, option)

	// handle the errors.
	if err != nil {

		fmt.Println()
		fmt.Println("Erro na resposta da função Query - {Database/Mongo.go} que esta sendo chamada na Função SearchClusters - {Function/Cluster.go}")
		fmt.Println()

		fmt.Println()
		fmt.Println(err.Error())
		fmt.Println()

		panic(err)
	}

	// le os documentos em partes, testei com 1000 documentos e deu certo
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var cluster Model.Cluster

		if err := cursor.Decode(&cluster); err != nil {

			fmt.Println()
			fmt.Println("Erro na resposta da função Decode que esta sendo chamada na Função SearchClusters - {Function/Cluster.go}")
			fmt.Println()

			fmt.Println()
			fmt.Println(err.Error())
			fmt.Println()

			log.Fatal(err)
		}

		result = append(result, cluster)

	}

	return result
}

func DeleteCluster(hash string, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) bool {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {

		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função DeleteCluster - {Function/Cluster.go}")
		fmt.Println()

		fmt.Println()
		fmt.Println(err.Error())
		fmt.Println()

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

		fmt.Println()
		fmt.Println("Erro na resposta da função DeleteOne - {Database/Mongo.go} que esta sendo chamada na Função DeleteCluster - {Function/Cluster.go}")
		fmt.Println()

		fmt.Println()
		fmt.Println(err.Error())
		fmt.Println()

		panic(err)
	}
	// verifica a quantidade de linhas afetadas
	if cursor.DeletedCount > 0 {

		return true
	} else {

		return false
	}
}

func SaveCluster(Cluster Model.Cluster, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	if len(Cluster.Hash) > 0 {
		cliente, contexto, cancel, errou := Database.Connect(ConnectionMongoDB)
		if errou != nil {

			fmt.Println()
			fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função SaveCluster - {Function/Cluster.go}")
			fmt.Println()

			fmt.Println()
			fmt.Println(errou.Error())
			fmt.Println()

			log.Fatal(errou)
		}

		Database.Ping(cliente, contexto)
		defer Database.Close(cliente, contexto, cancel)

		Database.ToDoc(Cluster)

		result, err := Database.InsertOne(cliente, contexto, DataBaseMongo, Collection, Cluster)

		// handle the error
		if err != nil {

			fmt.Println()
			fmt.Println("Erro na resposta da função InsertOne - {Database/Mongo.go} que esta sendo chamada na Função SaveCluster - {Function/Cluster.go}")
			fmt.Println()

			fmt.Println()
			fmt.Println(errou.Error())
			fmt.Println()

			panic(err)
		}

		if result.InsertedID != nil {
			return true
		} else {
			return false
		}

	} else {
		fmt.Println("Cluster invalido: Hash da Transação esta vazio")
		return false
	}

	return false
}

func PutListCluster(Hash string, Input []string, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	for _, item := range Input {
		PutCluster(Hash, item, ConnectionMongoDB, DataBaseMongo, Collection)
	}
	return true
}

func PutCluster(Hash string, Input string, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	if len(Input) > 0 {
		cliente, contexto, cancel, errou := Database.Connect(ConnectionMongoDB)
		if errou != nil {

			fmt.Println()
			fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função PutCluster - {Function/Cluster.go}")
			fmt.Println()

			fmt.Println()
			fmt.Println(errou.Error())
			fmt.Println()

			log.Fatal(errou)
		}

		Database.Ping(cliente, contexto)
		defer Database.Close(cliente, contexto, cancel)

		var result *mongo.UpdateResult
		var err error

		filter := bson.M{
			"hash": Hash,
		}
		update := bson.M{"$addToSet": bson.M{"input": Input}}

		result, err = Database.UpdateOne(cliente, contexto, DataBaseMongo, Collection, filter, update)

		// handle the error
		if err != nil {

			fmt.Println()
			fmt.Println("Erro na resposta da função UpdateOne - {Database/Mongo.go} que esta sendo chamada na Função PutCluster - {Function/Cluster.go}")
			fmt.Println()

			fmt.Println()
			fmt.Println(errou.Error())
			fmt.Println()
			panic(err)
		}

		if result.ModifiedCount > 0 {
			return true
		} else {
			return false
		}
	} else {
		fmt.Println("O Hash do Cluster esta vazio")
		return false
	}
	return false
}
