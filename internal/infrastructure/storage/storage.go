package storage

import (
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	DB *sql.DB
}

func InitDatabase(logger *logrus.Logger, connectionString string) (*Storage, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger.Debug("Bad connection string")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Debug("Connection Failed")
		logger.Info("Connection Failed")
		return nil, err
	}
	logger.Debug(connectionString)

	logger.Info("Database connected successfully")
	storage := &Storage{DB: db}
	return storage, nil
}

func RunMigrations(logger *logrus.Logger, connectionString string) error {
	migrationsPath := os.Getenv("MIGRATION_PATH")

	m, err := migrate.New(
		migrationsPath,
		connectionString,
	)
	if err != nil {
		logger.Debug("Migration error")
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		logger.Debug("Migration error")
		return err
	}

	logger.Debug("Migrations applied successfully!")
	return nil
}
