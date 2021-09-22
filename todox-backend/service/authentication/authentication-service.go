package authentication

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type AuthenticationSecret struct {
	Secret string `json:"key"`
}

func GetAuthentication(password string) bool {
	actual := encode(password)
	secretKey := readSecretJson().Secret
	if actual == secretKey {
		log.Println("Authenticated successfully!")
		return true
	} else {
		log.Println("Authentication failure!")
		return false
	}
}

func encode(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Default secret key is "admin"
func readSecretJson() AuthenticationSecret {
	jsonFile, err := os.Open("secret.json")
	if err != nil {
		log.Fatalf("Cannot read sectet.json file. Error: %v", err)
	}
	var byteValue, _ = ioutil.ReadAll(jsonFile)
	var authenticationSecret AuthenticationSecret
	json.Unmarshal(byteValue, &authenticationSecret)
	defer jsonFile.Close()
	return authenticationSecret
}
