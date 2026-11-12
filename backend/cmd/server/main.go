package main

import (
	"log/slog"
	"net/http"
	"pkl/backend/gen/pkl/users/v1/usersv1connect"
)

func main() {
	mux := http.NewServeMux()
	path, handler := usersv1connect.NewUserServiceHandler(&userServiceImpl{})

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := "Hello Rizki"
		w.Write([]byte(response))
	})
	mux.Handle(path, handler)

	slog.Info("Server berjalan di localhost:8080")
	// err := http.ListenAndServe("localhost:8080",
	// 	h2c.NewHandler(mux, &http2.Server{}),
	// )
	// if err != nil {
	// 	panic(err)
	// }
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetHTTP2(true)
	protocols.SetUnencryptedHTTP2(true)

	appServer := &http.Server{
		Addr:      "localhost:8080",
		Handler:   mux,
		Protocols: protocols,
	}

	err := appServer.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
