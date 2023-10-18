package main

import (
	"fmt"
	"log" // Log package for logging error messages
	"os"  // OS package used for interacting with the OS such as reading files

	"github.com/golang-jwt/jwt/v5" // Importing the JWT library for Go to handle JSON Web Tokens
)

// JWT token given as string
var tkn = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTY5NzYyMzI5MCwiaWF0IjoxNjk3NjIwMjkwfQ.QZ642KlWsLbTxpUw5TrLkWfI4-gu9mepYIbQmnmfqq47qQMS51P-MoKMeU6i9_QHlkoQLEiRNE-jIfHTKLh8vwGdpbSJnYH7awJzQIZWxHWtFjIkAg05exPDa4oonEyAh1VbsNuMykeScQSrqq2oyKkiqK5JeWGHLP5YkngtZDnn24n_p--YrnN-lnUDYNM6G72igFPxGF-PCEWOmngwFwgAjRGKMlwJqnB-dYUaI3-1cJQlfjWbmUCN26i2Zf1MSpIjezcRMr__BEv_A0jxCooYUIDsz64sgxZHY5D5j0XMOu07D1hY7Jg2TuQOr6orqHesPqHE9sF3Rf5TlNPHGw`

func main() {
	// Reads the public key from pubkey.pem file
	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		// If there's an error reading the file, print an error message and stop execution
		log.Fatalln("not able to read pem file")
	}

	// Parse the read public key to RSA public key format
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)
	if err != nil {
		// If there's an error parsing the public key, log the error and stop execution
		log.Fatalln(err)
	}
	var c jwt.RegisteredClaims
	// Parsing the JWT token with the claims
	token, err := jwt.ParseWithClaims(tkn, &c, func(token *jwt.Token) (interface{}, error) {
		// Provides the public key for validating the JWT token
		return publicKey, nil
	})

	if err != nil {
		// If error while parsing the token, print the error and exit
		log.Println("parsing token", err)
		return

	}
	if !token.Valid {
		// If the token is not valid, log the error and exit
		log.Println("invalid token")
		return
	}

	// Print the claims from the token
	fmt.Printf("%+v", c)

}
