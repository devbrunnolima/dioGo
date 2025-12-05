package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Clientes struct {
	Lista []Cliente `json:"clientes"`
}

type Cliente struct {
	Id          int    `json:"id"`
	Nome        string `json:"nome"`
	NivelAcesso string `json:"nivel_acesso"`
	Senha       string `json:"senha"`
	FotoCaminho string `json:"foto"`
}

const arquivoJSON = "clientes.json"

func carregarClientes() (Clientes, error) {
	var dados Clientes

	jsonBytes, err := ioutil.ReadFile(arquivoJSON)
	if err != nil {
		return dados, err
	}

	json.Unmarshal(jsonBytes, &dados)
	return dados, nil
}

func salvarClientes(dados Clientes) error {
	jsonBytes, err := json.MarshalIndent(dados, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(arquivoJSON, jsonBytes, 0644)
}

func getClientes(w http.ResponseWriter, r *http.Request) {
	dados, err := carregarClientes()
	if err != nil {
		http.Error(w, "Não foi possível carregar o arquivo JSON", 500)
		return
	}

	json.NewEncoder(w).Encode(dados)
}

func getClientePorID(w http.ResponseWriter, r *http.Request) {
	dados, _ := carregarClientes()
	param := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(param)

	for _, cliente := range dados.Lista {
		if cliente.Id == id {
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}

	http.Error(w, "Cliente não encontrado", 404)
}

func criarCliente(w http.ResponseWriter, r *http.Request) {
	dados, _ := carregarClientes()

	var novo Cliente
	json.NewDecoder(r.Body).Decode(&novo)

	novo.Id = len(dados.Lista) + 1

	dados.Lista = append(dados.Lista, novo)
	salvarClientes(dados)

	json.NewEncoder(w).Encode(novo)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/clientes", getClientes).Methods("GET")
	r.HandleFunc("/clientes/{id}", getClientePorID).Methods("GET")
	r.HandleFunc("/clientes", criarCliente).Methods("POST")

	fmt.Println("API rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
