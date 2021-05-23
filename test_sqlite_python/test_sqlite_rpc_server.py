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
#
# -----------------------------------------------------------------



from include.server import *
import threading



def main():
    rpc_server = CServer()

    # JSON-RPC
    th_jsonrpc = threading.Thread( target = rpc_server.run_jsonrpc_server )
    th_jsonrpc.start()
    th_jsonrpc.join()



if __name__ == "__main__":
    main()

