package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writePathTrack() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("pathTrack").Struct(
		Id("_iterations").Int(),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return Id(configType.Name).Map(Id(title(configType.Name) + "ID")).Id("path")
		}),
	)

	decls.File.Func().Id("newPathTrack").Params().Id("pathTrack").Block(
		Return(Id("pathTrack").Block(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				return Id(configType.Name).Id(":").Make(Map(Id(title(configType.Name) + "ID")).Id("path")).Id(",")
			}),
		)),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeIdentifiers() *EngineFactory {
	decls := NewDeclSet()

	alreadyWrittenCheck := make(map[string]bool)
	identifierValue := 0
	decls.File.Const().Defs(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			typeIdentifierShouldBeWritten := alreadyWrittenCheck[configType.Name]
			alreadyWrittenCheck[configType.Name] = true
			if !typeIdentifierShouldBeWritten {
				identifierValue -= 1
			}
			return &Statement{
				onlyIf(!typeIdentifierShouldBeWritten, Id(configType.Name+"Identifier").Int().Op("=").Lit(identifierValue).Line()),
				ForEachFieldInType(configType, func(field ast.Field) *Statement {
					if alreadyWrittenCheck[field.Name] {
						return Empty()
					}
					if field.ValueType().IsBasicType || field.HasPointerValue {
						return Empty()
					}
					alreadyWrittenCheck[field.Name] = true
					identifierValue -= 1
					return Id(field.Name + "Identifier").Int().Op("=").Lit(identifierValue)
				}),
			}
		}),
	)

	decls.Render(s.buf)
	return s
}

func writePathSegmentMethod(decls DeclSet, name string) {
	decls.File.Func().Params(Id("p").Id("path")).Id(name).Params().Id("path").Block(
		Id("newPath").Op(":=").Make(Id("[]int"), Len(Id("p")), Len(Id("p")).Op("+").Lit(1)),
		Copy(Id("newPath"), Id("p")),
		Id("newPath").Op("=").Append(Id("newPath"), Id(name+"Identifier")),
		Return(Id("newPath")),
	)
}

func (s *EngineFactory) writePathSegments() *EngineFactory {
	decls := NewDeclSet()

	alreadyWrittenCheck := make(map[string]bool)
	s.config.RangeTypes(func(configType ast.ConfigType) {

		if !alreadyWrittenCheck[configType.Name] {
			writePathSegmentMethod(decls, configType.Name)
			alreadyWrittenCheck[configType.Name] = true
		}

		configType.RangeFields(func(field ast.Field) {
			if alreadyWrittenCheck[field.Name] {
				return
			}
			if field.ValueType().IsBasicType || field.HasPointerValue {
				return
			}
			alreadyWrittenCheck[field.Name] = true
			writePathSegmentMethod(decls, field.Name)
		})

	})

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writePath() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("path").Id("[]int")

	decls.File.Func().Id("newPath").Params(Id("elementIdentifier"), Id("id").Int()).Id("path").Block(
		Return(Id("[]int").Values(Id("elementIdentifier"), Id("id"))),
	)

	decls.File.Func().Params(Id("p").Id("path")).Id("index").Params(Id("i").Id("int")).Id("path").Block(
		Id("newPath").Op(":=").Make(Id("[]int"), Len(Id("p")), Len(Id("p")).Op("+").Lit(1)),
		Copy(Id("newPath"), Id("p")),
		Id("newPath").Op("=").Append(Id("newPath"), Id("i")),
		Return(Id("newPath")),
	)

	decls.File.Func().Params(Id("p").Id("path")).Id("equals").Params(Id("parentPath").Id("path")).Bool().Block(
		If(Len(Id("p")).Op("!=").Len(Id("parentPath"))).Block(
			Return(False()),
		),
		For(Id("i"), Id("segment").Op(":=").Range().Id("parentPath")).Block(
			If(Id("segment").Op("!=").Id("p").Index(Id("i"))).Block(
				Return(False()),
			),
		),
		Return(True()),
	)

	decls.Render(s.buf)
	return s
}