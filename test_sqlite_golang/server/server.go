/* --------------------------------------------------------------
Project:    SQLite, RPC test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 19, 2021
Filename:   server.go

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
package server



//! Header
// ---------------------------------------------------------------

import (
    // SQLite
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    //"os"

    //"errors"
)



//! Definition
// --------------------------------------------------------------------



//! Implementation
// --------------------------------------------------------------------

type SqliteDB struct {
    Db *sql.DB
}

func (t *SqliteDB) Test_sqlitedb(query string, response* string) error {
    if t.Db == nil {
        fmt.Println( "SQLite: connection lost" )
        return nil
    }

    if len(query) <= 0 {
        fmt.Println( "SQLite: query is empty" )
        return nil
    }

    var result string

    //rows, err := t.Db.Query("select id, name from foo limit 1")
    rows, err := t.Db.Query( query )
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

        result = fmt.Sprintf( "id = %d, name = %s", id, name )
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    *response = result

    return nil
}
