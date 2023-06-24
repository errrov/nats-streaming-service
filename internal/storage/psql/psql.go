package psql

import (
	//"wildberries_L0/internal/model"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"fmt"
	"context"
)

type ConnectionInfo struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type Postgresql struct {
	ConnectionString ConnectionInfo
	Db               *pgxpool.Pool
}

func InitConnectionInfo() ConnectionInfo {
	User := os.Getenv("POSTGRES_USER")
	Password := os.Getenv("POSTGRES_PASSWORD")
	Host := "postgres"
	Port := "5432"
	Name := os.Getenv("POSTGRES_DB")
	return ConnectionInfo{User: User, Password: Password, Host: Host, Port: Port, Name: Name}
}

func (d *Postgresql) connect() {
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", d.ConnectionString.User, d.ConnectionString.Password, d.ConnectionString.Host, d.ConnectionString.Port, d.ConnectionString.Name)
	cfg, err := pgxpool.ParseConfig(connectionStr)
	if err != nil {
		log.Panic(err)
	}
	dpPool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Panic(err)
	}
	d.Db = dpPool
}

func Connect(d ConnectionInfo) *Postgresql {
	var newPsqlconnection Postgresql
	newPsqlconnection.ConnectionString = d
	newPsqlconnection.connect()
	log.Println("Connected")
	return &newPsqlconnection
}