# Cart: a toy shopping cart server

## Development setup

To run this application, you will need:

* Golang installed.
* Your `GOPATH` and `GOROOT` environment variables set, and the `go` binaries in your `PATH`. Placing
  this in your ~/.bash (or ~/.zshrc or etc.) will ensure this is the case:
  ```
  export GOPATH="${HOME}/go"
  export GOROOT="$(brew --prefix golang)/libexec"
  export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin:${GOPATH}/bin"
  ```
* sqlite3 installed
* The "pop" and "soda" command line tools installed, with sqlite3 support:
  ```
  go get -u -v -tags sqlite github.com/gobuffalo/pop/...
  go install -tags sqlite github.com/gobuffalo/pop/soda
  ```

### Installation

Clone this repo:

```
mkdir -p $GOPATH/src/github.com/matt-harvey
cd $GOPATH/src/github.com/matt-harvey
git clone git@github.com:matt-harvey/cart.git
```

Now `cd cart` to start working.

### Setup

Set up unversioned config files:

```
cp .env.template .env
cp config/database.yml.template config/database.yml
```

Adjust `.env` and `config/database.yml` if you're not happy with ENV variables and
database paths etc..

Now run `make dbsetup` to set up the development and test databases. This will also
seed the database with a cart and some products and discounts.

### Testing

#### Automated tests

`make test`

**NOTE**: Proper post-test database cleaning is not yet in place, so you may want to run
`make dbreset` to start the test database on a fresh footing before running `make test` on
subsequent occasions.

#### Manually testing using cURL

Start the server using `make run`. Then in a separate console tab, test using cURL as
follows:

##### View a given cart

`curl http://localhost:3000/carts/1`

This should output a JSON representation of the currently seeded cart, which should contain
3 trousers and 1 belt, with the belt subject to a 15% discount. Note prices are calculated in
cents (to avoid issues around base-2 floating point arithetic); and the price for each
"item" in the cart represents to total price for all products of that type in the cart.

##### Add various items to the cart

```
curl -X PATCH -H 'Content-Type: application/json' -d '{"product_id":2,"quantity":2}' http://localhost:3000/carts/1/adjust_items
curl -X PATCH -H 'Content-Type: application/json' -d '{"product_id":6,"quantity":2}' http://localhost:3000/carts/1/adjust_items
curl -X PATCH -H 'Content-Type: application/json' -d '{"product_id":1,"quantity":1}' http://localhost:3000/carts/1/adjust_items
```

If you now run `curl http://localhost:3000/carts/1`, you should see that 2 shirts, 2 ties and a
extra belt have been added to the existing cart. Note that the ties and shirts are not yet subject to a
discount.

Now add another shirt:

```
curl -X PATCH -H 'Content-Type: application/json' -d '{"product_id":2,"quantity":1}' http://localhost:3000/carts/1/adjust_items
```

Now run `curl http://localhost:3000/carts/1` again, and notice that a 50% discount has now been
applied to the ties, in virtue of 3 shirts now being included in the cart; and that the third shirt,
moreover, is now discounted to $45 ("4500 cents").

##### Remove items from the cart

Items can be removed from a cart by PATCH-ing a request to the same endpoint as used for adding items to the cart,
but with negative quantities representing the desired adjustment. For example:

```
curl -X PATCH -H 'Content-Type: application/json' -d '{"product_id":2,"quantity":-1}' http://localhost:3000/carts/1/adjust_items
```

Run `curl http://localhost:3000/carts/1` again to see the effect of the adjustment.

Experiment with different quantities to see what happens when you try to remove items that are not
present in the cart, or to remove larger quantity of a given product than is present in the cart.

