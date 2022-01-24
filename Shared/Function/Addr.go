package Function

import (
	"fmt"
	"log"
	"main/Shared/API"
	"main/Shared/Database"
	"main/Shared/Model"

	"gopkg.in/mgo.v2/bson"
)

// Funções para Addr

func SaveAddr(Addr Model.UnicoEndereco, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	if len(Addr.Address) > 0 {
		cliente, contexto, cancel, errou := Database.Connect(ConnectionMongoDB)
		if errou != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go}  que esta sendo chamada na Função SaveAddr - {Function/Addr.go}")
			fmt.Println()

			panic(errou)
		}

		Database.Ping(cliente, contexto)
		defer Database.Close(cliente, contexto, cancel)

		Database.ToDoc(Addr)

		result, err := Database.InsertOne(cliente, contexto, DataBaseMongo, Collection, Addr)
		// handle the error
		if err != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função InsertOne - {Database/Mongo.go}  que esta sendo chamada na Função SaveAddr - {Function/Addr.go}")
			fmt.Println()

			panic(err)
		}

		if result.InsertedID != nil {
			return true
		} else {
			return false
		}

	} else {
		fmt.Println("Addr esta vazio -", Addr.Address, "-")
		return false
	}
}

func SaveAddrSimplificado(Addr Model.Endereco, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	if len(Addr.Address) > 0 {
		cliente, contexto, cancel, errou := Database.Connect(ConnectionMongoDB)
		if errou != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go}  que esta sendo chamada na Função SaveAddrSimplificado - {Function/Addr.go}")
			fmt.Println()

			panic(errou)
		}

		Database.Ping(cliente, contexto)
		defer Database.Close(cliente, contexto, cancel)

		Database.ToDoc(Addr)

		result, err := Database.InsertOne(cliente, contexto, DataBaseMongo, Collection, Addr)
		// handle the error
		if err != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função InsertOne - {Database/Mongo.go}  que esta sendo chamada na Função SaveAddrSimplificado - {Function/Addr.go}")
			fmt.Println()

			panic(err)
		}

		if result.InsertedID != nil {
			return true
		} else {
			return false
		}

	} else {
		fmt.Println("Addr esta vazio -", Addr.Address, "-")
		return false
	}
}

func SalveAddrSimplicadoMongoDB(addr Model.Endereco, ConnectionMongoDB, DataBaseMongo, Collection string) bool {
	confirm := CheckAddr(ConnectionMongoDB, DataBaseMongo, Collection, "address", addr.Address)
	if confirm {
		fmt.Println("Esse addr ja existe nessa Collection: ", Collection)
		return false
	}
	return SaveAddrSimplificado(addr, ConnectionMongoDB, DataBaseMongo, Collection)
}

func SalveAddrMongoDB(addr Model.UnicoEndereco, ConnectionMongoDB, DataBaseMongo, Collection string) bool {
	confirm := CheckAddr(ConnectionMongoDB, DataBaseMongo, Collection, "address", addr.Address)
	if confirm {
		fmt.Println("Esse addr ja existe nessa Collection: ", Collection)
		return false
	}
	return SaveAddr(addr, ConnectionMongoDB, DataBaseMongo, Collection)
}

func GetAddr(endereco, urlAPI, RawAddr string) Model.UnicoEndereco {
	return API.GetUnicoEndereco(endereco, urlAPI, RawAddr)
}

func GetAddrMongoDB(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) (addr Model.UnicoEndereco) {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go}  que esta sendo chamada na Função GetAddrMongoDB - {Function/Addr.go}")
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
		fmt.Println("Erro na resposta da função Query - {Database/Mongo.go}  que esta sendo chamada na Função GetAddrMongoDB - {Function/Addr.go}")
		fmt.Println()

		panic(err)
	}

	// le os documentos em partes, testei com 1000 documentos e deu certo
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		if err := cursor.Decode(&addr); err != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função Decode que esta sendo chamada na Função GetAddrMongoDB - {Function/Addr.go}")
			fmt.Println()

			log.Fatal(err)
		}

		return addr
	}

	return Model.UnicoEndereco{}
}

func CheckAddr(ConnectionMongoDB, dataBase, col, key, code string) bool {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função CheckAddr - {Function/Addr.go}")
		fmt.Println()
		panic(err)
	}

	// Free the resource when mainn dunction is  returned
	defer Database.Close(client, ctx, cancel)
	count, err := Database.CountElemento(client, ctx, dataBase, col, key, code)
	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função CountElemento - {Database/Mongo.go} que esta sendo chamada na Função CheckAddr - {Function/Addr.go}")
		fmt.Println()
		panic(err)
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}

