package psql

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"nats-streaming-service/internal/config"
	"nats-streaming-service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgresql struct {
	ConnectionString config.PostgresConfig
	Db               *pgxpool.Pool
}

func (d *Postgresql) connect() error {
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", d.ConnectionString.User, d.ConnectionString.Password, d.ConnectionString.Host, d.ConnectionString.Port, d.ConnectionString.Name)
	cfg, err := pgxpool.ParseConfig(connectionStr)
	if err != nil {
		return err
	}
	dpPool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return err
	}
	d.Db = dpPool
	return nil
}

func Connect(l *log.Logger) (*Postgresql, error) {
	d := config.InitPsqlConfig(l)
	var newPsqlconnection Postgresql
	newPsqlconnection.ConnectionString = d
	err := newPsqlconnection.connect()
	if err != nil {
		return nil, err
	}
	err = newPsqlconnection.EnsureTableExists()
	if err != nil {
		return nil, err
	}
	return &newPsqlconnection, nil
}

func (d *Postgresql) EnsureTableExists() error {
	_, err := d.Db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS orders(
		order_uid varchar PRIMARY KEY,
		order_data jsonb NOT NULL
	)`)
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

	var uid string
	var searchOrder model.Order
	for rows.Next() {
		rows.Scan(&uid, &searchOrder)
		orders[uid] = &searchOrder
	}
	return orders, nil
}
