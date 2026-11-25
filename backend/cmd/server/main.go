package main

import (
	"log/slog"
	"net/http"
	"pkl/backend/database"
	"pkl/backend/pkgs/san_config"
	"pkl/backend/services/user_service"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
)

func main() {
	cfg, err := san_config.RenderYamlConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	user_service.UserServiceRegister(mux, db)

	slog.Info("running server", "Listen", cfg.ListenAddress)

	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)

	appServer := &http.Server{
		Addr:      cfg.ListenAddress,
		Handler:   withCORS(mux),
		Protocols: protocols,
	}

	err = appServer.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func withCORS(handler http.Handler) http.Handler {
	// connectcors.AllowedHeaders() covers only the CONNECT PROTOCOL headers
	// (Content-Type, Connect-Protocol-Version, timeouts, …). It does NOT include Authorization
	// — that is ours to add.
	//
	// Omitting it fails in a nasty, asymmetric way: Login carries no Authorization header, so it
	// sails through, and EVERY authenticated call afterwards is killed by the browser at
	// preflight — before the request ever reaches the server. Nothing appears in the server log.
	// From the UI it looks like the API returned nothing.
	allowedHeaders := append(connectcors.AllowedHeaders(), "Authorization")

	// PUT is not a Connect method; it is added for the local document file endpoint, which clients
	// upload to directly with a cross-origin PUT.
	allowedMethods := append(connectcors.AllowedMethods(), http.MethodPut)

	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: allowedMethods,
		AllowedHeaders: allowedHeaders,
		ExposedHeaders: connectcors.ExposedHeaders(),
	})

	app := middleware.Handler(handler)

	// An AGENT's requests are not our app's — ChatGPT, Claude — so they bypass the app's origin list and answer their own
	// CORS. The bearer, or PKCE, is what authorizes them, never an origin.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		app.ServeHTTP(w, r)
	})
}