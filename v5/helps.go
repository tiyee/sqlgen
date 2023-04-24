package v5

import (
	"fmt"
	"runtime"
	"strings"
)

func tildeWrapper(s string) string {
	return fmt.Sprintf("`%s`", s)
}
func findKeyword(s string) int {
	for k, v := range keywords {
		if strings.ToLower(s) == strings.ToLower(k) {
			return v
		}
	}
	for k, v := range ftm {
		if strings.ToLower(s) == strings.ToLower(k) {
			return v
		}
	}
	return 0
}
func unexpectedNode(x node, lexer yyLexer) {
	var buff [1024]byte
	runtime.Stack(buff[:], true)
	fmt.Println(string(buff[:]))
	lexer.Error(fmt.Sprintf("unexpected node: %d", x.Type()))
}

type PosArray[T IPos] []T

func (p PosArray[T]) Len() int {
	return len(p)
}

func (p PosArray[T]) Less(i, j int) bool {
	return p[i].Pos() < p[j].Pos()
}

func (p PosArray[T]) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
