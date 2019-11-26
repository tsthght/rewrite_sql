package ast

import (
	"fmt"
	"github.com/kr/pretty"
	"github.com/pingcap/parser/format"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	sqls := []string{
		`create table test(a int not null auto_increment,
                           primary key(id))`,
	}
	stmts, err := Parser(sqls[0], "", "")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	for _, stmt := range stmts {
		pretty.Println(stmt)
		var sb strings.Builder
		stmt.Restore(format.NewRestoreCtx(format.DefaultRestoreFlags, &sb))
		fmt.Printf("## %s\n", sb.String())
	}
}
