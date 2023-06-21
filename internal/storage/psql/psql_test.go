package psql

import (
	"context"
	"testing"
	//"github.com/jackc/pgx/v5/pgxpool"
)

func TestConnection(t *testing.T) {
	d := InitConnectionInfo()
	dpPool := Connect(d)
	defer dpPool.Db.Close()
	if err := dpPool.Db.Ping(context.Background()); err != nil {
		t.Errorf("Error pinging db %v", err)
	}
}
