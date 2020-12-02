package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	ID       string   `json:"id,omitempty"`
	Pname    string   `json:"pname,omitempty"`
	Pcompany string   `json:"pcompany,omitempty"`
	Address  *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var product []Product

func GetProductIdEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range product {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}
func GetProductEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(emp)
}
func CreateProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var product Product
	_ = json.NewDecoder(req.Body).Decode(&person)
	product.ID = params["id"]
	product = append(product, product)
	json.NewEncoder(w).Encode(product)
}
func DeleteProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range product {
		if item.ID == params["id"] {
			product = append(product[:index], product[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(product)
}
func main() {
	router := mux.NewRouter()
	product = append(product, Product{ID: "1", Pname: "Laptop", Pcompany: "Dell",
		Address: &Address{City: "Dublin", State: "CA"}})
	Product = append(product, Product{ID: "2", Pname: "mobile", Pcompany: "OOPO"})
	router.HandleFunc("/product", GetProductEndpoint).Methods("GET")
	router.HandleFunc("/product/{id}", GetproductIdEndpoint).Methods("GET")
	router.HandleFunc("/product/{id}", CreateProductEndpoint).Methods("POST")
	router.HandleFunc("/product/{id}", DeleteProductEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
