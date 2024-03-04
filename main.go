package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "1",
	})
	token, err := t.SignedString([]byte("hellogolang"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)

	decodeToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("hellogolang"), nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(decodeToken.Claims.(jwt.MapClaims)["name"])
}
