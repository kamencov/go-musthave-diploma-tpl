package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/handlers/authorize"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/handlers/balance"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/handlers/order"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/handlers/register"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/handlers/withdraw"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/logger"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/middleware"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/service/auth"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/service/orders"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/storage/db"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/workers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// инициализируем Config
	cfg := NewConfig()
	cfg.Parsed()

	// инициализируем Logger
	logs := logger.NewLogger(logger.WithLevel(cfg.LogLevel))
	logs.Info("Logger start")

	// инициализируем DB
	repo, err := db.NewDB(logs, cfg.AddrConDB)
	if err != nil {
		logs.Error("Fatal = not connect DB", "customerrors = ", err)
		panic(err)
	}
	logs.Info("DB connection")
	defer repo.Close()

	// инициализируем Service
	serv := orders.NewService(repo, logs)
	logs.Info("Service run")

	// инициализируем проверку авторизацию
	serviceAuth := auth.NewService(cfg.TokenSalt, cfg.PasswordSalt, repo)
	logs.Info("Service authorize run")

	// запускаем воркер
	worker := workers.NewWorkerAccrual(serv, logs)

	//инициализируем middleware
	authorization := middleware.NewAuthMiddleware(serviceAuth)

	// инициализируем Handlers
	registerHandler := register.NewHandlers(serviceAuth, logs)
	authorizeHandler := authorize.NewHandler(serviceAuth, logs)
	ordersHandler := order.NewHandler(serv, logs)
	balanceHandler := balance.NewHandler(serv, logs)
	withdrawHandler := withdraw.NewHandler(serv, logs)

	// инициализировали роутер и создали запросы
	r := chi.NewRouter()
	r.Use(middleware.WithLogging)

	r.Post("/api/user/register", registerHandler.Post)
	r.Post("/api/user/login", authorizeHandler.Post)

	r.Route("/api/user", func(r chi.Router) {
		r.Use(authorization.ValidAuth)
		r.Post("/orders", ordersHandler.Post)
		r.Get("/orders", ordersHandler.Get)
		r.Get("/balance", balanceHandler.Get)
		r.Post("/balance/withdraw", withdrawHandler.Post)
		r.Get("/withdrawals", withdrawHandler.Get)
	})

	// инициализируем запись Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker.StartWorkerAccrual(ctx, cfg.AccrualSystemAddress)

	// Слушаем сигналы
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			logs.Error("Err:", logger.ErrAttr(err))
		}

	}()

	<-sigs
	cancel()

	// Ждем завершения всех работников
	time.Sleep(time.Second * 2)

	logs.Info("Сервер завершил работу грациозно")

}
