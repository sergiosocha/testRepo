package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"
)

func Apply(dbConn *sql.DB) error {
	return ApplyFile(dbConn, "scripts/create_table_users.sql")
}

func ApplyFile(dbConn *sql.DB, path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("leyendo %s: %w", path, err)
	}

	stmts := strings.Split(string(b), ";")
	for i, st := range stmts {
		s := strings.TrimSpace(st)
		if s == "" {
			continue
		}
		if _, err := dbConn.ExecContext(ctx, s); err != nil {
			return fmt.Errorf("stmt #%d en %s: %w\nSQL: %s", i+1, path, err, s)
		}
	}
	return nil
}
