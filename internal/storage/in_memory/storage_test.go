package inMemoryStorage

import (
	"encoding/json"
	"testing"
	"wildberries_L0/internal/model"
)

var testJSON = `
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
	  },
	  {
		"chrt_id": 9934932,
		"track_number": "WBILMTESTTRACK2",
		"price": 460,
		"rid": "ab4219087a764ae0btest2",
		"name": "Mascaras2",
		"sale": 32,
		"size": "1",
		"total_price": 330,
		"nm_id": 2389216,
		"brand": "Vivienne Sabo2",
		"status": 203		
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

func TestAdd(t *testing.T) {
	TestMemoryStorage := NewInMemory()
	var testOrder model.Order
	err := json.Unmarshal([]byte(testJSON), &testOrder)
	if err != nil {
		t.Errorf("Error unmarshalling json (%v)", err)
	}
	err = TestMemoryStorage.Add(&testOrder)
	if err != nil {
		t.Errorf("Error adding to map, %v", err)
	}
}

func TestSameUID(t *testing.T) {
	TestMemoryStorage := NewInMemory()
	var testOrder model.Order
	err := json.Unmarshal([]byte(testJSON), &testOrder)
	if err != nil {
		t.Errorf("Error unmarshalling json (%v)", err)
	}
	err = TestMemoryStorage.Add(&testOrder)
	if err != nil {
		t.Errorf("Error adding to map, %v", err)
	}
	err = TestMemoryStorage.Add(&testOrder)
	if err != model.ErrAlreadyExist {
		t.Errorf("Error with rejecting adding of the same UID %v", err)
	}
}

func TestFindByUID(t *testing.T) {
	TestMemoryStorage := NewInMemory()
	var testOrder model.Order
	var gotOrder *model.Order
	err := json.Unmarshal([]byte(testJSON), &testOrder)
	if err != nil {
		t.Errorf("Error unmarshalling json (%v)", err)
	}
	err = TestMemoryStorage.Add(&testOrder)
	if err != nil {
		t.Errorf("Error adding to map, %v", err)
	}
	gotOrder, err = TestMemoryStorage.FindByUID(testOrder.OrderUID)
	if err != nil {
		t.Errorf("Error with finding by UID %v", err)
	}
	if gotOrder.OrderUID != testOrder.OrderUID {
		t.Errorf("Orders are not the same. Got %v, want %v", *gotOrder, testOrder)
	}

	_, err = TestMemoryStorage.FindByUID("abcd")
	if err != model.ErrNotFound {
		t.Errorf("Wanted error not found, got %v", err)
	}
}

