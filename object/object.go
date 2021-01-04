package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ryan-berger/monkey/ast"
)

type Type string

const (
	INTEGER  = "INTEGER"
	BOOLEAN  = "BOOLEAN"
	NULL     = "NULL"
	RETURN   = "RETURN"
	ERROR    = "ERROR"
	FUNCTION = "FUNCTION"
	STRING   = "STRING"
)

type Object interface {
	Type() Type
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type      { return INTEGER }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type      { return BOOLEAN }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() Type      { return NULL }
func (n *Null) Inspect() string { return "null" }

type Return struct {
	Value Object
}

func (r *Return) Type() Type      { return RETURN }
func (r *Return) Inspect() string { return r.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() Type      { return ERROR }
func (e *Error) Inspect() string { return fmt.Sprintf("ERROR: %s", e.Message) }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type { return FUNCTION }

func (f *Function) Inspect() string {
	var out bytes.Buffer

	var params []string

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (s String) Type() Type      { return STRING }
func (s String) Inspect() string { return s.Value }
