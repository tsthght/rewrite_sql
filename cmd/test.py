#!/usr/bin/env python
import ctypes

lib = ctypes.CDLL('../bin/rewrite.so')
rewrite = lib.SQLRewrite
rewrite.argtypes = [ctypes.c_char_p]
rewrite.restype = ctypes.c_char_p
sql1 = b"create table t1 (a int, b int, primary key(a, b))"
print lib.SQLRewrite(sql1)

sql2 = b"create table t1 (a int, primary key(a))"
print lib.SQLRewrite(sql2)

sql3 = b"create table t1 (a int, unique key(a))"
print lib.SQLRewrite(sql3)

sql4 = b"create table t1 (a int, b varchar(32))"
print lib.SQLRewrite(sql4)

sql5 = b"create table t1 (a int, b varchar(32));create table t2 (a int, primary key(a))"
print lib.SQLRewrite(sql5)

sql6 = b"abc"
print lib.SQLRewrite(sql6)