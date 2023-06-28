CREATE TABLE IF NOT EXISTS ORDERS(
    order_uid varchar(20) PRIMARY KEY,
    track_number varchar(30) NOT NULL UNIQUE,
    entry varchar(20) NOT NULL,
    locale varchar(3) NOT NULL,
    internal_signature varchar(30),
    delivery_service varchar(30) NOT NULL,
    shardkey varchar(30) NOT NULL,
    sm_id BIGINT NOT NULL,
    date_created TIMESTAMP NOT NULL,
    oof_shard varchar(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS DELIVERY(
    order_id varchar(20) REFERENCES ORDERS(order_uid),
    name varchar(255) NOT NULL,
    phone varchar(30) NOT NULL,
    zip varchar(30) NOT NULL,
    city varchar(30) NOT NULL,
    address varchar(50) NOT NULL,
    region varchar(50) NOT NULL,
    email varchar(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS ITEMS(
    chrt_id BIGINT NOT NULL,
    track_number varchar(30) REFERENCES ORDERS(track_number),
    price int NOT NULL,
    rid varchar(50) NOT NULL,
    name varchar(50) NOT NULL,
    sale int NOT NULL,
    size varchar(5) NOT NULL,
    total_price int NOT NULL,
    nm_id BIGINT NOT NULL,
    brand varchar(50) NOT NULL,
    status int NOT NULL
);

CREATE TABLE IF NOT EXISTS PAYMENT(
    transaction_code varchar(20) REFERENCES ORDERS,
    request_id varchar(20) NOT NULL,
    currency varchar(10) NOT NULL,
    provider varchar(30) NOT NULL,
    amount integer NOT NULL,
    payment_dt BIGINT NOT NULL,
    bank varchar(30),
    delivery_cost INT NOT NULL,
    goods_total INT NOT NULL,
    custom_fee INT NOT NULL
);


CREATE TABLE IF NOT EXISTS orders(
    order_uid varchar PRIMARY KEY
    order_data jsonb NOT NULL
)