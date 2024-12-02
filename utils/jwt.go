package utils

import (
	"fmt"
	"gdsc-core-be/database"
	"gdsc-core-be/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJwt(user models.User, expiration time.Time, tokenType string) (string, error) {
	SecretKey, err := database.SecretKeyInit()
	if err != nil {
		log.Fatal(err)
	}

	batch := models.Batch{}
	member := models.Member{}
	role := models.Role{}

	database.DB.Select("id_batch").Where("year = ?", user.CurrentBatch).First(&batch)
	database.DB.Select("id_role, id_member").Where("id_user = ? AND id_batch = ?", user.IDUser, batch.IDBatch).First(&member)
	database.DB.Select("nama").Where("id_role = ?", member.IDRole).First(&role)

	claims := jwt.MapClaims{
		"id":    user.IDUser,
		"id_member": member.IDMember,
		"role":  role.Nama,
		"batch": user.CurrentBatch,
		"exp":   expiration.Unix(),
		"type":  tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func GenerateNewJwt(id uuid.UUID, id_member uuid.UUID, role string, batch int, expiration time.Time, tokenType string) (string, error) {
	SecretKey, err := database.SecretKeyInit()
	if err != nil {
		fmt.Println("Error fetching secret key:", err)
		return "", err
	}

	claims := jwt.MapClaims{
		"id":    id,
		"id_member": id_member,
		"role":  role,
		"batch": batch,
		"exp":   expiration.Unix(),
		"type":  tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	return tokenString, nil
}

func VerifyJwt(tokenString string) error {
	SecretKey, err := database.SecretKeyInit()
	if err != nil {
		return err
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return err
	}

	return nil
}

func VerifyJwtWithRole(tokenString string) (*jwt.Token, error) {
	SecretKey, err := database.SecretKeyInit()
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeJwt(tokenString string) error {
	err := VerifyJwt(tokenString)
	if err != nil {
		return err
	}
	return nil
}

func DecodeJwtWithRole(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyJwtWithRole(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if !isOk || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func GenerateTokens(user models.User) (string, string, error) {
	accessToken, err := GenerateJwt(user, time.Now().Add(30*time.Hour), "access")
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateJwt(user, time.Now().Add(24*time.Hour), "refresh")
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}