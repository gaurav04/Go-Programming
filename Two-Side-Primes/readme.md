First Install the go package mux
```
go get -u github.com/gorilla/mux
```
Then run
```
go get github.com/gaurav04/Go-Programming/Two-Side-Primes
```

Start the http server by running the below command  
```
cd $GOPATH/src/github.com/gaurav04/Go-Programming/Two-Side-Primes
go run main.go
```

Once server start on port 8080, run the following command from the browser.
```
http://localhost:8080/{any_number}

like: http://localhost:8080/3797
```

Expected Ouput:

```
The number is Two-sided Prime: True/False
```
