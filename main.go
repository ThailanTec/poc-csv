package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Dados da estrutura de exemplo
type Dados struct {
	Nome   string
	Idade  int
	Cidade string
}

func main() {
	// Suponha que você tenha uma função que consulta os dados do banco de dados e retorna uma slice de structs
	dadosDoBanco := consultarDadosDoBanco()

	// Chame a função para salvar os dados em um arquivo CSV
	err := salvarDadosCSV("empresa"+".csv", dadosDoBanco)
	if err != nil {
		fmt.Println(err)
	}
}

// Função para consultar os dados do banco de dados (substitua com sua lógica real)
func consultarDadosDoBanco() []Dados {
	// Lógica para consultar os dados do banco de dados e retornar uma slice de structs
	return []Dados{
		{"João", 30, "São Paulo"},
		{"Maria", 25, "Rio de Janeiro"},
	}
}

func salvarDadosCSV(nomeArquivo string, dados []Dados) error {
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	escritor := csv.NewWriter(arquivo)
	defer escritor.Flush()

	// Escreva o cabeçalho do CSV
	err = escritor.Write([]string{"Nome", "Idade", "Cidade"})
	if err != nil {
		return err
	}

	for _, dado := range dados {
		err := escritor.Write([]string{dado.Nome, strconv.Itoa(dado.Idade), dado.Cidade})
		if err != nil {
			return err
		}
	}

	return nil
}
