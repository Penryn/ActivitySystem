package jwt

import (
	"errors"
	"strings"
	"time"

	"activitySystem/internal/global"
	"github.com/golang-jwt/jwt/v5"
)

var (
	key string
	t   *jwt.Token
)

// NewJWT 生成 JWT
func NewJWT(uid uint) string {
	key = global.Config.GetString("jwt.secret")
	duration := time.Hour * 24 * 7
	expirationTime := time.Now().Add(duration).Unix() // 设置过期时间
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": expirationTime,
	})
	s, err := t.SignedString([]byte(key))
	if err != nil {
		return ""
	}
	return s
}

// ParseJWT 解析 JWT 并返回 stuId
func ParseJWT(authHeader string) (uint, error) {
	const bearerPrefix = "Bearer "

	// 解析 Bearer 令牌
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return 0, errors.New("invalid authorization header")
	}
	token := strings.TrimPrefix(authHeader, bearerPrefix)
	secret := global.Config.GetString("jwt.secret")

	t, err := jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !t.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// 验证 exp 是否有效
	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		return 0, errors.New("token expired")
	}

	uid, ok := claims["uid"].(float64)
	if !ok {
		return 0, errors.New("invalid user claims")
	}

	// 将 float64 转换为 uint
	return uint(uid), nil

}
