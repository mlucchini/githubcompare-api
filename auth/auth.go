package auth

import (
	"os"
	"encoding/json"
	"log"
)

type credentials struct {
	githubClientId string
	githubClientSecret string
}

var cred = &credentials{}

func readCredentials() {
	file, _ := os.Open("credentials")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&cred)
	if err != nil {
		log.Print("Error: unable to load Github credentials")
	}
}

func GithubClientId() string {
	if cred == nil {
		readCredentials()
	}
	return cred.githubClientId
}

func GithubClientSecret() string {
	if cred == nil {
		readCredentials()
	}
	return cred.githubClientSecret
}

func GithubUrlSuffix() string {
	return "?client_id=" + GithubClientId() + "&client_secret=" + GithubClientSecret()
}