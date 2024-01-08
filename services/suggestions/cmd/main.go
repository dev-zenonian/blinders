package main

import (
	pkg "blinders/packages/suggestion"
	service "blinders/services/suggestion"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func initService() *service.Service {
	port := os.Getenv("PORT")
	config := &service.ServiceConfig{
		Port: port,
	}

	openaiKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(openaiKey)

	suggester, err := pkg.NewGPTSuggester(client)
	if err != nil {
		panic(err)
	}

	sv, err := service.NewTransporter(suggester, config)
	if err != nil {
		panic(err)
	}
	return sv
}

func main() {
	fmt.Println("Suggestion service")
	service := initService()
	panic(service.Listen())
}
