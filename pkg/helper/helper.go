package helper

import (
	"errors"
	"fmt"
	"main/pkg/domain"
	"main/pkg/utils/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// AuthCustomClaims represents custom claims for JWT
type AuthCustomClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateTokensAdmin(admin domain.Admin) (string, string, error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    uint(admin.ID),
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshTokenClaims := &AuthCustomClaims{
		Id:    uint(admin.ID),
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(viper.GetString("KEY")))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(viper.GetString("KEY")))
	if err != nil {
		return "", "", err
	}

	fmt.Println("Admin tokens created")
	return accessTokenString, refreshTokenString, nil
}

func GenerateTokensUser(user models.UserResponse) (string, string, error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    uint(user.Id),
		Email: user.Email,
		Role:  "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshTokenClaims := &AuthCustomClaims{
		Id:    uint(user.Id),
		Email: user.Email,
		Role:  "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(viper.GetString("KEY")))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(viper.GetString("KEY")))
	if err != nil {
		return "", "", err
	}

	fmt.Println("User tokens created")
	return accessTokenString, refreshTokenString, nil
}

/*
validateToken is for decrypting a jwt token using HMAC256 algorithm

Parameters:
- token: JWT token string.
*/
func ValidateToken(token string) (*jwt.Token, error) {
	fmt.Println("Token validating.........")
	jwttoken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		secret := viper.GetString("KEY")
		return []byte(secret), nil
	})

	return jwttoken, err
}

// using for generating tokens when access token expires

func TokensFromRefreshToken(prevRefreshTokenString string) (string, string, error) {
	// Parse the previous refresh token
	prevRefreshToken, err := jwt.Parse(prevRefreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("KEY")), nil
	})

	if err != nil {
		return "", "", err
	}

	// Extract claims from the previous refresh token
	prevRefreshClaims, ok := prevRefreshToken.Claims.(jwt.MapClaims)
	if !ok || !prevRefreshToken.Valid {
		return "", "", errors.New("invalid refresh token")
	}

	// Use the claims to generate a new access token
	newAccessTokenClaims := &AuthCustomClaims{
		Id:    uint(prevRefreshClaims["id"].(float64)),
		Email: prevRefreshClaims["email"].(string),
		Role:  prevRefreshClaims["role"].(string),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newAccessTokenClaims)
	newAccessTokenString, err := newAccessToken.SignedString([]byte(viper.GetString("KEY")))
	if err != nil {
		return "", "", err
	}

	// Generate a new refresh token for the next cycle
	newRefreshTokenClaims := &AuthCustomClaims{
		Id:    uint(prevRefreshClaims["id"].(float64)),
		Email: prevRefreshClaims["email"].(string),
		Role:  prevRefreshClaims["role"].(string),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newRefreshTokenClaims)
	newRefreshTokenString, err := newRefreshToken.SignedString([]byte(viper.GetString("KEY")))
	if err != nil {
		return "", "", err
	}

	return newAccessTokenString, newRefreshTokenString, nil
}

/*
GetUserID returns the userID stored in the context

Parameters:
- c: gin context

Returns:
- int: userID
- error: error is returned
*/
func GetUserID(c *gin.Context) (int, error) {
	var key models.UserKey = "userID"
	val := c.Request.Context().Value(key)

	// Check if the value is not nil
	if val == nil {
		return 0, errors.New("userID not found in context")
	}

	// Use type assertion to convert to the expected type
	userKey, ok := val.(models.UserKey)
	if !ok {
		return 0, errors.New("failed to convert userID to the expected type")
	}

	ID := userKey.String()
	userID, err := strconv.Atoi(ID)
	if err != nil {
		return 0, errors.New("failed to convert userID to int")
	}

	return userID, nil
}

/*
PasswordHashing hashes a password.

Parameters:
- password: Password to be hashed.

Returns:
- string: Hashed Password.
- error: Error is returned if any.
*/
func PasswordHashing(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

func FindMostBoughtProduct(products []domain.ProductReport) []int {

	productMap := make(map[int]int)

	for _, v := range products {
		productMap[v.InventoryID] += v.Quantity
	}

	maxQty := 0
	for _, v := range productMap {
		if v > maxQty {
			maxQty = v
		}
	}

	var bestSellers []int
	for k, v := range productMap {
		if v == maxQty {
			bestSellers = append(bestSellers, k)
		}
	}
	return bestSellers
}
