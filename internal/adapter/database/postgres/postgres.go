//go:generate ifacemaker -f $GOFILE -s Postgres -i PostgresRepo -p postgres -o postgres_repo.go
//go:generate mockgen -source=postgres_repo.go -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	params *PostgresParams
	db     *pgxpool.Pool
}

func Connection(opts ...Options) (*Postgres, error) {
	params, err := newPostgresParams(opts...)
	if err != nil {
		return nil, err
	}

	// configure the pgx with pool
	conf, err := pgxpool.ParseConfig(params.String())
	if err != nil {
		return nil, err
	}

	conf.MaxConnIdleTime = params.ConnMaxIdleTime()
	conf.MaxConnLifetime = params.ConnMaxLifetime()
	conf.MaxConns = params.MaxIdleConns()

	ctx := context.Background()
	conn, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if the connection is alive
	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &Postgres{
		params: params,
		db:     conn,
	}, nil
}

func (d *Postgres) Close() {
	d.db.Close()
}

func (d *Postgres) Ping(ctx context.Context) error {
	return d.db.Ping(ctx)
}

func (d *Postgres) String() string {
	return d.params.String()
}

func (d *Postgres) DB() *pgxpool.Pool {
	return d.db
}

func (d *Postgres) Insert(ctx context.Context, query string,
	args ...any) (pgconn.CommandTag, error) {
	return d.db.Exec(ctx, query, args...)
}

func (d *Postgres) Query(ctx context.Context, query string,
	args ...any) (pgx.Rows, error) {
	return d.db.Query(ctx, query, args...)
}

func (d *Postgres) QueryRow(ctx context.Context, query string,
	args ...any) pgx.Row {
	return d.db.QueryRow(ctx, query, args...)
}

func (d *Postgres) Begin(ctx context.Context) (pgx.Tx, error) {
	return d.db.Begin(ctx)
}
