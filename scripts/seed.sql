-- Intended to be run on development database to seed it with
-- initial data set.

INSERT INTO products
  (created_at,  updated_at,  name,       price_cents)
VALUES
  (time('now'), time('now'), "Belts",    2000),
  (time('now'), time('now'), "Shirts",   6000),
  (time('now'), time('now'), "Suits",    30000),
  (time('now'), time('now'), "Trousers", 7000),
  (time('now'), time('now'), "Shoes",    12000),
  (time('now'), time('now'), "Ties",     2000);

INSERT INTO transactions (created_at, updated_at) VALUES (time('now'), time('now'));

INSERT INTO stock_entries
  (
    created_at,
    updated_at,
    transaction_id,
    product_id,
    quantity
  )
VALUES
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions LIMIT 1),
    (SELECT id FROM products WHERE name = "Belts"),
    10
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions LIMIT 1),
    (SELECT id FROM products WHERE name = "Shirts"),
    5
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions LIMIT 1),
    (SELECT id FROM products WHERE name = "Suits"),
    2
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions LIMIT 1),
    (SELECT id FROM products WHERE name = "Trousers"),
    4
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions LIMIT 1),
    (SELECT id FROM products WHERE name = "Shoes"),
    1
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions LIMIT 1),
    (SELECT id FROM products WHERE name = "Ties"),
    8
  );
