package rewrite

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
	"strings"
)

func RewriteSQL(stmt ast.StmtNode) (ret string, err error){
	switch node := stmt.(type) {
	case *ast.CreateTableStmt:
		modifyPK2UK(node)
	default:
	}
	var sb strings.Builder
	err = stmt.Restore(format.NewRestoreCtx(format.DefaultRestoreFlags, &sb))
	return sb.String(), err
}

func modifyPK2UK(node *ast.CreateTableStmt) error {
	for _, column := range node.Cols {
		modifyColumnOption(column, ast.ColumnOptionPrimaryKey, ast.ColumnOptionUniqKey)
	}

	for _, constraint := range node.Constraints {
		if constraint.Tp == ast.ConstraintPrimaryKey {
			constraint.Tp = ast.ConstraintUniqKey
		}
	}
	return nil
}

func modifyColumnOption(column *ast.ColumnDef, b, a ast.ColumnOptionType) {
	for _, columnOption := range column.Options {
		if columnOption.Tp == b {
			columnOption.Tp = a
		}
	}
}
