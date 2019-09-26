package main

import (
	"log"
	"net/http"

	//"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"
)

type token struct {
	Token string `json:"token"`
}

func main() {
	lambda.Start(checkLoginUnico)

	//http.HandleFunc("/testego", checkLoginUnico)
	//http.ListenAndServe(":8822", nil)
}

//func checkLoginUnico(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
func checkLoginUnico(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//var loginUnicoURL string = os.Getenv("LoginUnicoURL")
	//log.Println(loginUnicoURL)
	//fmt.Println(loginUnicoURL)
	var loginUnicoURL string = "http://loginunico-dev.sa-east-1.elasticbeanstalk.com/loginunico/token"

	client := &http.Client{
		CheckRedirect: nil,
	}

	request, err := http.NewRequest("GET", loginUnicoURL, nil)
	if err != nil {
		//writer.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		//return
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 401}, nil
	}
	//localhost
	//decoder := json.NewDecoder(req.Body)
	//decoder := json.NewDecoder(req.Body)
	//var token string = req.Header.Get("Authorization")
	token := token{Token: ""}

	//lambda
	err = json.Unmarshal([]byte(req.Body), &token)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 401}, nil
	}

	//err = decoder.Decode(&token)
	var tokenString string = token.Token
	request.Header.Add("Authorization", "Bearer "+tokenString)

	response, err := client.Do(request)
	if err != nil {
		//writer.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		//return
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 401}, nil
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		//var redirectURL string = os.Getenv("redirectSucessURL")
		//log.Println(redirectURL)
		//fmt.Println(redirectURL)
		var redirectURL string = "https://onwidy7zdj.execute-api.us-east-1.amazonaws.com/default/Bucsan"
		var writer http.ResponseWriter
		//http.Redirect(writer, req, redirectURL, http.StatusSeeOther)
		http.Redirect(writer, request, redirectURL, http.StatusSeeOther)
		log.Println("token: " + tokenString)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
	}// else {
		log.Println("login falhou com o token: " + tokenString)
		//writer.WriteHeader(http.StatusUnauthorized)
		//return
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 401}, nil
	//}
}
