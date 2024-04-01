package middleware

import (
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/helpers"
	"testing"
)

func TestThatWhenAUserFirstLogsInThatWeCreateAUserRecord(t *testing.T) {
	helpers.ConfigureTests(t)
	db_connection := db.InitDB()

	RetrieveOrCreateClerkUser("aayush.manocha@gmail.com")

	var user db.User
	db_connection.Where("email = ?", "aayush.manocha@gmail.com").First(&user)

	if user.ID == 0 {
		t.Errorf("User not created")
	}

}

func TestThatDuplicateUserRecordsAreNotCreated(t *testing.T) {
	helpers.ConfigureTests(t)
	db_connection := db.InitDB()

	db_connection.Create(&db.User{Email: "aayush.manocha@gmail.com"})

	var initial_user_count int64
	var users []db.User
	db_connection.Find(&users).Count(&initial_user_count)

	RetrieveOrCreateClerkUser("aayush.manocha@gmail.com")

	var final_user_count int64
	db_connection.Find(&users).Count(&final_user_count)

	if initial_user_count != final_user_count {
		t.Errorf("Duplicate user record created")
	}
}
