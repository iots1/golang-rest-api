package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"peepo/rest-api/models"
)

func main() {
	var goVersion float32 = 1.18
	fmt.Println("Hello Go ", goVersion)

	var name string
	name = "PEEPO"
	fmt.Printf("My name is %s \n", name)

	address := "138/12 เชียงใหม่"
	fmt.Printf("ที่อยู่ %s \n", address)

	postalCode := 50140
	fmt.Printf("รหัสไปรษณีย์ %d \n", postalCode)

	myProduct := models.Product{
		ProductId:   1,
		ProductName: "milk",
		ProductQty:  5,
	}
	fmt.Printf("%v \n", myProduct)
	fmt.Printf("%v \n", myProduct.ProductName)

	if myProduct.ProductName == "milk" {
		fmt.Println("ดื่มนม")
	} else {
		fmt.Println("ไม่ใช่นม")
	}

	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) { // (1)
	fmt.Fprint(w, "Welcome to the HomePage!") // (2)
}

func userPage(w http.ResponseWriter, r *http.Request) {
	responseData := models.User{
		Username: "PEEPO",
		Password: "1234",
	}
	bytes, _ := json.Marshal(responseData)
	fmt.Fprint(w, string(bytes))
}

func productPage(w http.ResponseWriter, r *http.Request) {
	responseData := models.Product{
		ProductId:   1,
		ProductName: "นม",
		ProductQty:  5,
	}
	json.NewEncoder(w).Encode(responseData)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/user", userPage)
	http.HandleFunc("/product", productPage)
	http.ListenAndServe(":8080", nil)
}
