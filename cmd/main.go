package main

import (
	"github.com/tsthght/rewrite_sql/ast"
	"github.com/tsthght/rewrite_sql/rewrite"
)

import "C"
//export SQLRewrite
func SQLRewrite(sql *C.char) *C.char {
	var ret string = ""
	stmts, err := ast.Parser(C.GoString(sql), "", "")
	if err != nil {
		return C.CString(err.Error())
	}
	for _, stmt := range stmts {
		s, e := rewrite.RewriteSQL(stmt)
		if e != nil {
			return C.CString(err.Error())
		}
		ret += s
		ret += ";"
	}
	return C.CString(ret)
}

func main() {}
