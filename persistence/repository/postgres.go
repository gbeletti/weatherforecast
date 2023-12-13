package repository

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const retryWait = 3 * time.Second
const maxRetryAttempts = 3

type SQLConfig struct {
	HOST           string
	PORT           string
	DATABASE       string
	USER           string
	PASSWORD       string
	REQUIRESSL     string
	MAX_CONNS      int
	MAX_IDLE_CONNS int
}

func GetPostgresRepo() (*ForecastRepo, error) {
	dbClient, err := GetDb()
	return NewForecastRepo(dbClient), err
}

func GetDb() (db *sqlx.DB, err error) {
	conCount, err := strconv.Atoi(os.Getenv("MAX_CONNS"))
	if err != nil {
		return nil, err
	}
	conIdleCount, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	if err != nil {
		return nil, err
	}
	connectionInfo := SQLConfig{
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DATABASE"),
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBSSL"),
		conCount,
		conIdleCount,
	}
	dbClient, err := connectClientWithRetry(connectionInfo, maxRetryAttempts)
	if err != nil {
		return nil, err
	}
	return dbClient, nil
}

func (i SQLConfig) getConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		i.HOST,
		i.PORT,
		i.USER,
		i.PASSWORD,
		i.DATABASE,
		i.REQUIRESSL,
	)
}

func newPgClient(info SQLConfig) (*sqlx.DB, error) {
	slog.Info(fmt.Sprintf("Connecting to %s/%s...",
		info.HOST,
		info.DATABASE),
	)

	db, err := sqlx.Open("postgres", info.getConnectionString())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(info.MAX_CONNS)
	db.SetMaxIdleConns(info.MAX_IDLE_CONNS)

	err = db.Ping()
	if err != nil {
		slog.Info(
			fmt.Sprintf("Unable to connect to %s/%s...",
				info.HOST,
				info.DATABASE),
		)
		return nil, err
	}

	slog.Info(
		fmt.Sprintf("Connected to %s/%s...",
			info.HOST,
			info.DATABASE),
	)
	return db, nil
}

func connectClientWithRetry(info SQLConfig, retries int) (db *sqlx.DB, err error) {
	for retries > 0 {
		db, err = newPgClient(info)
		if err != nil {
			slog.Error(fmt.Sprintf("Error getting new pg client %s", err))
			retries--
			time.Sleep(retryWait)
			continue
		}

		break
	}
	return db, err
}
