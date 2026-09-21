package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	// translate password to []byte and translate this object to hash
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	// return hash
	return string(bytes), err
}

func verifyPassword(hash, password string) bool {
	// // translate hash and password to []byte, if hash == password => return true
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Claims object
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func generateToken(userID int) (string, error) {
	// create a token claims
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// token active time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// create a new jwtToken with claims and one of Methods. token == object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return string token
	return token.SignedString(jwtSecret)
}

func verifyToken(tokenString string) (int, error) {
	// parse token to object
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// validation token and return UserID
	return claims.UserID, nil
}

/*
Клиент регистрируется: POST /register с пароль "password123"

Сервер: hashPassword("password123") → хеш вроде $2a$10$...

Сервер: сохраняет хеш в БД

Сервер: generateToken(1) → токен "eyJhbGc..."

Клиент получает токен и сохраняет

Клиент логинится: POST /login с email и пароль "password123"

Сервер: достаёт хеш из БД

Сервер: verifyPassword($2a$10$..., "password123") → true

Сервер: generateToken(1) → новый токен

Клиент получает токен

Клиент удаляет себя: DELETE /users/1 с заголовком Authorization: Bearer eyJhbGc...

Middleware: verifyToken("eyJhbGc...") → userID=1

Middleware: проверяет, что userID=1 совпадает с id=1 в URL

Middleware: вызывает deleteUser()

Хендлер удаляет пользователя
*/
