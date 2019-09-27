package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(checkLoginUnico)
}

func checkLoginUnico(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	var loginUnicoURL string = os.Getenv("LoginUnicoURL")

	client := &http.Client{}

	response, err := client.Get(loginUnicoURL)
	if err != nil {
		log.Println("Erro client.Get: " + err.Error())
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
		}

	request, err := http.NewRequest("GET", loginUnicoURL, nil)
	if err != nil {
		log.Println("Erro http.NewRequest: " + err.Error())
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
		}
	log.Println("headers: " + req.Headers["Authorization"])
	request.Header.Add("Authorization", req.Headers["Authorization"])

	response, err = client.Do(request)
	if err != nil {
		log.Println("Erro client.Do " + err.Error())
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
		}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		return events.APIGatewayProxyResponse{Body: "ok", StatusCode: http.StatusOK}, nil
	} 
	log.Println("loginUnico falhou com a Authorization: " + req.Headers["Authorization"])
	return events.APIGatewayProxyResponse{Body: "unauthorized", StatusCode: http.StatusUnauthorized}, nil
}
