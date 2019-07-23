// Copyright 2019 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package builtin

import (
	"go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func pre(mustFound *bool) func(c *astutil.Cursor) bool {
	return func(c *astutil.Cursor) bool {

		currentNode := c.Node()
		parentNode := c.Parent()
		idx := c.Index() // Can be negative

		switch n := parentNode.(type) { // Must be parent to access Index
		case *ast.BlockStmt:

			if len(chain) == 0 || current().ref != n {
				// not current so update latest to chain
				chain = append(chain, tracker{n, idx})
			} else {
				if idx > current().idx {
					chain = append(chain, tracker{n, idx})
				}
			}
		case *ast.CommClause:

			if len(chain) == 0 || current().ref != n {
				// not current so update latest to chain
				chain = append(chain, tracker{n, idx})
			} else {
				if idx > current().idx {
					chain = append(chain, tracker{n, idx})
				}
			}
		case *ast.CaseClause:

			if len(chain) == 0 || current().ref != n {
				// not current so update latest to chain
				chain = append(chain, tracker{n, idx})
			} else {
				if idx > current().idx {
					chain = append(chain, tracker{n, idx})
				}
			}
		}

		switch n := currentNode.(type) {
		case *ast.CallExpr:
			ok, arg1, arg2 := isMustFunc(n)
			if ok {
				*mustFound = false
				replaceMustFunc(c, arg1, arg2)
			}

		}

		return true
	}
}

func post(c *astutil.Cursor) bool {

	currentNode := c.Node()

	switch n := currentNode.(type) {
	case *ast.BlockStmt, *ast.CommClause, *ast.CaseClause:
		pop(n)
	}

	return true
}
