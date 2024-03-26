package postgres

import (
	"fmt"
	"time"

	"github.com/gookit/slog"
)

type Options func(s *PostgresParams) error

type PostgresParams struct {
	port            string
	host            string
	password        string
	dbName          string
	user            string
	sslmode         string
	connMaxIdleTime time.Duration
	connMaxLifetime time.Duration
	maxIdleConns    int32
	slog            *slog.SugaredLogger
}

func newPostgresParams(opts ...Options) (*PostgresParams, error) {
	s := &PostgresParams{}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func WithPort(port string) Options {
	return func(s *PostgresParams) error {
		s.port = port
		return nil
	}
}

func WithHost(host string) Options {
	return func(s *PostgresParams) error {
		s.host = host
		return nil
	}
}

func WithPassword(password string) Options {
	return func(s *PostgresParams) error {
		s.password = password
		return nil
	}
}

func WithUser(user string) Options {
	return func(s *PostgresParams) error {
		s.user = user
		return nil
	}
}

func WithSSLMode(sslmode string) Options {
	return func(s *PostgresParams) error {
		s.sslmode = sslmode
		return nil
	}
}

func WithDatabaseName(dbName string) Options {
	return func(s *PostgresParams) error {
		s.dbName = dbName
		return nil
	}
}

func WithConnMaxIdleTime(connMaxIdleTime string) Options {
	return func(s *PostgresParams) error {
		t, err := time.ParseDuration(connMaxIdleTime)
		if err != nil {
			return err
		}
		s.connMaxIdleTime = t
		return nil
	}
}

func WithConnMaxLifetime(connMaxLifetime string) Options {
	return func(s *PostgresParams) error {
		t, err := time.ParseDuration(connMaxLifetime)
		if err != nil {
			return err
		}
		s.connMaxLifetime = t
		return nil
	}
}

func WithMaxIdleConns(MaxIdleConns int32) Options {
	return func(s *PostgresParams) error {
		s.maxIdleConns = MaxIdleConns
		return nil
	}
}

func WithSlog(l *slog.SugaredLogger) Options {
	return func(s *PostgresParams) error {
		s.slog = l
		return nil
	}
}

// getters -----

func (s *PostgresParams) Port() string {
	return s.port
}

func (s *PostgresParams) Host() string {
	return s.host
}

func (s *PostgresParams) Password() string {
	return s.password
}

func (s *PostgresParams) User() string {
	return s.user
}

func (s *PostgresParams) SSLMode() string {
	return s.sslmode
}

func (s *PostgresParams) DatabaseName() string {
	return s.dbName
}

func (s *PostgresParams) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s dbname=%s", s.host, s.port, s.user, s.password, s.sslmode, s.dbName)
}

func (s *PostgresParams) WithParams() map[string]interface{} {
	return map[string]interface{}{
		"host":     s.host,
		"port":     s.port,
		"user":     s.user,
		"password": s.password,
		"sslmode":  s.sslmode,
		"dbname":   s.dbName,
	}
}

func (s *PostgresParams) WithDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", s.user, s.password, s.host, s.port, s.dbName, s.sslmode)
}

func (s *PostgresParams) MaxIdleConns() int32 {
	return s.maxIdleConns
}

func (s *PostgresParams) ConnMaxIdleTime() time.Duration {
	return s.connMaxIdleTime
}

func (s *PostgresParams) ConnMaxLifetime() time.Duration {
	return s.connMaxLifetime
}

func (s *PostgresParams) GetSlog() *slog.SugaredLogger {
	return s.slog
}

// setters -----

func (s *PostgresParams) SetPort(port string) {
	s.port = port
}

func (s *PostgresParams) SetHost(host string) {
	s.host = host
}

func (s *PostgresParams) SetPassword(password string) {
	s.password = password
}

func (s *PostgresParams) SetUser(user string) {
	s.user = user
}

func (s *PostgresParams) SetSSLMode(sslmode string) {
	s.sslmode = sslmode
}

func (s *PostgresParams) SetDatabaseName(dbName string) {
	s.dbName = dbName
}

func (s *PostgresParams) SetMaxIdleConns(maxIdleConns int32) {
	s.maxIdleConns = maxIdleConns
}

func (s *PostgresParams) SetConnMaxIdleTime(connMaxIdleTime time.Duration) {
	s.connMaxIdleTime = connMaxIdleTime
}

func (s *PostgresParams) SetConnMaxLifetime(connMaxLifetime time.Duration) {
	s.connMaxLifetime = connMaxLifetime
}

func (s *PostgresParams) SetSlog(l *slog.SugaredLogger) {
	s.slog = l
}
