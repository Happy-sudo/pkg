package jwt

import (
	"context"
	"fmt"
	krserror "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
	"time"
)

type JwObj struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	StaffID  string `json:"staff_id"`
}

var (
	secretKey []byte
	expire    int
	expirePc  int
)

// CustomClaims 自定义payload结构体,不建议直接使用 dgrijalva/jwt-go `jwt.StandardClaims`结构体.因为他的payload包含的用户信息太少.
type CustomClaims struct {
	jwt.StandardClaims
	JwObj
}

// GenerateToken 生成token params
func GenerateToken(obj JwObj, equipment string) (int64, int64, string, error) {

	var ExpiresAt int64

	if len(equipment) <= 0 {
		return 0, 0, "", errors.New("参数不能为空")
	}
	if strings.ToLower(equipment) == "app" {
		ExpiresAt = time.Now().Unix() + int64(expire) //过期时间
	}
	if strings.ToLower(equipment) == "pc" {
		ExpiresAt = time.Now().Unix() + int64(expirePc) //过期时间
	}
	if len(equipment) >= 0 && strings.ToLower(equipment) != "app" && strings.ToLower(equipment) != "pc" {
		return 0, 0, "", errors.New("请求参数错误")
	}

	switch equipment {
	case "app":
		ExpiresAt = time.Now().Unix() + int64(expire) //过期时间
		break
	case "pc":
		ExpiresAt = time.Now().Unix() + int64(expirePc) //过期时间
		break
	default:
		return 0, 0, "", errors.New("请求参数错误")
	}

	IssuedAt := time.Now().Unix() //签发时间

	stdClaims := jwt.StandardClaims{
		ExpiresAt: ExpiresAt,
		IssuedAt:  IssuedAt,
		Id:        fmt.Sprintf("%s", obj.UserId),
		Issuer:    "clouderpSrv",
	}

	uClaims := CustomClaims{
		StandardClaims: stdClaims,
		JwObj:          obj,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return 0, 0, "", errors.Wrap(err, "生成Token失败")
	}
	return IssuedAt, ExpiresAt, tokenString, err
}

// ParseToken 解析token
func ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

//AuthMiddleware 根据header读取jwt token
func AuthMiddleware(c *conf.ApolloList, logger log.Logger) middleware.Middleware {
	secretKey = []byte(c.JwtSecret.GetAccessSecret())
	expire, _ = strconv.Atoi(c.JwtSecret.GetAccessExpire())
	expirePc, _ = strconv.Atoi(c.JwtSecret.GetAccessExpirePc())
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var token string
			tr, ok := transport.FromServerContext(ctx)
			if ok {
				authorization := tr.RequestHeader().Get("Authorization")
				if authorization != "" {
					token = authorization
				}
			}
			md, ok := metadata.FromIncomingContext(ctx)
			if ok && len(md.Get("Authorization")) > 0 {
				token = md.Get("Authorization")[0]
			}
			if token == "" {
				//logger.Log(log.LevelError, err, errs.SystemForbiddenErrorMsg, errs.SystemParameterCannotEmpty)
				return nil, krserror.BadRequest("暂无权限访问", "参数不能为空")
			}

			ctx = context.WithValue(ctx, "Authorization", token)

			// 解析token
			claims, err := ParseToken(token)
			if err != nil {
				//logger.Log(log.LevelError, err, errs.SystemForbiddenErrorMsg, errs.SystemInvalidToken)
				return nil, krserror.Unauthorized("暂无权限访问", "无效的Token")
			} else if time.Now().Unix() > claims.ExpiresAt {
				//logger.Log(log.LevelError, err, errs.SystemForbiddenErrorMsg, errs.SystemTokenHasExpired)
				return nil, krserror.Unauthorized("暂无权限访问", "Token已过期")
			}
			ctx = context.WithValue(ctx, "UserId", claims.UserId)
			ctx = context.WithValue(ctx, "UserName", claims.UserName)
			ctx = context.WithValue(ctx, "StaffID", claims.StaffID)
			return handler(ctx, req)
		}
	}
}
