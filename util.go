package unexport

import (
	"fmt"
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/types"
	"unicode"
	"unicode/utf8"
)

func wholePath(obj types.Object, pkg *types.Package, prog *loader.Program) string {
	if v, ok := obj.(*types.Var); ok && v.IsField() {
		structName := getDeclareStructOrInterface(prog, v)
		return fmt.Sprintf("(%s.%s).%s", pkg.Path(), structName, obj.Name())
	} else if f, ok := obj.(*types.Func); ok {
		if r := recv(f); r != nil {
			structName := r.Type().String()
			return fmt.Sprintf("(%s.%s).%s", pkg.Path(), structName, obj.Name())
		}
	}
	return fmt.Sprintf("%s.%s", pkg.Path(), obj.Name())
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// recv returns the method's receiver.
func recv(meth *types.Func) *types.Var {
	return meth.Type().(*types.Signature).Recv()
}
