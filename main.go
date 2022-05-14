package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "gopkg.in/square/go-jose.v2"
	// "gopkg.in/square/go-jose.v2"
	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	// "gopkg.in/square/go-jose.v2"
	// "gopkg.in/square/go-jose.v2/jwt"
)

// type KeyRes struct {
// 	Kty string      `json:"kty"`
// 	Kid string      `json:"kid"`
// 	Use string      `json:"use"`
// 	Alg string      `json:"alg"`
// 	N   interface{} `json:"n"`
// 	E   string      `json:"e"`
// }

// type KeysObj struct {
// 	Keys []KeyRes `json:"keys"`
// }

const token = "eyJraWQiOiJXNldjT0tCIiwiYWxnIjoiUlMyNTYifQ.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiaG9zdC5leHAuRXhwb25lbnQiLCJleHAiOjE2NTI2MzIyMzYsImlhdCI6MTY1MjU0NTgzNiwic3ViIjoiMDAwNzkwLjE3YzA2YjBiZTljMjQ2MjI5YzY5ZDcyMGFjZjE2MmQ4LjAyNDUiLCJjX2hhc2giOiJ4NVVSN080WjVkeTVmcnlQdzloTGh3IiwiZW1haWwiOiJqcGVyc2luZzE5OTlAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNjUyNTQ1ODM2LCJub25jZV9zdXBwb3J0ZWQiOnRydWV9.dWWb2BB14y-fzsnmT3qz_ZIOOy2Ea4Sc8GcS--wJb7o-fKWFWnRiN4t9VNvatJgmPu-1gpYEhqE7F8rd9xS8pDHvVCmlt5YSQLiWAITmxqlm5-xrhrX_Cyddi2EQwEg5ITv8nc9n-Ryv1XVxdctNzmFdfzFRWRYBRmFf4SQ4T7QXUs-GU_vBWS6fssjbOwguUMIPKCUS3XyyrzFUSdeWrsmR457iV-GH3mCkblvV5YvnjlneRTZVK2o0QTSLgClFIeIQ9LORNPIcYGRPmqEWxYIxVlGUGnbLUlxVFjbqQ1s1aF-GviFBZx-xYeaWGziS4_hwFQLBP0HiBBHSirK6Xg"
const jwksURL = "https://appleid.apple.com/auth/keys"

func main() {

	//"gopkg.in/square/go-jose.v2/jwt"
	// View claims without checking validity

	// var claims map[string]interface{}

	// token, _ := jwt.ParseSigned("eyJraWQiOiJXNldjT0tCIiwiYWxnIjoiUlMyNTYifQ.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiaG9zdC5leHAuRXhwb25lbnQiLCJleHAiOjE2NTI2MzIyMzYsImlhdCI6MTY1MjU0NTgzNiwic3ViIjoiMDAwNzkwLjE3YzA2YjBiZTljMjQ2MjI5YzY5ZDcyMGFjZjE2MmQ4LjAyNDUiLCJjX2hhc2giOiJ4NVVSN080WjVkeTVmcnlQdzloTGh3IiwiZW1haWwiOiJqcGVyc2luZzE5OTlAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNjUyNTQ1ODM2LCJub25jZV9zdXBwb3J0ZWQiOnRydWV9.dWWb2BB14y-fzsnmT3qz_ZIOOy2Ea4Sc8GcS--wJb7o-fKWFWnRiN4t9VNvatJgmPu-1gpYEhqE7F8rd9xS8pDHvVCmlt5YSQLiWAITmxqlm5-xrhrX_Cyddi2EQwEg5ITv8nc9n-Ryv1XVxdctNzmFdfzFRWRYBRmFf4SQ4T7QXUs-GU_vBWS6fssjbOwguUMIPKCUS3XyyrzFUSdeWrsmR457iV-GH3mCkblvV5YvnjlneRTZVK2o0QTSLgClFIeIQ9LORNPIcYGRPmqEWxYIxVlGUGnbLUlxVFjbqQ1s1aF-GviFBZx-xYeaWGziS4_hwFQLBP0HiBBHSirK6Xg")
	// _ = token.UnsafeClaimsWithoutVerification(&claims)
	// fmt.Printf(token)

	//Valid token

	res, err := http.Get("https://appleid.apple.com/auth/keys")
	if err != nil {
		fmt.Println("error making http req")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading body")
	}

	jwks, err := keyfunc.NewJSON(body)

	token, err := jwt.Parse("eyJraWQiOiJXNldjT0tCIiwiYWxnIjoiUlMyNTYifQ.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiaG9zdC5leHAuRXhwb25lbnQiLCJleHAiOjE2NTI2Mzk2MjIsImlhdCI6MTY1MjU1MzIyMiwic3ViIjoiMDAwNzkwLjE3YzA2YjBiZTljMjQ2MjI5YzY5ZDcyMGFjZjE2MmQ4LjAyNDUiLCJjX2hhc2giOiJrOUhMUUxib3k5S1dFR09wYXZvLWFRIiwiZW1haWwiOiJqcGVyc2luZzE5OTlAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNjUyNTUzMjIyLCJub25jZV9zdXBwb3J0ZWQiOnRydWV9.LvgcXxZkmDxwQlpcPvRYIEbaiqgoMxAgBhxMGSmmJAQP4U9ksq4v8lXM88WEVkzmRDzwO07KgrjrTcPbBSBPvDHCPWGv8h0bho1tzs11nNf-pur0Qgbh8GInDG_WGCH-MBjohGGYHSu-wo9cU4i2ZdsglZmv_KaOc0uKc021I9z2DAEYfR_W6xHVJaT_dhQeNk566nocZGKO5BdR6aocxfuL582bxUbBIw57I6XCNui3Kj4DJse9SfNWakv0k8mmi6kx39qHwKyenDSTEapQC-_U3MxzeljzoN3IgjLiswDVlbn07jTxBQGEV-N4RfXfXivfppAIxvpizkO8KtnGGA", jwks.Keyfunc)

	if err != nil {
		log.Fatalf("Failed to parse the JWT.\nError: %s", err.Error())
	}

	// Check if the token is valid.
	if !token.Valid {
		log.Fatalf("The token is not valid.")
	}

	if token.Valid {
		fmt.Println("TOKEN IS VALID")
	}
	fmt.Println(token.Claims)

}
