/* --------------------------------------------------------------
Project:    SQLite, RPC test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 19, 2021
Filename:   test_sqlite_rpc_client.go

Last modified:  April 19, 2021
License:

*
* Copyright (C) 2021 Ho-Jung Kim (godmode2k@hotmail.com)
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
-----------------------------------------------------------------
Note:
-----------------------------------------------------------------
Reference:
 - https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
 - https://github.com/mattn/go-sqlite3/issues/204
 - https://golang.org/pkg/net/rpc/
 - https://pkg.go.dev/database/sql
 - https://pkg.go.dev/github.com/mattn/go-sqlite3

Dependencies:
$ go get github.com/mattn/go-sqlite3


1. Build, Run:
	$ go run test_sqlite_rpc_server.go
	$ go run test_sqlite_rpc_client.go
-------------------------------------------------------------- */
package main



//! Header
// ---------------------------------------------------------------

import (
    "fmt"
    "log"

    "net/rpc"

    //"test_sqlite_golang/server"
)



//! Definition
// --------------------------------------------------------------------



//! Implementation
// --------------------------------------------------------------------

func main() {
    client, err := rpc.DialHTTP("tcp", ":1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    /*
    // Synchronous call
    args := &server.Args{7,8}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)


    // Asynchronous call
    quotient := new(Quotient)
    divCall := client.Go("Arith.Divide", args, quotient, nil)
    replyCall := <-divCall.Done	// will be equal to divCall
    // check errors, print, etc.
    */



    // Synchronous call
    var result string
    query := "select id, name from foo limit 1"
    err = client.Call( "SqliteDB.Test_sqlitedb", query, &result )
    if err != nil {
        log.Fatal( "RPC: error: ", err )
    }
    fmt.Printf( "RPC: result = %s\n", result )
}
