package model

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
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

func TestJSONParsing(t *testing.T) {
	testOrder := &Order{}
	err := json.Unmarshal([]byte(jStr), testOrder)
	if err != nil {
		t.Errorf("Error unmarshalling json (%v)", err)
	}
}

func TestValidation(t *testing.T) {
	var testOrder Order
	err := json.Unmarshal([]byte(jStr), &testOrder)
	if err != nil {
		t.Errorf("Error unmarshalling json (%v)", err)
	}
	validate := validator.New()
	err = validate.Struct(testOrder)
	if err != nil {
		t.Errorf("Error validating struct (%v)", err)
	}

	testOrder.Payment.Amount = 0
	err = validate.Struct(testOrder)
	if err == nil {
		t.Errorf("Struct was not validated field (%v) with value(%v)", "Payment.Amount", testOrder.Payment.Amount)
	}

	testOrder.Payment.Amount = 1
	testOrder.Payment.DeliveryCost = 0
	err = validate.Struct(testOrder)
	if err == nil {
		t.Errorf("Struct was not validated field (%v) with value(%v)", "Payment.Delivery_cost", testOrder.Payment.DeliveryCost)
	}

	testOrder.Payment.DeliveryCost = 1
	testOrder.Items[0].Price = 0
	err = validate.Struct(testOrder.Items[0])
	if err == nil {
		t.Errorf("Struct was not validated field (%v) with value(%v)", "Items.Price", testOrder.Items[0].Price)
	}

	testOrder.Items[0].Price = 1
	testOrder.Items[0].TotalPrice = 0
	err = validate.Struct(testOrder.Items[0])
	if err == nil {
		t.Errorf("Struct was not validated field (%v) with value(%v)", "Items.TotalCost", testOrder.Items[0].TotalPrice)
	}
}
