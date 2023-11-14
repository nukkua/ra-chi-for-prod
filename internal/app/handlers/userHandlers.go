package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nukkua/ra-chi/internal/app/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JwtKey = []byte("secret_key")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GetUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User
		db.Find(&users)
		json.NewEncoder(w).Encode(users)
	}
}

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userWithoutHash models.UserWithoutHash

		err := json.NewDecoder(r.Body).Decode(&userWithoutHash)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		hashedPassword, errorHashing := bcrypt.GenerateFromPassword([]byte(userWithoutHash.Password), 1)
		if errorHashing != nil {
			http.Error(w, errorHashing.Error(), http.StatusInternalServerError)
			return
		}

		user := models.User{
			Username: userWithoutHash.Username,
			Email:    userWithoutHash.Email,
			Password: hashedPassword,
		}
		result := db.Create(&user)

		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "user created"})

	}
}
func LoginUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&creds)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := db.Where("email = ?", creds.Email).First(&user)

		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusUnauthorized)
			return
		}
		error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
		if error != nil {
			http.Error(w, error.Error(), http.StatusUnauthorized)
			return
		}
		expirationTime := time.Now().Add(30 * time.Minute)

		claims := &Claims{
			Email: creds.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(JwtKey)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"access_token": tokenString, "username": user.Username, "email": user.Email})
	}
}
