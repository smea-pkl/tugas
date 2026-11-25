package user_service

import (
	"context"
	"fmt"
	"net/http"
	"pkl/backend/gen/pkl/users/v1/usersv1connect"
	"pkl/backend/pkgs/san_config"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type CheckLoginInterceptor struct{}

// WrapStreamingClient implements [connect.Interceptor].
func (c *CheckLoginInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

// WrapStreamingHandler implements [connect.Interceptor].
func (c *CheckLoginInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

// WrapUnary implements [connect.Interceptor].
func (c *CheckLoginInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		token :=(req.Header().Get("Authorization"))

		// if strings.TrimSpace(token) == "" {
		// 	return nil, fmt.Errorf("Token kosong")
		// }

		// token = strings.Replace(token, "Bearer ", "", -1)

		cfg, err := san_config.RenderYamlConfig()
		if err != nil {
			return nil, err
		}

		tkn, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (any, error) { return []byte(cfg.JwtKey), nil }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			return nil, err
		}

		if _, ok := tkn.Claims.(*CustomClaims); ok && tkn.Valid {
			return next(ctx, req)
		}

		fmt.Println(token)
		return nil, fmt.Errorf("Invalid token")
	}
}

func UserServiceRegister(mux *http.ServeMux, db *gorm.DB) error {

	interceptor := connect.WithInterceptors(
		validate.NewInterceptor(),
		&CheckLoginInterceptor{},
	)

	path, handler := usersv1connect.NewUserServiceHandler(
		&userServiceImpl{db: db},
		interceptor,
	)

	mux.Handle(path, handler)

	return nil
}
