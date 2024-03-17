package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Dados struct {
	Nome    string
	Idade   int
	Cidade  string
	Empresa string
}

func main() {
	dadosDoBanco := consultarDadosDoBanco()

	err := salvarDadosCSV("empresa"+".csv", dadosDoBanco)
	if err != nil {
		fmt.Println(err)
	}
}

func consultarDadosDoBanco() []Dados {
	return []Dados{
		{"João", 30, "São Paulo", "Itau"},
		{"Maria", 25, "Rio de Janeiro", "Bradesco"},
		{"Maria", 25, "Rio de Janeiro", "Loud"},
		{"Cleiton", 25, "Rio de Janeiro", "Furia"},
		{"John", 25, "Rio de Janeiro", "Imperial"},
		{"Clebin", 25, "Rio de Janeiro", "Brastemp"},
		{"Calton", 25, "Rio de Janeiro", "Serasa"},
		{"Mirley", 25, "Rio de Janeiro", "SBT"},
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

	err = escritor.Write([]string{"Nome", "Idade", "Cidade", "Empresa"})
	if err != nil {
		return err
	}

	for _, dado := range dados {
		err := escritor.Write([]string{dado.Nome, strconv.Itoa(dado.Idade), dado.Cidade, dado.Empresa})
		if err != nil {
			return err
		}
	}

	return nil
}
