package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/config"
)

func GenerateToken(claim dto.UserContext) (string, error) {
	secretKey := []byte(config.AppConfig.JWT_SECRET)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":            float64(claim.ID),
		"email":         claim.Email,
		"roles":         convertUintSliceToFloat64(claim.Roles),
		"selected_role": float64(claim.SelectedRole),
		"exp":           GetExpiryTime().Unix(),
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(config.AppConfig.JWT_SECRET)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok {
			return nil, fmt.Errorf("token expiration (exp) is missing or invalid")
		}
		if time.Now().Unix() > int64(exp) {
			return nil, fmt.Errorf("token has expired")
		}
	} else {
		return nil, fmt.Errorf("invalid token claims")
	}

	return token, nil
}

func convertUintSliceToFloat64(roles []uint) []float64 {
	var result []float64
	for _, r := range roles {
		result = append(result, float64(r))
	}
	return result
}

func GetExpiryTime() time.Time {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local)
	expiry := time.Until(midnight)
	return now.Add(expiry)
}

func SetTokenCookie(c *fiber.Ctx, token string) {

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  GetExpiryTime(),
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: fiber.CookieSameSiteStrictMode,
		Path:     "/",
	})
}

func DeleteTokenCookie(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: fiber.CookieSameSiteStrictMode,
		Path:     "/",
	})
}
