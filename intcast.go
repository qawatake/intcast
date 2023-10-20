package intcast

import (
	"go/ast"
	"go/types"
	"slices"
	"strconv"
	"strings"

	"github.com/gostaticanalysis/comment"
	"github.com/gostaticanalysis/comment/passes/commentmap"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "intcast identifies integer type casts that can potentially cause overflow"

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
									pass.Reportf(f.Pos(), "unsafe cast: converting %s to %s could lead to integer overflow.", trimPackage(argType.String()), trimPackage(f.String()))
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
								line := pass.Fset.Position(f.Pos()).Line
								if !cmaps.IgnoreLine(pass.Fset, line, "intcast") {
									dstType := pass.TypesInfo.TypeOf(f)
									pass.Reportf(f.Pos(), "unsafe cast: converting %s to %s could lead to integer overflow.", trimPackage(argType.String()), trimPackage(dstType.String()))
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
	for _, cast := range problematicCasts {
		if src.Kind() == cast.from && dst.Kind() == cast.to {
			return true
		}
	}
	return false
}

var problematicCasts []cast = []cast{
	// int
	{types.Int, types.Int16},
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

type cast struct {
	from types.BasicKind
	to   types.BasicKind
}

func init() {
	if strconv.IntSize == 64 {
		problematicCasts = slices.Grow(problematicCasts, 1)
		problematicCasts = append(problematicCasts,
			cast{types.Int, types.Int32},
		)
	}
	if strconv.IntSize == 32 {
		problematicCasts = slices.Grow(problematicCasts, 2)
		problematicCasts = append(problematicCasts,
			cast{types.Int64, types.Int},
			cast{types.Uint32, types.Int},
			cast{types.Uint32, types.Int},
		)
	}
}

func trimPackage(pkg string) string {
	if pkg == "" {
		return ""
	}
	parts := strings.Split(pkg, "/")
	return parts[len(parts)-1]
}
