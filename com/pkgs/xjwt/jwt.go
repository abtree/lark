package xjwt

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
)

// token 结构
type JWTToken struct {
	Token    string //token的验证信息
	Expire   int64  //过期时间
	UserData []byte //用户数据
}

/*
Payload载荷：
jti：该jwt的唯一标识
iss：该jwt的签发者
iat：该jwt的签发时间
aud：该jwt的接收者
sub：该jwt的面向的用户
nbf：该jwt的生效时间,可不设置,若设置,一定要大于当前Unix UTC,否则token将会延迟生效
exp：该jwt的过期时间
*/

// JWT使用的常量值
const (
	JWT_ISSUER           = "lark.com"
	JWT_TOKEN_SECRET_KEY = "lark_jwt_token_2022"
	JWT_PREFIX           = "jwt="
)

// JWT使用的Key
const (
	JWT_KEY_ISS = "iss" //签发者
	JWT_KEY_EXP = "exp" //过期时间
	JWT_KEY_IAT = "iat" //生效时间

	//用户自定义
	USER_DATA = "user_data" //用户数据
)

/*
创建JWTToken

access 是否为返回参数(设置为返回参数 会加上JWT_PREFIX)
expireIn 过期时间
*/
func CreateToken(user []byte, access bool, expireIn int) (t *JWTToken, err error) {
	//创建JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	//拼装jwt body 数据
	userStr := base64.StdEncoding.EncodeToString(user)
	now := time.Now().Unix()
	claims[JWT_KEY_ISS] = JWT_ISSUER
	claims[JWT_KEY_EXP] = now + int64(expireIn)
	claims[JWT_KEY_IAT] = now
	claims[USER_DATA] = userStr
	//生成Token sign
	tokenStr, err := token.SignedString([]byte(JWT_TOKEN_SECRET_KEY))
	if err != nil {
		return nil, err
	}
	//如果为传出参数 加上参数名
	if access {
		tokenStr = JWT_PREFIX + tokenStr
	}
	//填充数据
	t = &JWTToken{}
	t.Token = tokenStr
	t.Expire = int64(expireIn)
	t.UserData = user
	return t, nil
}

// 获取jwt key的函数(用于传入参数)
func keyFunc(t *jwt.Token) (interface{}, error) {
	return []byte(JWT_TOKEN_SECRET_KEY), nil
}

// 通过request获取 jwt.Token
func ParseFromRequest(req *http.Request) (*jwt.Token, error) {
	return request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, keyFunc)
}

// 将 token string 解析为 jwt.Token
func ParseFromToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, keyFunc)
}

/*
将token string解析为JWTToken对象
tokenStr 收到的token string
*/
func Decode(tokenStr string) (*JWTToken, error) {
	//将 token string 解析为 jwt.Token
	token, err := ParseFromToken(tokenStr)
	if err != nil {
		return nil, err
	}
	//填充JWTToken
	t := &JWTToken{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		switch key {
		case USER_DATA:
			{
				userStr := value.(string)
				t.UserData, _ = base64.StdEncoding.DecodeString(userStr)
			}
		case JWT_KEY_EXP:
			{
				t.Expire = int64(value.(float64))
			}
		default:
		}
	}
	return t, nil
}
