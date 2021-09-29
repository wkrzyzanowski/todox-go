package authentication

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

const USER = "admin"
const ACCESS_TOKEN_TYPE = "ACCESS"
const ACCESS_TOKEN_DURATION_MINUTES = time.Minute * 5
const REFRESH_TOKEN_TYPE = "REFRESH"
const REFRESH_TOKEN_DURATION_MINUTES = time.Minute * 60

type AuthorizationToken struct {
	AccessToken  string
	RefreshToken string
}

type AuthenticationSecret struct {
	Secret string `json:"key"`
}

func GetAuthentication(user string, password string) (AuthorizationToken, error) {
	userPassword := encode(password)
	secretKey := readSecretJson()
	if userPassword == secretKey {
		log.Println("Authenticated successfully!")
		at, errAccess := createAccessToken(user, secretKey)
		rt, errRefresh := createRefreshToken(user, secretKey)
		if errAccess != nil || errRefresh != nil {
			return AuthorizationToken{}, errors.New("token creation failed")
		}
		return AuthorizationToken{
			AccessToken:  at,
			RefreshToken: rt,
		}, nil
	} else {
		log.Println("Authentication failure!")
		return AuthorizationToken{}, errors.New("authentication failed")
	}
}

func RefreshAuthentication(refreshToken string) (AuthorizationToken, error) {
	parsedToken, err := VerifyToken(refreshToken)
	if err != nil {
		return AuthorizationToken{}, err
	}
	userId := parsedToken.Claims.(jwt.MapClaims)["userId"]
	at, errAccess := createAccessToken(fmt.Sprintf("%v", userId), readSecretJson())
	if errAccess != nil {
		return AuthorizationToken{}, errors.New("token creation failed")
	}
	return AuthorizationToken{
		AccessToken:  at,
		RefreshToken: ExtractToken(refreshToken),
	}, nil
}

func VerifyToken(authorizationToken string) (*jwt.Token, error) {
	tokenString := ExtractToken(authorizationToken)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(readSecretJson()), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(token string) string {
	//normally Authorization the_token_xxx
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func createAccessToken(userId string, secret string) (string, error) {
	return createToken(userId, secret, ACCESS_TOKEN_DURATION_MINUTES, ACCESS_TOKEN_TYPE)
}

func createRefreshToken(userId string, secret string) (string, error) {
	return createToken(userId, secret, REFRESH_TOKEN_DURATION_MINUTES, REFRESH_TOKEN_TYPE)
}

func createToken(userId string, secret string, duration time.Duration, tokenType string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["userId"] = userId
	atClaims["exp"] = time.Now().Add(duration).Unix()
	atClaims["type"] = tokenType
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func encode(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Default secret key is "admin"
func readSecretJson() string {
	jsonFile, err := os.Open("secret.json")
	if err != nil {
		log.Fatalf("Cannot read sectet.json file. Error: %v", err)
	}
	var byteValue, _ = ioutil.ReadAll(jsonFile)
	var authenticationSecret AuthenticationSecret
	json.Unmarshal(byteValue, &authenticationSecret)
	defer jsonFile.Close()
	return authenticationSecret.Secret
}
