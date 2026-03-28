package main

import (
	"rest-api-2/internal/handlers"
	"rest-api-2/internal/repositories"
	"rest-api-2/internal/usecases"
)

func main() {
	repos := repositories.New()
	usecases := usecases.New(repos)
	h := handlers.New(usecases)
	h.Listen(3000)

}
