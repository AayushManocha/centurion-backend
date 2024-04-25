package middleware

import (
	"AayushManocha/centurion/centurion-backend/db"
	"AayushManocha/centurion/centurion-backend/services"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func verifyJWT(tokenString string) (jwt.Claims, error) {
	secretKey := `-----BEGIN PUBLIC KEY-----
	MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwecEPNjM4W5ZZYMXcZ/J
	k5LI6fOFHPztHdn6yfSY1cN3AjdG4pCoBpGPtsJGHz2dvAx6rP4UnH22p4JsGkU5
	GYYJQ84RBdd3AMn0e5i09GD8M1BHAqKaZj6ShBmq+6xb/6nmm1MsgnVtAtwTrlHX
	8D7r5peyBMBTuX/XH9zNdMHcMlWGCiPhT4u8NdiQOV+8LH4dM4FHcUri+BOQ0rX7
	3GntM6YmfH9hqc/fsT7xnHEGjjaOE99SXtyjtHIZ7oexdw5EqT3WnoVX7pFBmwwF
	AQXFqBnV7YbtvfWUkBK6J+HDh9mcLzxk+FqiZQIOf+KuMrZIEEpszVt0ciJ9yZqf
	ewIDAQAB
	-----END PUBLIC KEY-----`

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	return token.Claims, err
}

func getUserFromJWT(token string) (string, error) {
	claims, _ := verifyJWT(token)
	if claims == nil {
		return "", fmt.Errorf("invalid token")
	}
	return claims.(jwt.MapClaims)["sub"].(string), nil
}

func RetrieveOrCreateClerkUserFromDatabase(userEmail string) db.User {
	db_conn := db.GetDB()
	current_user := db.User{}

	db_conn.Where("email = ?", userEmail).First(&current_user)
	if current_user.ID == 0 {
		current_user = db.User{Email: userEmail}
		db_conn.Create(&current_user)
		db_conn.Where("email = ?", userEmail).First(&current_user)
	}

	return current_user
}

func AuthenticatedUser(c *fiber.Ctx) (db.User, error) {
	token := c.Get("Authorization")
	if token == "" {
		return db.User{}, fmt.Errorf("no token provided")
	}

	token = token[7:]

	fmt.Println("Token: ", token)

	if os.Getenv("ENVIRONMENT") == "testing" {
		test_user := db.User{}
		db_conn := db.InitDB()
		db_conn.Where("email = ?", "aayush.manocha@gmail.com").First(&test_user)
		fmt.Printf("Test user: %+v\n", test_user)
		return test_user, nil
	}

	current_user := db.User{}
	userId, err := getUserFromJWT(token)

	fmt.Println("User ID: ", userId)

	if err == nil {
		// Use Clerk client to get user email
		user, err := services.GetClerkUserById(userId)
		if err != nil {
			return current_user, err
		}
		userEmail := user.EmailAddresses[0].EmailAddress
		current_user = RetrieveOrCreateClerkUserFromDatabase(userEmail)
	}

	return current_user, err
}

func EnsureAuthenticated(c *fiber.Ctx) error {
	fmt.Println("Checking for auth")
	_, err := AuthenticatedUser(c)
	if err != nil {
		fmt.Println("Unauthorized on request: ", c.Path())
		fmt.Println("error: ", err.Error())
		return c.Status(401).SendString("Could not authorize user")
	}
	return c.Next()
}
