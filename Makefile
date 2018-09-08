# Environment variable default values.
BINARY_NAME=cart

# Default target.
.PHONY: all
all: build

# Unless VERBOSE env var is defined, do not output executed commands.
# Warning: do _not_ put this at top above default target.
ifndef VERBOSE
.SILENT:
endif

.PHONY: build
build: vendor
	echo Compiling binary
	go build -o $(BINARY_NAME) -tags sqlite

.PHONY: test
test: vendor
	echo Running tests
	ENV=test go test -tags sqlite ./...

.PHONY: clean
clean:
	echo Removing compiled binaries
	go clean
	echo Cleaning cached tests
	go clean -testcache
	echo Uninstalling dependencies
	rm -rf vendor

.PHONY: run
run: build
	./$(BINARY_NAME)

.PHONY: dbmigrate
dbmigrate:
	echo Migrating databases up
	soda migrate up && soda migrate -e test up

.PHONY: dbmigratedown
dbmigratedown:
	echo Migrating databases down
	soda migrate down && soda migrate -e test down

.PHONY: dbdrop
dbdrop:
	echo Dropping databases
	soda drop -a && rm -f dbseed

.PHONY: dbcreate
dbcreate:
	echo Creating databases
	soda create -a

dbseed:
	echo Seeding development database
	sqlite3 /tmp/cart_development.sqlite < scripts/seed.sql && touch dbseed

.PHONY: dbreset
dbreset: dbdrop dbcreate dbmigrate dbseed

vendor: Gopkg.toml Gopkg.lock
	echo Installing dependencies
	dep ensure
