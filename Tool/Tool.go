package Tool

import (
	"demo/Global"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)


func KeyRand(n int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
 
	var sb strings.Builder
	for i := 0; i < n; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}

var jwtKey = []byte(Global.Config.Jwt.Key)

func Token(U Global.User) (string,error){
	claims := &Global.Claims{
		// 使用 ID、Username 作为有效载荷
		Email: U.Email,
		Password: U.Password,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(Global.Config.Jwt.Expires), // 签名过期时间
			NotBefore: time.Now().Unix() - 1000,                      // 签名生效时间
			Issuer:    Global.Config.Jwt.Issuer,                       // 签名发行人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Global.Claims, error) {
	claims := &Global.Claims{}
	// 解码
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}