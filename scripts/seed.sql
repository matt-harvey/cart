-- Intended to be run on development database to seed it with
-- initial data set.

INSERT INTO products
  (created_at, updated_at, name, price_cents)
VALUES
  (time('now'), time('now'), 'Belts', 2000),
  (time('now'), time('now'), 'Shirts', 6000),
  (time('now'), time('now'), 'Suits', 30000),
  (time('now'), time('now'), 'Trousers', 7000),
  (time('now'), time('now'), 'Shoes', 12000),
  (time('now'), time('now'), 'Ties', 2000);

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
    (SELECT id FROM transactions ORDER BY id DESC LIMIT 1),
    (SELECT id FROM products WHERE name = 'Belts'),
    10
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions ORDER BY id DESC LIMIT 1),
    (SELECT id FROM products WHERE name = 'Shirts'),
    5
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions ORDER BY id DESC LIMIT 1),
    (SELECT id FROM products WHERE name = 'Suits'),
    2
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions ORDER BY id DESC LIMIT 1),
    (SELECT id FROM products WHERE name = 'Trousers'),
    4
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions ORDER BY id DESC LIMIT 1),
    (SELECT id FROM products WHERE name = 'Shoes'),
    1
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM transactions ORDER BY id DESC LIMIT 1),
    (SELECT id FROM products WHERE name = 'Ties'),
    8
  );

INSERT INTO carts (created_at, updated_at) VALUES (time('now'), time('now'));

INSERT INTO cart_items
  (
    created_at,
    updated_at,
    cart_id,
    product_id,
    quantity,
    standard_price_cents, -- FIXME This should be calculated
    discounted_price_cents -- FIXME This should be calculated
  )
VALUES
  (
    time('now'),
    time('now'),
    (SELECT id FROM carts ORDER BY id DESC LIMIT 1),
    (SELECT id FROM products WHERE name = 'Trousers'),
    3,
    14000,
    14000
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM carts ORDER BY id DESC LIMIT 1),
    (SELECT id FROM products WHERE name = 'Belts'),
    1,
    2000,
    1700
  );

INSERT INTO percentage_discounts
  (id, created_at, updated_at, amount)
VALUES
  (1, time('now'), time('now'), 0.15),
  (2, time('now'), time('now'), 0.5);

INSERT INTO pinned_discounts
  (id, created_at, updated_at, amount_cents)
VALUES
  (1, time('now'), time('now'), 4500);

INSERT INTO promotions
  (
    label,
    created_at,
    updated_at,
    required_product_id,
    required_product_quantity,
    discount_type,
    discount_id
  )
VALUES
  (
    'trousers_belts_shoes',
    time('now'),
    time('now'),
    (SELECT id FROM products WHERE name = 'Trousers'),
    2,
    "PercentageDiscount",
    1
  ),
  (
    'shirts_45',
    time('now'),
    time('now'),
    (SELECT id FROM products WHERE name = 'Shirts'),
    2,
    "FixedDiscount",
    1
  ),
  (
    'shirts_ties',
    time('now'),
    time('now'),
    (SELECT id FROM products WHERE name = 'Shirts'),
    3,
    "PercentageDiscount",
    2
  );

INSERT INTO promotion_discounted_products
  (
    created_at,
    updated_at,
    promotion_id,
    product_id
  )
VALUES
  (
    time('now'),
    time('now'),
    (SELECT id FROM promotions WHERE label = 'trousers_belts_shoes'),
    (SELECT id FROM products WHERE name = 'Belts')
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM promotions WHERE label = 'trousers_belts_shoes'),
    (SELECT id FROM products WHERE name = 'Shoes')
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM promotions WHERE label = 'shirts_45'),
    (SELECT id FROM products WHERE name = 'Shirts')
  ),
  (
    time('now'),
    time('now'),
    (SELECT id FROM promotions WHERE label = 'shirts_ties'),
    (SELECT id FROM products WHERE name = 'Ties')
  );
