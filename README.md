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

