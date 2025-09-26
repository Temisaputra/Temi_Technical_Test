package router

import (
	"net/http"

	"github.com/Temisaputra/warOnk/delivery/handler"
	"github.com/Temisaputra/warOnk/delivery/middleware"
	_ "github.com/Temisaputra/warOnk/docs" // wajib untuk register doc
	"github.com/Temisaputra/warOnk/pkg/auth"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type Handlers struct {
	VoucherHandler *handler.VoucherHandler
	UserHandler    *handler.UserHandler
	AuthHandler    *handler.AuthHandler
	Logger         *zap.Logger
	JwtService     auth.JwtService
}

// NewRouter bikin router dan register semua endpoint
func NewRouter(handlers *Handlers) http.Handler {
	router := mux.NewRouter()
	authMW := middleware.NewAuthMiddleware(handlers.JwtService)

	router.Use(middleware.LoggingMiddleware(handlers.Logger)) // <- inject logger

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := router.PathPrefix("/api/").Subrouter()

	// ---------------- Public ----------------
	// Auth endpoints
	api.HandleFunc("/register", handlers.AuthHandler.Register).Methods("POST")
	api.HandleFunc("/login", handlers.AuthHandler.Login).Methods("POST")

	// ---------------- Protected ----------------
	protected := api.PathPrefix("").Subrouter()
	protected.Use(authMW.Authorization)

	// User endpoints
	protected.HandleFunc("/users", handlers.UserHandler.GetAllUser).Methods("GET")
	protected.HandleFunc("/user/{id}", handlers.UserHandler.GetUserByID).Methods("GET")
	protected.HandleFunc("/user-create", handlers.UserHandler.CreateUser).Methods("POST")
	protected.HandleFunc("/user-update/{id}", handlers.UserHandler.UpdateUser).Methods("PUT")
	protected.HandleFunc("/user-delete/{id}", handlers.UserHandler.DeleteUser).Methods("DELETE")

	// Voucher endpoints
	protected.HandleFunc("/vouchers", handlers.VoucherHandler.GetListVoucher).Methods("GET")
	protected.HandleFunc("/voucher/{id}", handlers.VoucherHandler.GetVoucherByID).Methods("GET")
	protected.HandleFunc("/voucher-create", handlers.VoucherHandler.CreateVoucher).Methods("POST")
	protected.HandleFunc("/voucher-update/{id}", handlers.VoucherHandler.UpdateVoucher).Methods("PUT")
	protected.HandleFunc("/voucher-delete/{id}", handlers.VoucherHandler.DeleteVoucher).Methods("DELETE")
	protected.HandleFunc("/vouchers/upload-csv", handlers.VoucherHandler.UploadVouchersFromCSV).Methods("POST")
	protected.HandleFunc("/vouchers/export-csv", handlers.VoucherHandler.ExportVouchersToCSV).Methods("GET")

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Mode"},
		MaxAge:             60,
		AllowCredentials:   true,
		OptionsPassthrough: false,
	})

	return c.Handler(router)
}
