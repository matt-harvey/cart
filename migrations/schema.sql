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
);
