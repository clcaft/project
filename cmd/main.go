package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"procurement-system/internal/config"
	"procurement-system/internal/database"
	"procurement-system/internal/handler"
	"procurement-system/internal/logger"
	"procurement-system/internal/repository"
	"procurement-system/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	cfg := config.Load()
	log := logger.New(cfg.LogLevel)
	log.Info("Starting Procurement System API")

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database", "error", err)
	}
	defer db.Close()
	log.Info("Database connected")

	repos := repository.NewRepositories(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandlers(services, log)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(logger.ChiMiddleware(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/departments", handlers.Department.Routes)
		r.Route("/warehouses", handlers.Warehouse.Routes)
		r.Route("/stores", handlers.Store.Routes)
		r.Route("/suppliers", handlers.Supplier.Routes)
		r.Route("/products", handlers.Product.Routes)
		r.Route("/product-categories", handlers.ProductCategory.Routes)
		r.Route("/product-units", handlers.ProductUnit.Routes)
		r.Route("/purchase-requests", handlers.PurchaseRequest.Routes)
		r.Route("/delivery-invoices", handlers.DeliveryInvoice.Routes)
		r.Route("/internal-transfer-invoices", handlers.InternalTransferInvoice.Routes)
		r.Route("/inventory-balances", handlers.InventoryBalance.Routes)
		r.Route("/reports", handlers.Report.Routes)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.APIPort),
		Handler: r,
	}

	go func() {
		log.Info("Server starting", "port", cfg.APIPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", "error", err)
	}
	log.Info("Server exited")
}
