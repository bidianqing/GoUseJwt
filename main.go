package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

/*
iss（全称为 issuer），指明 JWT 是由谁签发的
sub（全称为 subject），指明 JWT 的主题（也可理解为面向用户的类型）
aud（全称为 audience），指明 JWT 希望谁签收
exp（全称为 expiration time），指明 JWT 的过期时间，过期时间需大于签发时间
nbf（全称为 not before time），指明 JWT 在哪个时间点生效
iat（全称为 issued at time），指明 JWT 的签发时间
jti（全称为 JWT ID），指明 JWT 唯一 ID，用于避免重放攻击
//*/

func main() {
	key := "hellojwthellocsharpbidianqingbidianqing"

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "bidianqing",
		"aud":  "golang",
		"sub":  "auth",
		"exp":  jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
		"nbf":  jwt.NewNumericDate(time.Now()),
		"iat":  jwt.NewNumericDate(time.Now()),
		"jti":  uuid.New(),
		"name": "1",
	})
	token, err := t.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)

	decodeToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(decodeToken.Claims.(jwt.MapClaims))
}
