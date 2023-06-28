package psql

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"wildberries_L0/internal/model"
)

var jStr = `
{
	"order_uid": "b563feb7b2b84b6test",
	"track_number": "WBILMTESTTRACK",
	"entry": "WBIL",
	"delivery": {
	  "name": "Test Testov",
	  "phone": "+9720000000",
	  "zip": "2639809",
	  "city": "Kiryat Mozkin",
	  "address": "Ploshad Mira 15",
	  "region": "Kraiot",
	  "email": "test@gmail.com"
	},
	"payment": {
	  "transaction": "b563feb7b2b84b6test",
	  "request_id": "",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 1817,
	  "payment_dt": 1637907727,
	  "bank": "alpha",
	  "delivery_cost": 1500,
	  "goods_total": 317,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 317,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	  }
	],
	"locale": "en",
	"internal_signature": "",
	"customer_id": "test",
	"delivery_service": "meest",
	"shardkey": "9",
	"sm_id": 99,
	"date_created": "2021-11-26T06:22:19Z",
	"oof_shard": "1"
}
`

func TestConnection(t *testing.T) {
	dpPool := Connect()
	defer dpPool.Db.Close()
	if err := dpPool.Db.Ping(context.Background()); err != nil {
		t.Errorf("Error pinging db %v", err)
	}
}

func TestInsertOrder(t *testing.T) {
	dpPool := Connect()
	defer dpPool.Db.Close()
	testOrder := &model.Order{}
	err := json.Unmarshal([]byte(jStr), testOrder)
	if err != nil {
		t.Errorf("Error unmarshalling json (%v)", err)
	}
	if err = dpPool.InsertOrder(testOrder); err != nil {
		t.Errorf("Error : %v", err)
	}
}

func TestFindByUID(t *testing.T) {
	dpPool := Connect()
	defer dpPool.Db.Close()
	testOrder, err := dpPool.FindByUID("b563feb7b2b84b6test")
	if err != nil {
		t.Errorf("Error while searching : %v", err)
	}
	log.Println(*testOrder)
}

func TestFindAll(t *testing.T) {
	dpPool := Connect()
	defer dpPool.Db.Close()
	m1, err := dpPool.FindAll()
	if err != nil {
		t.Errorf("error scanning rows %v", err)
	}
	log.Println(m1)
}
