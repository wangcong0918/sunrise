package jwt
import (
	"log"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"time"
	"errors"
)

const (
	SecretKey = "sunrise"
)
type User struct {
	UserID  string `json:"userID"`
	Phone string  `json:"phone"`
	OpenID  string  `json:"openID"`
	UserName    string `json:"userName"`
	Address string  `json:"address"`
	Remark   string `json:"remark"`
}

//自定义payload结构体
type userStdClaims struct {
	jwt.StandardClaims
	*User
}

func JwtGenerateToken(u *User,d time.Duration) (string, error) {
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        u.UserID,
		Issuer:    SecretKey,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		User:           u,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Logger.Info("config is wrong, can not generate jwt------------->", err)
	}
	return tokenString,err
}
func GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
//JwtParseUser 解析payload的内容,得到用户信息
//gin-middleware 会使用这个方法
func JwtParseUser(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := userStdClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claims.User, err
}