package main

import (
	"log"
	"net/http"
	"os"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(checkLoginUnico)
	//http.HandleFunc("/testego", checkLoginUnico)
	//http.ListenAndServe(":8822", nil)
}

func checkLoginUnico(writer http.ResponseWriter, req *http.Request) {
	var loginUnicoURL string = os.Getenv("loginUnicoUrl")

	client := &http.Client{
		CheckRedirect: nil,
	}
	request, err := http.NewRequest("GET", loginUnicoURL, nil)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	var token string = req.Header.Get("Authorization")
	request.Header.Add("Authorization", "Bearer " + token)

	response, err := client.Do(request)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		var redirectURL string = os.Getenv("redirectSucessUrl")
		http.Redirect(writer, req, redirectURL, http.StatusSeeOther)
		return
	} else {
		log.Println("login falhou com o token: " + token)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
}
