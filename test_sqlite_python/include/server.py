#!/usr/bin/python3
# -*- coding: utf-8 -*-



# -----------------------------------------------------------------
# Purpose: SQLite, RPC test
# Author: Ho-Jung Kim (godmode2k@hotmail.com)
# Filename: test_sqlite_rpc_server.py
# Date: Since April 14, 2021
#
#
# Reference:
#
#
# License:
#
# *
# * Copyright (C) 2021 Ho-Jung Kim (godmode2k@hotmail.com)
# *
# * Licensed under the Apache License, Version 2.0 (the "License");
# * you may not use this file except in compliance with the License.
# * You may obtain a copy of the License at
# *
# *      http://www.apache.org/licenses/LICENSE-2.0
# *
# * Unless required by applicable law or agreed to in writing, software
# * distributed under the License is distributed on an "AS IS" BASIS,
# * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# * See the License for the specific language governing permissions and
# * limitations under the License.
# *
# -----------------------------------------------------------------
# Note:
# -----------------------------------------------------------------
# Dependencies:
# $ pip3 install pysqlite3
# $ pip3 install jsonrpclib
# -----------------------------------------------------------------



# JSON-RPC
# https://github.com/joshmarshall/jsonrpclib
#import jsonrpclib
from jsonrpclib.SimpleJSONRPCServer import SimpleJSONRPCServer
import sqlite3



class CServer:
    def __init__(self):
        pass

    def test_sqlite(self):
        result = ""
        db_filename = "test_sqlite3.db"
        #conn = sqlite3.connect( db_filename )
        conn = sqlite3.connect( "file:?mode=memory&cache=shared" )


        cur = conn.cursor()

        sql = "CREATE TABLE testdb (id integer not null primary key, f1 text, f2 text);"
        cur.execute( sql )
        conn.commit()

        sql = "INSERT INTO testdb (f1, f2) VALUES (?, ?)"
        cur.execute( sql, ('test1', 'test2') )
        conn.commit()

        cur = conn.cursor()
        cur.execute( "SELECT * FROM testdb" )
        rows = cur.fetchall()
        for row in rows:
            print( row )
            result = row
        conn.close()

        return str(result)

    def rpc_call__test(self):
        result = "test_call"
        #return json.dumps(result).encode("utf-8")
        return '{"result": "' + result + '"}'

    def rpc_call__test_sqlite(self):
        result = self.test_sqlite()
        #return json.dumps(result).encode("utf-8")
        return '{"result": "' + result + '"}'


    def run_jsonrpc_server(self):
        # https://github.com/joshmarshall/jsonrpclib

        # e.g.,
        # $ curl -X POST --data '{"jsonrpc":"2.0","method":"<method name>","params":[{"key":"value"}],"id":0}' -H "Content-Type: application/json" http://127.0.0.1:8080/
        # $ curl -X POST --data '{"jsonrpc":"2.0","method":"test","params":[],"id":0}' -H "Content-Type: application/json" http://127.0.0.1:8080/
        # $ curl -X POST --data '{"jsonrpc":"2.0","method":"test_sqlite","params":[],"id":0}' -H "Content-Type: application/json" http://127.0.0.1:8080/

        print( "JSON-RPC Server starting..." )
        server = SimpleJSONRPCServer( ('0.0.0.0', 8080) )


        server.register_function( self.rpc_call__test, "test" )
        server.register_function( self.rpc_call__test_sqlite, "test_sqlite" )

        # e.g.,
        #server.register_function( pow )
        #server.register_function( lambda x, y: x + y, 'add' )
        #server.register_function( lambda x: x, 'ping' )

        print( "running..." )
        server.serve_forever()


