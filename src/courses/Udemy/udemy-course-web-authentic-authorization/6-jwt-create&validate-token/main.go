package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//customised claims, instead of session Id it can be anything
//godoc many methods and claims are given for standard claim to use
type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

var key = []byte{}

//customised claim method
func (u *UserClaims) valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("Invalid session ID")
	}
	return nil
}

//among signing methods there are HMAC, RSA and ECDSA from go doc
//HMAC uses only single key while RSA and ECDSA uses two key(public and private)
//private key to send to client and once slient send back public key is use to compare
/*ECDSA is most recommended compare to RSA coz its req smaller key for same amt of power
for the same amount of security and signature and is faster than RSA */
//besides signing method also need claims can use type MapClaims or custom claims like standar claims
func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

}

func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Error in createToken when signing token: %w", err)
	}

	return signedToken, nil
}

//verify token and return token if its verified
func parseToken(signedToken string) (*UserClaims, error) {
	//t below is unverified token and in func (t *jwt.Token)it get verified
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}
		return key, nil

	})
	if err != nil {
		return nil, fmt.Errorf("Error in parseToken while parsing token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("Error in parseToken, token is not valid")
	}
	return t.Claims.(*UserClaims), nil
}
