package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const TokenExpireDuration = time.Hour*2
var MySercet = []byte("onethreeonethreeoneonethree")

type MyClaims struct {
	UserID int64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GetToken 生成JWT
func GetToken(userID int64,username string) (string, error) {
	// 创建一个我们自己的声明
	claims := MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "hew1s", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySercet)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString,mc, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return MySercet, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}