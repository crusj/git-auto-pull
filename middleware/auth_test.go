package middlware

import (
	"github.com/wonderivan/logger"
	"testing"
)

var (
	token *Token = &Token{}
)

func TestToken_CreateToken(t *testing.T) {
	tokenString, err := token.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("生成TOKEN为:%s\n", tokenString)
}
func TestToken_DecodeToken(t *testing.T) {
	tokenString, _ := token.CreateToken(1)
	id, err := token.DecodeToken(tokenString)
	if err != nil {
		t.Fatal(err)
	}
	if id != 1 {
		t.Fatal("ID解析错误")
	}
	logger.Info("ID为", id)
}
