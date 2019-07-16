package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const ClientId = "TruAuthDemoResourceServer"
const ClientSecret = "some-client-secret-123"
const AuthServer = "http://localhost:4820"
const GrantType = "authorization_code"
const RedirectURI = "http://localhost:3000/redirectEndpoint"

type TokenRequest struct {
	AuthorizationCode string `json:"authorizationCode"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/token", getToken) // Todo: change to code

	fmt.Println("Server Started on Port 4821")
	http.ListenAndServe(":4821", mux)
}

func getToken(w http.ResponseWriter, req *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	tokenRequest := TokenRequest{}
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&tokenRequest)

	tokenEndpoint := fmt.Sprintf("%s/token?grant_type=%s&code=%s&redirect_uri=%s&client_id=%s&client_secret=%s", AuthServer, GrantType, tokenRequest.AuthorizationCode, RedirectURI, ClientId, ClientSecret)
	response, err := http.Post(tokenEndpoint, "application/json", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := ioutil.ReadAll(response.Body)

	w.Write(body)
}
