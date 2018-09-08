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
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (cart_id) REFERENCES carts (id) ON DELETE cascade,
FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE cascade
);
CREATE INDEX "cart_items_cart_id_idx" ON "cart_items" (cart_id);
CREATE INDEX "cart_items_product_id_idx" ON "cart_items" (product_id);
