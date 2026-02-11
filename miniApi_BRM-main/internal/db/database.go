package db

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string

	TLS   string
	SSLCA string
}

func NewConnection(cfg Config) (*sql.DB, error) {

	tlsParam := "false"

	switch cfg.TLS {
	case "verify":

		caPath := cfg.SSLCA
		if caPath == "" {
			caPath = "internal/db/certs/aiven-ca.pem"
		}
		abs, _ := filepath.Abs(caPath)
		pem, err := os.ReadFile(abs)
		if err != nil {
			return nil, fmt.Errorf("leyendo CA (%s): %w", abs, err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("no se pudo cargar CA desde %s", abs)
		}

		if err := mysql.RegisterTLSConfig("custom", &tls.Config{
			MinVersion: tls.VersionTLS12,
			RootCAs:    pool,
			ServerName: cfg.Host,
		}); err != nil {
			return nil, fmt.Errorf("registrando TLS custom: %w", err)
		}
		tlsParam = "custom"

	case "true":

		tlsParam = "true"

	case "skip-verify":

		_ = mysql.RegisterTLSConfig("skip", &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true,
			ServerName:         cfg.Host,
		})
		tlsParam = "skip"

	case "false", "":
		tlsParam = "false"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=%s&parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		tlsParam,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	return db, nil
}
