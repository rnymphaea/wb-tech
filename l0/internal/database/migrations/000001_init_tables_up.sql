CREATE TABLE IF NOT EXISTS orders (
  uid VARCHAR(20) PRIMARY KEY,
  track_number VARCHAR(20) NOT NULL,
  entry VARCHAR(5) NOT NULL,
  locale VARCHAR(2) NOT NULL,
  internal_signature TEXT,
  customer_id TEXT NOT NULL,
  delivery_service TEXT NOT NULL, 
  shardkey TEXT,
  sm_id INTEGER NOT NULL,
  date_created TIMESTAMPTZ NOT NULL,
  oof_shard TEXT
);

CREATE TABLE IF NOT EXISTS deliveries (
  order_uid VARCHAR(20) NOT NULL, 
  name TEXT,
  phone TEXT,
  zip TEXT,
  city TEXT NOT NULL,
  address TEXT NOT NULL,
  region TEXT NOT NULL,
  email TEXT,
  FOREIGN KEY (order_uid) REFERENCES orders(uid) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS payments (
  order_uid VARCHAR(20) NOT NULL,
  transaction TEXT NOT NULL,
  request_id TEXT NOT NULL,
  currency VARCHAR(4) NOT NULL,
  provider TEXT NOT NULL,
  amount INTEGER NOT NULL,
  payment_dt BIGINT NOT NULL,
  bank TExT NOT NULL,
  delivery_cost INTEGER NOT NULL,
  goods_total INTEGER NOT NULL,
  custom_fee INTEGER,
  FOREIGN KEY (order_uid) REFERENCES orders(uid) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS items (
  order_uid VARCHAR(20) NOT NULL,
  chrt_id BIGINT NOT NULL,
  track_number TEXT NOT NULL,
  price INTEGER NOT NULL,
  rid TEXT NOT NULL,
  name TEXT NOT NULL,
  sale INTEGER NOT NULL,
  size TEXT,
  total_price INTEGER NOT NULL,
  nm_id BIGINT NOT NULL,
  brand TEXT NOT NULL,
  status INTEGER NOT NULL,
  FOREIGN KEY (order_uid) REFERENCES orders(uid) ON DELETE CASCADE
);



