package Function

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Funcao que le o conteudo do arquivo e retorna um slice the string com todas as linhas do arquivo
func LerTexto(caminhoDoArquivo string) ([]string, error) {
	// Abre o arquivo
	arquivo, err := os.Open(caminhoDoArquivo)
	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return nil, err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	// Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
	return linhas, scanner.Err()
}

// Funcao que escreve um texto no arquivo e retorna um erro caso tenha algum problema
func EscreverTexto(linhas []string, caminhoDoArquivo string) error {
	// Cria o arquivo de texto
	arquivo, err := os.Create(caminhoDoArquivo)
	// Caso tenha encontrado algum erro retornar ele
	if err != nil {
		return err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	escritor := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		fmt.Fprintln(escritor, linha)
	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return escritor.Flush()
}

func EscreverTextoSemApagar(linhas []string, caminhoDoArquivo string) error {
	valoresAntigos, er := LerTexto(caminhoDoArquivo)

	if er != nil {
		return er
	}
	valoresAtual := append(valoresAntigos, linhas...)
	// Cria o arquivo de texto
	arquivo, err := os.Create(caminhoDoArquivo)
	// Caso tenha encontrado algum erro retornar ele
	if err != nil {
		return err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	escritor := bufio.NewWriter(arquivo)
	for _, linha := range valoresAtual {
		fmt.Fprintln(escritor, linha)
	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return escritor.Flush()
}

func GetIndiceLogIndice(nomeArquivoIndice string) int {
	valorLogIndice, err := LerTexto(nomeArquivoIndice)
	indiceInicial := 0
	if len(valorLogIndice) > 0 {
		if err != nil {
			fmt.Print(err.Error())
			fmt.Println("Erro na função GetIndiceLogIndice na leitura do arquivo")
		}
		var er error
		indiceInicial, er = strconv.Atoi(valorLogIndice[0])
		if er != nil {
			fmt.Print(er.Error())
			fmt.Println("Erro na função GetIndiceLogIndice na conversão de string para int")
		}

	}

	return indiceInicial
}
