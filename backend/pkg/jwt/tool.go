package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Code string
	Data map[string]string
	jwt.StandardClaims
}

type Tool struct {
	Key string
}

func NewTool(key string) *Tool {
	return &Tool{
		Key: key,
	}
}

// NewClaims 创建认证数据
func (j *Tool) NewClaims(code string, data map[string]string, expires time.Duration) Claims {
	return Claims{
		Code: code,
		Data: data,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",                             // 接收jwt的客户端或者其他
			ExpiresAt: time.Now().Add(expires).Unix(), // 过期时间
			Id:        code,                           // 唯一识别，一般为Id
			IssuedAt:  time.Now().Unix(),              // 签发时间
			Issuer:    "gostc",                        // 签发者
			NotBefore: time.Now().Unix(),              // 此时间之前失效
			Subject:   code,                           // 面向的用户
		},
	}
}

// GenerateToken 生成Token
func (j *Tool) GenerateToken(claims Claims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Key))
	return token, err
}

// ValidToken 验证Token
func (j *Tool) ValidToken(token string) (myClaims Claims, err error) {
	withClaims, err := jwt.ParseWithClaims(token, &myClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Key), nil
	})
	if err != nil {
		return myClaims, err
	}
	if withClaims != nil {
		if validClaims, ok := withClaims.Claims.(*Claims); ok && withClaims.Valid {
			return *validClaims, nil
		}
	}
	return myClaims, errors.New("已失效")
}