func DeleteAddrMongo(addr string, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	confirm := CheckAddr(ConnectionMongoDB, DataBaseMongo, Collection, "address", addr)
	if !confirm {
		fmt.Println("Esse Addr não existe nessa Collection, por isso não tem como excluir: ", Collection)
		return false
	}
	return DeleteAddr(addr, ConnectionMongoDB, DataBaseMongo, Collection)
}

func DeleteAddr(addr string, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função DeleteAddr - {Function/Addr.go}")
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
		"address": addr,
	}

	cursor, err := Database.DeleteOne(client, ctx, DataBaseMongo, Collection, filter)

	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função DeleteOne - {Database/Mongo.go} que esta sendo chamada na Função DeleteAddr - {Function/Addr.go}")
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

func GetAllAddr(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) (addrs []Model.UnicoEndereco) {

	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go}  que esta sendo chamada na Função GetAllAddr - {Function/Addr.go}")
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
		fmt.Println("Erro na resposta da função Query - {Database/Mongo.go}  que esta sendo chamada na Função GetAllAddr - {Function/Addr.go}")
		fmt.Println()

		panic(err)
	}

	// le os documentos em partes, testei com 1000 documentos e deu certo
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var addr Model.UnicoEndereco

		if err := cursor.Decode(&addr); err != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função Decode que esta sendo chamada na Função GetAllAddr - {Function/Addr.go}")
			fmt.Println()

			log.Fatal(err)
		}

		addrs = append(addrs, addr)

	}

	return addrs
}

func SearchAddr(addr string, ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) (result []Model.UnicoEndereco) {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {

		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função SearchAddr - {Function/Cluster.go}")
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
		fmt.Println("Erro na resposta da função Query - {Database/Mongo.go} que esta sendo chamada na Função SearchAddr - {Function/Cluster.go}")
		fmt.Println()

		fmt.Println()
		fmt.Println(err.Error())
		fmt.Println()

		panic(err)
	}

	// le os documentos em partes, testei com 1000 documentos e deu certo
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var addr Model.UnicoEndereco

		if err := cursor.Decode(&addr); err != nil {

			fmt.Println()
			fmt.Println("Erro na resposta da função Decode que esta sendo chamada na Função SearchAddr - {Function/Cluster.go}")
			fmt.Println()

			fmt.Println()
			fmt.Println(err.Error())
			fmt.Println()

			log.Fatal(err)
		}

		result = append(result, addr)

	}

	return result
}

// AP - AwaitingProcessing
// P  - Processing
func MudancaStatusAddr_AP_P(addr Model.UnicoEndereco, ConnectionMongoDB, DataBase, awaitingProcessing, processing string) bool {

	salvo := SalveAddrMongoDB(addr, ConnectionMongoDB, DataBase, processing)

	if !salvo {
		fmt.Println("Não foi salvo com Sucesso")
		return false
	} else {
		fmt.Println("Salvo com sucesso a Addr na collection ", processing)
	}

	deletado := DeleteAddrMongo(addr.Address, ConnectionMongoDB, DataBase, awaitingProcessing)

	if !deletado {
		fmt.Println("Address: ", addr.Address, " não foi deletado de ", awaitingProcessing)
		return false
	} else {
		fmt.Println("Deletado com sucesso a Addr da collection ", awaitingProcessing)
	}

	return true
}

func MudancaStatusAddr_Processing_Processed(addr Model.UnicoEndereco, ConnectionMongoDB, DataBase, processing, processed string) bool {

	salvo := SalveAddrMongoDB(addr, ConnectionMongoDB, DataBase, processed)

	if !salvo {
		fmt.Println("Não foi salvo com Sucesso")
		return false
	} else {
		fmt.Println("Salvo com sucesso a Addr na collection ", processed)
	}

	deletado := DeleteAddrMongo(addr.Address, ConnectionMongoDB, DataBase, processing)

	if !deletado {
		fmt.Println("Address: ", addr.Address, " não foi deletado de ", processing)
		return false
	} else {
		fmt.Println("Deletado com sucesso a Addr da collection ", processing)
	}

	return true
}

// Funções para MultiAddr

func SaveMultiAddr(MultiAddr Model.MultiEndereco, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	if len(MultiAddr.Addresses) > 0 {
		cliente, contexto, cancel, errou := Database.Connect(ConnectionMongoDB)
		if errou != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go}  que esta sendo chamada na Função SaveMultiAddr - {Function/Addr.go}")
			fmt.Println()

			panic(errou)
		}

		Database.Ping(cliente, contexto)
		defer Database.Close(cliente, contexto, cancel)

		Database.ToDoc(MultiAddr)

		result, err := Database.InsertOne(cliente, contexto, DataBaseMongo, Collection, MultiAddr)

		// handle the error
		if err != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função InsertOne - {Database/Mongo.go}  que esta sendo chamada na Função SaveMultiAddr - {Function/Addr.go}")
			fmt.Println()

			panic(err)
		}

		if result.InsertedID != nil {
			return true
		} else {
			return false
		}

	} else {
		fmt.Println("Array de MultiAddr esta vazio -", len(MultiAddr.Addresses), "-")
		return false
	}
}

