/* --------------------------------------------------------------
Project:    SQLite, RPC test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 19, 2021
Filename:   test_sqlite_rpc_server.go

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
    "database/sql"

    //_ "github.com/mattn/go-sqlite3"

    "net"
    "net/http"
    "net/rpc"
    "test_sqlite_golang/server"
)



//! Definition
// --------------------------------------------------------------------



//! Implementation
// --------------------------------------------------------------------

/*
// Source: https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
func test_sqlite() {
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    sqlStmt := `
    create table foo (id integer not null primary key, name text);
    delete from foo;
    `
    _, err = db.Exec(sqlStmt)
    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
        return
    }

    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()
    for i := 0; i < 100; i++ {
        _, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
        if err != nil {
            log.Fatal(err)
        }
    }
    tx.Commit()

    rows, err := db.Query("select id, name from foo")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    stmt, err = db.Prepare("select name from foo where id = ?")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()
    var name string
    err = stmt.QueryRow("3").Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)

    _, err = db.Exec("delete from foo")
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
    if err != nil {
        log.Fatal(err)
    }

    rows, err = db.Query("select id, name from foo")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}
*/


func run_rpc_server() {
    // SQLite
    // ------------------------------------------

    //db, err := sql.Open("sqlite3", "./foo.db")
    //db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
    //db_filename := randomString(16) // func creates random string
    db_filename := "test_sqlite3.db"
    db, err := sql.Open( "sqlite3", fmt.Sprintf("file:%s?mode=memory&cache=shared", db_filename) )
    //fmt.Print( "%v", db )
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()



    sqlStmt := `
    CREATE TABLE foo (id integer not null primary key, name text);
    DELETE FROM foo;
    `
    _, err = db.Exec( sqlStmt )
    if err != nil {
        log.Printf( "%q: %s\n", err, sqlStmt )
        return
    }

    tx, err := db.Begin()
    if err != nil {
        log.Fatal( err )
    }
    stmt, err := tx.Prepare( "INSERT INTO foo(id, name) VALUES(?, ?)" )
    if err != nil {
        log.Fatal( err )
    }
    defer stmt.Close()
    for i := 0; i < 100; i++ {
        _, err = stmt.Exec( i, fmt.Sprintf("こんにちわ世界%03d", i) )
        if err != nil {
            log.Fatal( err )
        }
    }
    tx.Commit()



    // RPC
    // ------------------------------------------

    sqlitedb := new( server.SqliteDB )
    sqlitedb.Db = db

    rpc.Register( sqlitedb )
    rpc.HandleHTTP()

    l, e := net.Listen( "tcp", ":1234" )
    if e != nil {
        log.Fatal( "listen error:", e )
    }

    //go http.Serve( l, nil )
    http.Serve( l, nil )
}


func main() {
    //test_sqlite()

    run_rpc_server()
}
