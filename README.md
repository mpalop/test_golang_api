# test_golang_api

## Solution implemented 

I focused in a neat implementation of the requirements, meaning: 
* API that allows the ingest of a well defined JSON structures (Orders)
* Store this ingested Orders into Memory
* The system has some restrictions: 
  * You cannot reload an order with an Id that is already loaded into memory
  * You cannot load an Order whose line numbers are not:
    * ordered, starting by 1
    * consecutive. There is no gap allowed 
  * Obviously, the JSON code has to be well formed and matches the defined format

The implemented endpoints are the following ones: 

| Endpoint | Method | Description |
|---|---|---|
| /order | POST | This is the ingester. You can use it by:<br> curl -X POST -H "Content-Type: application/json" -d @sample1.json http://localhost:8000/order |
| /order | GET | This method shows all the Orders loaded into the system |
| /order/<order_id> | GET | This method shows the Order with Id = <order_id> | 

The possible erros that system shows are: 

| Endpoint | Method | Errors |
|---|---|---|
|/order | POST | 400 If the file is malformed, the order has been already loaded or the order of the lines is not valid |
|/order | GET | Never. If there is no Orders loaded, shows empty |
|/order/<order_id> | GET | 400 in case <order_id> is not an integer<br> 404 in case there is not loaded Order with Id = <order_id> |

Concerning the organization of the code, it spreads into three folders:
* models: Contains the structures and methods to work with the Orders 
* persistence: With all the logic to build a memory warehouse and manage the Orders 
* controllers: Has the code that raise the http server and prepares the endpoints and their code

All the test are into the folder test, as well some fixtures with sample Orders.

## Installation 

```bash 
cd $GOPATH
go get github.com/mpalop/sample_go_api 
```

## Dependencies: vendor approach

I got one dependency: the package httprouter from Julienschmid. 
To manage this dependency I use dep (check the files Gopkg.* on root directory)

```bash
test_golang_api git:(master) dep status
PROJECT                              CONSTRAINT  VERSION  REVISION  LATEST  PKGS USED
github.com/julienschmidt/httprouter  ^1.2.0      v1.2.0   348b672   v1.2.0  1
```

## Github

The repo is [https://github.com/mpalop/test_golang_api](https://github.com/mpalop/test_golang_api)
There is the code, sample fixtures and the current Readme file


## Makefile

There is a Makefile to help with the different parts of the app. 
The current implemented verbs are:
* **build**: to build the app 
* **run**: executes the app, starting the server that listens port 8000
* **run-with-samples**: the same as above, but starts the server (in background) and loads
a couple of orders to check the basic functions of the system
* **run-doc**: runs the golang documentation, showing in packages, in third party, 
the documentation generated of the app

There is some samples of the execution of the app

### Build the executable

Just use make: 
```bash
test_golang_api git:(master) make build
go build -o test_golang_api -v
```

### Run the executable and load fixtures 

In that case run the following make verb:
```bash
test_golang_api git:(master) make run-with-samples
./test_golang_api &
sleep 2
REST test. Use:
/order <POST> to add orders
/order <GET> to return the list of orders
/order/:orderId <GET> to get back the order <orderId> or 404 if not found, or 400 if <orderId> is not an int

starting server...
curl -X POST -H "Content-Type: application/json" -d @/Users/manel/work/go_base/src/github.com/mpalop/test_golang_api/tests/fixtures/sample1.json http://localhost:8000/order
Order 1 stored successfully
curl -X POST -H "Content-Type: application/json" -d @/Users/manel/work/go_base/src/github.com/mpalop/test_golang_api/tests/fixtures/sample2.json http://localhost:8000/order
Order 2 stored successfully
```

You can check the system running using curl, for instance: 

```bash
test_golang_api git:(master) curl http://localhost:8000/order
List of Orders

Order:     1, StoreId:    20
-line:  1 SKU: blue_sock
-line:  2 SKU: red_sock

Order:     2, StoreId:    30
-line:  1 SKU: yellow_sock
-line:  2 SKU: purple_sock
```

**NOTE**: Remember to stop the server using:  
```bash 
killall test_golang_api
```

*Yes, it is very provisional*

### Running the tests

This is an output of the execution of the tests: 
```bash 
test_golang_api git:(master) make test
go test -v ./...
?   	github.com/mpalop/test_golang_api	[no test files]
?   	github.com/mpalop/test_golang_api/controllers	[no test files]
?   	github.com/mpalop/test_golang_api/models	[no test files]
?   	github.com/mpalop/test_golang_api/persistence	[no test files]
=== RUN   TestParseOrder
--- PASS: TestParseOrder (0.00s)
    orders_test.go:11: Testing normal file
    orders_test.go:16: Order processed OK:
         Order:     1, StoreId:    20
        -line:  1 SKU: blue_sock
        -line:  2 SKU: red_sock
    orders_test.go:20: Testing bad syntax file
    orders_test.go:25: Detected bad syntax OK. Error invalid character '{' looking for beginning of object key string
    orders_test.go:30: Testing bad order lines file
    orders_test.go:35: Detected bad order lines OK. Error line numbers are not ordered
    orders_test.go:40: Testing bad structure file
    orders_test.go:45: Detected bad structure OK. Error json: cannot unmarshal string into Go struct field Order.id of type int
=== RUN   TestPersistence
--- PASS: TestPersistence (0.00s)
    persistence_test.go:15: Testing store normal file
    persistence_test.go:26: Order 1 was saved correctly. OK
    persistence_test.go:37: Testing store normal file twice
    persistence_test.go:40: Order 1 returns error when tried to save twice: Order 1 already exists. OK
PASS
ok  	github.com/mpalop/test_golang_api/tests	(cached)
?   	github.com/mpalop/test_golang_api/vendor	[no test files]
```

## Documentation 

The code is documented in order to allow godoc integrate it into the standard documentation
to check that, just run: 
```bash
test_golang_api git:(master) make run-doc
godoc -http=:6060 -http=:6060
```

Then, visit with the browser the url: [http://127.0.0.1:6060/pkg/](http://127.0.0.1:6060/pkg/) and look in "Third Party" zone
