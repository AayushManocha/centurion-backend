package services

import (
	"testing"
)

func TestClerkGetUserByID(t *testing.T) {

	t.Setenv("CLERK_SECRET_KEY", "sk_test_xOqtz2D5sJqrtVtYGeyWWHfpix1AxJBLm66acYhpFN")
	aayush_test_user_id := "user_2eCl6AS5Sg7GvRgXkteF8I88gsn"
	user, err := GetUserById(aayush_test_user_id)

	if err != nil {
		t.Errorf("Error getting user by ID: %s", err.Error())
	}

	emailAddress := user.EmailAddresses[0].EmailAddress

	if emailAddress != "aayush.manocha@gmail.com" {
		t.Errorf("User email is incorrect")
	}
}