// resolver esse metodo
func SalveMultiAddrMongoDB(MultiAddr Model.MultiEndereco, ConnectionMongoDB, DataBaseMongo, Collection string) bool {
	//confirm := CheckMultiAddr(ConnectionMongoDB, DataBaseMongo, Collection, "address", MultiAddr.Addresses)
	//if confirm {
	//	fmt.Println("Esse MultiAddr ja existe nessa Collection: ", Collection)
	//	return false
	//}
	return SaveMultiAddr(MultiAddr, ConnectionMongoDB, DataBaseMongo, Collection)
}

// Busca na API
func GetMultiAddr(MultiAddr []string, urlAPI, endpoint string) Model.MultiEndereco {
	return API.GetMultiEnderecos(MultiAddr, urlAPI, endpoint)
}

func GetAddrMultiMongoDB(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) (MultiAddr Model.MultiEndereco) {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go}  que esta sendo chamada na Função GetAddrMultiMongoDB - {Function/Addr.go}")
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
		fmt.Println("Erro na resposta da função Query - {Database/Mongo.go}  que esta sendo chamada na Função GetAddrMultiMongoDB - {Function/Addr.go}")
		fmt.Println()

		panic(err)
	}

	// le os documentos em partes, testei com 1000 documentos e deu certo
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		if err := cursor.Decode(&MultiAddr); err != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função Decode que esta sendo chamada na Função GetAddrMultiMongoDB - {Function/Addr.go}")
			fmt.Println()

			log.Fatal(err)
		}

		return MultiAddr
	}

	return Model.MultiEndereco{}
}

// resolver esse metodo
func CheckMultiAddr(ConnectionMongoDB, dataBase, col, key string, code []string) bool {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função CheckMultiAddr - {Function/Addr.go}")
		fmt.Println()
		panic(err)
	}

	// Free the resource when mainn dunction is  returned
	defer Database.Close(client, ctx, cancel)
	for _, item := range code {
		count, err := Database.CountElemento(client, ctx, dataBase, col, key, item)
		if err != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função CountElemento - {Database/Mongo.go} que esta sendo chamada na Função CheckMultiAddr - {Function/Addr.go}")
			fmt.Println()
			panic(err)
		}
		if count > 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

// resolver esse metodo
func DeleteMultiAddrMongo(MultiAddr string, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	//confirm := CheckMultiAddr(ConnectionMongoDB, DataBaseMongo, Collection, "address", MultiAddr)
	//if !confirm {
	//	fmt.Println("Esse Addr não existe nessa Collection, por isso não tem como excluir: ", Collection)
	//	return false
	//}
	return DeleteMultiAddr(MultiAddr, ConnectionMongoDB, DataBaseMongo, Collection)
}

// resolver esse metodo
func DeleteMultiAddr(MultiAddr string, ConnectionMongoDB string, DataBaseMongo string, Collection string) bool {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go} que esta sendo chamada na Função DeleteMultiAddr - {Function/Addr.go}")
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
		"address": MultiAddr,
	}

	cursor, err := Database.DeleteOne(client, ctx, DataBaseMongo, Collection, filter)

	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função DeleteOne - {Database/Mongo.go} que esta sendo chamada na Função DeleteMultiAddr - {Function/Addr.go}")
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

func GetAllMultiAddr(ConnectionMongoDB string, DataBaseMongo string, CollectionRecuperaDados string) (multiAddrs []Model.MultiEndereco) {

	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Database.Connect(ConnectionMongoDB)
	if err != nil {
		fmt.Println()
		fmt.Println("Erro na resposta da função Connect - {Database/Mongo.go}  que esta sendo chamada na Função GetAllAddr - {Function/Addr.go}")
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
		fmt.Println("Erro na resposta da função Query - {Database/Mongo.go}  que esta sendo chamada na Função GetAllAddr - {Function/Addr.go}")
		fmt.Println()

		panic(err)
	}

	// le os documentos em partes, testei com 1000 documentos e deu certo
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var multiAddr Model.MultiEndereco

		if err := cursor.Decode(&multiAddr); err != nil {
			fmt.Println()
			fmt.Println("Erro na resposta da função Decode que esta sendo chamada na Função GetAllAddr - {Function/Addr.go}")
			fmt.Println()

			log.Fatal(err)
		}

		multiAddrs = append(multiAddrs, multiAddr)

	}

	return multiAddrs
}
