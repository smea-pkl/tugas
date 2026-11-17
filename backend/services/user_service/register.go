package user_service

import (
	"net/http"
	"pkl/backend/configs"
	"pkl/backend/gen/pkl/users/v1/usersv1connect"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"gorm.io/gorm"
)

func UserServiceRegister(mux *http.ServeMux, db *gorm.DB, cfg *configs.AppConfig) error {
	path, handler := usersv1connect.NewUserServiceHandler(
		&userServiceImpl{db: db, jwtKey: []byte(cfg.JwtKey)},
		connect.WithInterceptors(
			validate.NewInterceptor(),
		),
	)
	mux.Handle(path, handler)

	return nil
}
