// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shaderir

type Program struct {
	Uniforms      []Type
	Attributes    []Type
	Varyings      []Type
	Funcs         []Func
	VertexEntry   string
	FragmentEntry string

	structNames map[string]string
	structTypes []Type
}

type Func struct {
	Name        string
	InParams    []Type
	InOutParams []Type
	OutParams   []Type
	Block       Block
}

type Block struct {
	LocalVars []Type
	Stmts     []Stmt
}

type Stmt struct {
	Type    StmtType
	Exprs   []Expr
	Blocks  []Block
	ForInit Expr
	ForRest Expr
}

type StmtType int

const (
	ExprStmt StmtType = iota
	BlockStmt
	Assign
	If
	For
	Continue
	Break
	Return
	Discard
)

type Expr struct {
	Type     ExprType
	Exprs    []Expr
	Variable Variable
	Num      float64
	Ident    string
	Op       Op
}

type ExprType int

const (
	Numeric ExprType = iota
	VarName
	Ident
	Unary
	Binary
	Call
	Selector
	Index
)

type Variable struct {
	Type  VariableType
	Index int
}

type VariableType int

const (
	Uniform VariableType = iota
	Attribute
	Varying
	Local
)

type Op string

const (
	Add    Op = "+"
	Sub    Op = "-"
	Neg    Op = "!"
	Mul    Op = "*"
	Div    Op = "/"
	Mod    Op = "%"
	LShift Op = "<<"
	RShift Op = ">>"
	LT     Op = "<"
	LE     Op = "<="
	GT     Op = ">"
	GE     Op = ">="
	Eq     Op = "=="
	NE     Op = "!="
	And    Op = "&"
	Xor    Op = "^"
	Or     Op = "|"
	AndAnd Op = "&&"
	OrOr   Op = "||"
	Cond   Op = "?:"
)