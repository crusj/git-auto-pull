package middlware

import (
	"errors"
	. "github.com/crusj/git-auto-pull/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	//秘钥
	secret []byte
)

func getSecret() []byte {
	if secret == nil {
		secret = []byte(Cfg.Section("auth").Key("secret").String())
	}
	return secret
}

type Token struct{}
type Claim struct {
	Id        int64 `json:"id"`
	ExpiresAt int64 `json:"exp"`
	jwt.StandardClaims
}

/**
 *生成TOKEN
 */
func (Token) CreateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 30).Unix(),
	})
	return token.SignedString(getSecret())
}

/**
 *解密TOKEN
 */
func (Token) DecodeToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecret()), nil
	})
	if err != nil {
		return 0, nil
	}
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		//检查过期
		if time.Now().Unix() > claims.ExpiresAt {
			return 0, errors.New("令牌过期")
		} else {
			return claims.Id, nil
		}
	} else {
		return 0, nil
	}

}
