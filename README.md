# SQLite3 test


Summary
----------
>



Environment
----------
```sh
Python: v3.7.4
Golang: v1.15
```



Installation and Run
----------
```sh
[Python verion]
$ pip3 install pysqlite3
$ pip3 install jsonrpclib

$ cd test_sqlite_python

// Server
$ python3 test_sqlite_rpc_server.py

// Client
$ curl -X POST --data '{"jsonrpc":"2.0","method":"test","params":[],"id":0}' -H "Content-Type: application/json" http://127.0.0.1:8080/
$ curl -X POST --data '{"jsonrpc":"2.0","method":"test_sqlite","params":[],"id":0}' -H "Content-Type: application/json" http://127.0.0.1:8080/



[Golang version]
$ go get github.com/mattn/go-sqlite3

$ cd test_sqlite_golang

// Server
$ go run test_sqlite_rpc_server.go

// Client
$ go run test_sqlite_rpc_client.go

```
