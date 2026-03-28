package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api-2/internal/models"
)

func main() {
	req := models.CreateUserRequest{
		Name:  "Alex Arlindo Nhabinde",
		Email: "alex@piripiri.chat",
	}

	b, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:3000/users", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	var responseApi models.CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusCreated {
		panic("error ao criar o user")
	}

	fmt.Println("Novo user criado com sucesso.", responseApi)

}
