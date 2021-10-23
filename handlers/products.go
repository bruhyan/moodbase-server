// This is for testing 0Auth
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
    Id int
    Name string
    Slug string
    Description string
}

type ProductHandler struct {
	logger *log.Logger
}

func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

func (ph *ProductHandler) ServeHTTP (rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ph.getProducts(rw, r)
		return
	}
}

func (ph *ProductHandler) getProducts(rw http.ResponseWriter, r *http.Request) {
	products = GetProducts()
	e := json.NewEncoder(rw)
	err := e.Encode(products)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}


// should be in data file from here on

type Products []*Product

var products = []*Product{
	{Id: 1, Name: "World of Authcraft", Slug: "world-of-authcraft", Description : "Battle bugs and protect yourself from invaders while you explore a scary world with no security"},
	{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description : "Explore the depths of the sea in this one of a kind underwater experience"},
	{Id: 3, Name: "Dinosaur Park", Slug : "dinosaur-park", Description : "Go back 65 million years in the past and ride a T-Rex"},
	{Id: 4, Name: "Cars VR", Slug : "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description : "Pick up the bow and arrow and master the art of archery"},
	{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description : "Explore the seven wonders of the world in VR"},
}

func GetProducts() Products {
	return products
}