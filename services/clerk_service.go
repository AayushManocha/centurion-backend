package services

import (
	"fmt"
	"os"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func GetUserById(id string) (*clerk.User, error) {
	clerk_api_key := os.Getenv("CLERK_SECRET_KEY")

	client, err := clerk.NewClient(clerk_api_key)

	if err != nil {
		fmt.Println(clerk_api_key)
		fmt.Println("Error creating Clerk client: " + err.Error())
		return nil, err
	}
	user, err := client.Users().Read(id)

	if err != nil {
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
