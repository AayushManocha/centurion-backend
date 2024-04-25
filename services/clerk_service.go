package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func GetClerkUserById(id string) (*clerk.User, error) {
	client := http.Client{}

	fmt.Println("Getting user from clerk with ID: " + id)

	req, err := http.NewRequest("GET", "https://api.clerk.dev/v1/users/"+id, nil)
	if err != nil {
		fmt.Println("Error creating request: " + err.Error())
	}

	authToken := os.Getenv("CLERK_SECRET_KEY")
	authToken = strings.Replace(authToken, "\n", "", -1)
	fmt.Printf("Auth token: %s\n", authToken)
	authHeader := "Bearer " + authToken
	req.Header.Add("Authorization", authHeader)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting user: " + err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading user: " + err.Error())
	}

	// Print body as string
	fmt.Println(string(body))

	// Parse body into Clerk.User

	var user *clerk.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		fmt.Println("Error parsing user: " + err.Error())
		return nil, err
	}

	return user, nil

}

func GenerateSessionToken() (*clerk.Session, error) {
	clerk_api_key := os.Getenv("CLERK_SECRET_KEY")

	client, err := clerk.NewClient(clerk_api_key)

	if err != nil {
		fmt.Println(clerk_api_key)
		fmt.Println("Error creating Clerk client: " + err.Error())
		return nil, err
	}

	session, err := client.Sessions().Read("test-token")

	if err != nil {
		return nil, err
	}

	return session, nil
}
