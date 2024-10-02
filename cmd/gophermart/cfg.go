package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Configs struct {
	RunAddress           string
	LogLevel             string
	AddrConDB            string
	AccrualSystemAddress string
	TokenSalt            []byte
	PasswordSalt         []byte
}

func NewConfig() *Configs {
	return &Configs{}
}

func (c *Configs) Parsed() {
	c.initSaltFromEnv()
	c.parseFlags()
	// Проверка переменной окружения RUN_ADDRESS
	if c.RunAddress == "" {
		if envRunAddress := os.Getenv("RUN_ADDRESS"); envRunAddress != "" {
			c.RunAddress = envRunAddress
		}
	}

	// Проверка переменной окружения LOG_LEVEL
	if c.LogLevel == "" {
		if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
			c.LogLevel = envLogLevel
		}
	}

	// Проверка переменной окружения DATABASE_URI

	if c.AddrConDB == "" {
		if envAddrConDB := os.Getenv("DATABASE_URI"); envAddrConDB != "" {
			c.AddrConDB = envAddrConDB
		}
	}

	// Проверка переменной окружения ACCRUAL_SYSTEM_ADDRESS
	if c.AccrualSystemAddress == "" {
		if envAccrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccrualSystemAddress != "" {
			c.AccrualSystemAddress = envAccrualSystemAddress
		}
	}

}

func (c *Configs) parseFlags() {
	// Флаг -a отвечает за адрес запуска HTTP-сервера (значение может быть таким: localhost:8080).
	flag.StringVar(&c.RunAddress, "a", ":8080", "Server address host:port")
	// Флаг -l отвечает за logger
	flag.StringVar(&c.LogLevel, "l", "info", "log level")
	// Флаг -p отвечает за адрес подключения DB
	flag.StringVar(&c.AddrConDB, "d", "", "address DB")
	// Флаг -r отвечает за адрес системы расчета начислений
	flag.StringVar(&c.AccrualSystemAddress, "r", "", "address accrual system address")
	flag.Parse()
}

func (c *Configs) initSaltFromEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		slog.Error("Fatal", "error loading .env file = ", err)
		return
	}

	c.TokenSalt = []byte(os.Getenv("TOKEN_SALT"))
	c.PasswordSalt = []byte(os.Getenv("PASSWORD_SALT"))
}
