CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "carts" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE IF NOT EXISTS "products" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"name" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, "price_cents" int NOT NULL DEFAULT '0');
CREATE TABLE IF NOT EXISTS "cart_items" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"cart_id" integer NOT NULL,
"product_id" integer NOT NULL,
"quantity" uint NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL, "standard_price_cents" int NOT NULL DEFAULT '0', "discounted_price_cents" int NOT NULL DEFAULT '0',
FOREIGN KEY (cart_id) REFERENCES carts (id) ON DELETE cascade,
FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE cascade
);
CREATE INDEX "cart_items_cart_id_idx" ON "cart_items" (cart_id);
CREATE INDEX "cart_items_product_id_idx" ON "cart_items" (product_id);
CREATE TABLE IF NOT EXISTS "transactions" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "stock_entries" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"transaction_id" integer NOT NULL,
"product_id" integer NOT NULL,
"quantity" integer NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (transaction_id) REFERENCES transactions (id) ON DELETE cascade,
FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE cascade
);
CREATE INDEX "stock_entries_transaction_id_idx" ON "stock_entries" (transaction_id);
CREATE INDEX "stock_entries_product_id_idx" ON "stock_entries" (product_id);
CREATE TABLE IF NOT EXISTS "promotions" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"label" TEXT NOT NULL,
"required_product_id" integer NOT NULL,
"required_product_quantity" uint NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL, "discount_type" TEXT, "discount_id" integer,
FOREIGN KEY (required_product_id) REFERENCES products (id) ON DELETE cascade
);
CREATE UNIQUE INDEX "promotions_label_idx" ON "promotions" (label);
CREATE INDEX "promotions_required_product_id_idx" ON "promotions" (required_product_id);
CREATE TABLE IF NOT EXISTS "percentage_discounts" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"amount" decimal NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "pinned_discounts" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"amount_cents" uint NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE INDEX "index_promotions_on_discount_type_discount_id" ON "promotions" (discount_type, discount_id);
CREATE TABLE IF NOT EXISTS "promotion_discounted_products" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"product_id" integer NOT NULL,
"promotion_id" integer NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE cascade,
FOREIGN KEY (promotion_id) REFERENCES promotions (id) ON DELETE cascade
);
CREATE INDEX "promotion_discounted_products_product_id_idx" ON "promotion_discounted_products" (product_id);
CREATE INDEX "promotion_discounted_products_promotion_id_idx" ON "promotion_discounted_products" (promotion_id);
