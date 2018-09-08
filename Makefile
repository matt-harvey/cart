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
	soda migrate up && soda migrate -e test up

.PHONY: dbmigratedown
dbmigratedown:
	soda migrate down && soda migrate -e test down

.PHONY: dbdrop
dbdrop:
	soda drop -a

.PHONY: dbcreate
dbcreate:
	soda create -a

.PHONY: dbreset
dbreset: dbdrop dbcreate dbmigrate

vendor: Gopkg.toml Gopkg.lock
	echo Installing dependencies
	dep ensure
