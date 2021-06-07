package factoryutils

import (
	"bar-cli/ast"
	"bytes"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
)

// pluralizeClient.Singular is used to find the singular of field names
// this is necessary for writing coherent method names, eg. in write_adders.go (toSingular)
// with getting the singular form of a plural, this field:
// { pieces []piece }
// can have the coherent adder method of "AddPiece"
var Singular func(string) string = pluralize.NewClient().Singular

func ForEachParamInAction(action ast.Action, fn func(param ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	action.RangeParams(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachFieldInType(configType ast.ConfigType, fn func(field ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	configType.RangeFields(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachValueOfField(field ast.Field, fn func(configType *ast.ConfigType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	field.RangeValueTypes(func(configType *ast.ConfigType) {
		statements = append(statements, fn(configType))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachReferenceOfType(configType ast.ConfigType, fn func(field *ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	configType.RangeReferencedBy(func(field *ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachTypeInAST(config *ast.AST, fn func(configType ast.ConfigType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeTypes(func(configType ast.ConfigType) {
		statements = append(statements, fn(configType))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachRefFieldInAST(config *ast.AST, fn func(field ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeRefFields(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachAnyFieldInAST(config *ast.AST, fn func(field ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeAnyFields(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachActionInAST(config *ast.AST, fn func(action ast.Action) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeActions(func(action ast.Action) {
		statements = append(statements, fn(action))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachFieldValueComparison(field ast.Field, comparator jen.Statement, fn func(configType *ast.ConfigType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	first := true
	field.RangeValueTypes(func(valueType *ast.ConfigType) {
		statement := jen.Empty()
		if !first {
			statement.Else()
		}
		_comparator := comparator
		statement.If(_comparator.Op("==").Id("ElementKind" + Title(valueType.Name))).Block(
			fn(valueType),
		)
		statements = append(statements, statement)
		first = false
	})
	return &statements
}

func Title(name string) string {
	return strings.Title(name)
}

func Lower(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func OnlyIf(is bool, statement *jen.Statement) *jen.Statement {
	if is {
		return statement
	}
	return jen.Empty()
}

type DeclSet struct {
	File *jen.File
}

func NewDeclSet() DeclSet {
	return DeclSet{
		File: jen.NewFile("main"),
	}
}

func (d DeclSet) Render(buf *bytes.Buffer) {
	var _buf bytes.Buffer
	err := d.File.Render(&_buf)
	if err != nil {
		panic(err)
	}
	code := strings.TrimPrefix(_buf.String(), "package main")
	code = strings.TrimSpace(code)
	buf.WriteString("\n" + code + "\n")
}
