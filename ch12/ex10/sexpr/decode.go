package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() // get the first token
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names.
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		v.SetInt(int64(i))
		lex.next()
		return
	case scanner.Float:
		switch v.Kind() {
		case reflect.Float32:
			f, _ := strconv.ParseFloat(lex.text(), 32) // NOTE: ignoring erros
			v.SetFloat(f)
		case reflect.Float64:
			f, _ := strconv.ParseFloat(lex.text(), 64) // NOTE: ignoring erros
			v.SetFloat(f)
		default:
			panic(fmt.Sprintf("unexpected type: %d", v.Kind()))
		}
		lex.next()
		return
	case '#':
		lex.next()
		lex.next()
		lex.next()
		r := lex.text()
		lex.next()
		i := lex.text()
		lex.next()
		lex.next() // consume ')'

		var bitSize int
		switch v.Kind() {
		case reflect.Complex64:
			bitSize = 32
		case reflect.Complex128:
			bitSize = 64
		default:
			panic(fmt.Sprintf("unexpected type: %d", v.Kind()))
		}
		fr, _ := strconv.ParseFloat(r, bitSize)
		fi, _ := strconv.ParseFloat(i, bitSize)
		v.SetComplex(complex(fr, fi))
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // consume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	case reflect.Interface:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		item := reflect.New(typeOf(s)).Elem()
		lex.next()
		read(lex, item)
		v.Set(item)

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}
func typeOf(name string) reflect.Type {
	t, ok := typeRegistered[name]
	if ok {
		return t
	}

	if strings.HasPrefix(name, "[]") {
		return reflect.SliceOf(typeOf(name[2:]))
	}

	if name[0] == '[' {
		i := strings.Index(name, "]")
		if i > 0 {
			size, _ := strconv.Atoi(name[1:i]) // NOTE: ignoring errors
			return reflect.ArrayOf(size, typeOf(name[i+1:]))
		}
	}

	if strings.HasPrefix(name, "map") {
		i := strings.Index(name, "]")
		if i > 0 {
			return reflect.MapOf(typeOf(name[4:i]), typeOf(name[i+1:]))
		}

	}

	panic(fmt.Sprintf("%s not supported\n", name))

}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
