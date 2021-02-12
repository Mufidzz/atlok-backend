package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var (
	signingKey = []byte("07306E5B9F65518D8485B4744EF36C6FD5806F30BBEAB9CBA6FCB954A2405FE7B54A7FE539B77AD8F235DF0421890F6DA2AA07CDB7D1B5CACC550DA8E5B02D75")
)

func Validate(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, getSigningKey)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid jwt token")
	}

	return token, nil
}

func Translate(token *jwt.Token) (Header, Body) {
	h := Header{
		Alg: token.Header["alg"].(string),
		Typ: token.Header["typ"].(string),
	}

	claims := token.Claims.(jwt.MapClaims)

	b := Body{
		Uid: parseUintFromFloat(claims["uid"].(float64)),
		Iat: int64(parseUintFromFloat(claims["iat"].(float64))),
		Acs: parseUintFromFloat(claims["acs"].(float64)),
	}

	return h, b
}

func Generate(body Body) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	c := token.Claims.(jwt.MapClaims)

	c["uid"] = body.Uid
	c["iat"] = time.Now().Unix()
	c["acs"] = body.Acs

	ts, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return ts, nil
}

func getSigningKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("get signing key failed")
	}
	return signingKey, nil
}

func parseUintFromFloat(f float64) uint {
	return parseUint(fmt.Sprintf("%.0f", f))
}

func parseUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
