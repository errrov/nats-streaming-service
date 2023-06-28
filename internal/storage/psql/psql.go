package psql

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	//"os"
	"wildberries_L0/internal/config"
	"wildberries_L0/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionInfo struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type Postgresql struct {
	ConnectionString config.PostgresConfig
	Db               *pgxpool.Pool
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

func Connect() *Postgresql {
	d := config.InitPsqlConfig()
	log.Println(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", d.User, d.Password, d.Host, d.Port, d.Name))
	var newPsqlconnection Postgresql
	newPsqlconnection.ConnectionString = d
	newPsqlconnection.connect()
	err := newPsqlconnection.EnsureTableExists()
	if err != nil {
		return nil
	}
	log.Println("Created connection")
	return &newPsqlconnection
}

func (d *Postgresql) EnsureTableExists() error {
	_, err := d.Db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS orders(
		order_uid varchar PRIMARY KEY,
		order_data jsonb NOT NULL
	)`)
	log.Println("Creating table", err)
	return err
}

func (d *Postgresql) InsertOrder(order *model.Order) error {
	orderQueue := `INSERT INTO ORDERS VALUES ($1, $2)`
	json, err := json.Marshal(*order)
	if err != nil {
		return err
	}
	_, err = d.Db.Exec(context.Background(), orderQueue, order.OrderUID, json)
	if err != nil {
		return err
	}
	return nil
}

func (d *Postgresql) FindByUID(uid string) (*model.Order, error) {
	sql := `SELECT order_data::jsonb FROM orders WHERE order_uid=$1`
	var searchingOrder model.Order
	err := d.Db.QueryRow(context.Background(), sql, uid).Scan(&searchingOrder)
	if err != nil {
		return nil, err
	}
	return &searchingOrder, nil
}

func (d *Postgresql) FindAll() (map[string]*model.Order, error) {
	orders := make(map[string]*model.Order)
	rows, err := d.Db.Query(context.Background(), "select * from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var uid string
		var searchOrder model.Order
		rows.Scan(&uid, &searchOrder)
		orders[uid] = &searchOrder
	}
	return orders, nil
}
