package intcast

import (
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/comment"
	"github.com/gostaticanalysis/comment/passes/commentmap"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "intcast finds integer type cast that can cause overflow"

var Analyzer = &analysis.Analyzer{
	Name: "intcast",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
		commentmap.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	cmaps := pass.ResultOf[commentmap.Analyzer].(comment.Maps)

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CallExpr:
			f := n.Fun
			args := n.Args
			switch f := f.(type) {
			case *ast.Ident:
				if isIntegerCast(f, pass) {
					if len(args) == 1 {
						arg := args[0]
						argType := pass.TypesInfo.TypeOf(arg)
						if isIntegerArg(arg, pass) {
							srcBasicType := argType.Underlying().(*types.Basic)
							dstBasicType := pass.TypesInfo.TypeOf(f).Underlying().(*types.Basic)
							if isProblematicIntegerCast(srcBasicType, dstBasicType) {
								line := pass.Fset.Position(f.Pos()).Line
								if !cmaps.IgnoreLine(pass.Fset, line, "intcast") {
									pass.Reportf(f.Pos(), "problematic integer type cast from %s to %s", argType, f)
								}
							}
						}
					}
				}
			case *ast.SelectorExpr:
				if isIntegerCast(f.Sel, pass) {
					if len(args) == 1 {
						arg := args[0]
						argType := pass.TypesInfo.TypeOf(arg)
						if isIntegerArg(arg, pass) {
							srcBasicType := argType.Underlying().(*types.Basic)
							dstBasicType := pass.TypesInfo.TypeOf(f).Underlying().(*types.Basic)
							if isProblematicIntegerCast(srcBasicType, dstBasicType) {
								if isProblematicIntegerCast(srcBasicType, dstBasicType) {
									line := pass.Fset.Position(f.Pos()).Line
									if !cmaps.IgnoreLine(pass.Fset, line, "intcast") {
										pass.Reportf(f.Pos(), "problematic integer type cast from %s to %s.%s", argType, f.X, f.Sel)
									}
								}
							}
						}
					}
				}
			}
		}
	})
	return nil, nil
}

func isIntegerCast(ident *ast.Ident, pass *analysis.Pass) bool {
	typ := pass.TypesInfo.TypeOf(ident).Underlying()
	basic, ok := typ.(*types.Basic)
	if !ok {
		return false
	}
	return basic.Info()&types.IsInteger != 0
}

func isIntegerArg(expr ast.Expr, pass *analysis.Pass) bool {
	typ := pass.TypesInfo.TypeOf(expr).Underlying()
	basic, ok := typ.(*types.Basic)
	if !ok {
		return false
	}
	return basic.Info()&types.IsInteger != 0
}

func isProblematicIntegerCast(src *types.Basic, dst *types.Basic) bool {
	for _, cast := range problematicCast {
		if src.Kind() == cast[0] && dst.Kind() == cast[1] {
			return true
		}
	}
	return false
}

// [0] is source type, [1] is destination type
var problematicCast [][2]types.BasicKind = [][2]types.BasicKind{
	// int
	{types.Int, types.Int16},
	{types.Int, types.Int32},
	{types.Int, types.Int8},
	{types.Int, types.Uint},
	{types.Int, types.Uint16},
	{types.Int, types.Uint32},
	{types.Int, types.Uint64},
	{types.Int, types.Uint8},
	// int16
	{types.Int16, types.Int8},
	{types.Int16, types.Uint},
	{types.Int16, types.Uint16},
	{types.Int16, types.Uint32},
	{types.Int16, types.Uint64},
	{types.Int16, types.Uint8},
	// int32
	{types.Int32, types.Int16},
	{types.Int32, types.Int8},
	{types.Int32, types.Uint},
	{types.Int32, types.Uint16},
	{types.Int32, types.Uint32},
	{types.Int32, types.Uint64},
	{types.Int32, types.Uint8},
	// int64
	{types.Int64, types.Int16},
	{types.Int64, types.Int32},
	{types.Int64, types.Int8},
	{types.Int64, types.Uint},
	{types.Int64, types.Uint16},
	{types.Int64, types.Uint32},
	{types.Int64, types.Uint64},
	{types.Int64, types.Uint8},
	// int8
	{types.Int8, types.Uint},
	{types.Int8, types.Uint16},
	{types.Int8, types.Uint32},
	{types.Int8, types.Uint64},
	{types.Int8, types.Uint8},
	// uint
	{types.Uint, types.Int},
	{types.Uint, types.Int16},
	{types.Uint, types.Int32},
	{types.Uint, types.Int8},
	{types.Uint, types.Int64},
	{types.Uint, types.Uint16},
	{types.Uint, types.Uint32},
	{types.Uint, types.Uint8},
	// uint16
	{types.Uint16, types.Int16},
	{types.Uint16, types.Int8},
	{types.Uint16, types.Uint8},
	// uint32
	{types.Uint32, types.Int},
	{types.Uint32, types.Int16},
	{types.Uint32, types.Int32},
	{types.Uint32, types.Int8},
	{types.Uint32, types.Uint16},
	{types.Uint32, types.Uint8},
	// uint64
	{types.Uint64, types.Int},
	{types.Uint64, types.Int16},
	{types.Uint64, types.Int32},
	{types.Uint64, types.Int64},
	{types.Uint64, types.Int8},
	{types.Uint64, types.Uint16},
	{types.Uint64, types.Uint32},
	{types.Uint64, types.Uint8},
	// uint8
	{types.Uint8, types.Int8},
}
