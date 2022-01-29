package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeOperationKind() *EngineFactory {

	s.file.Type().Id("OperationKind").String()

	s.file.Const().Defs(
		Id("OperationKindDelete").Id("OperationKind").Op("=").Lit("DELETE"),
		Id("OperationKindUpdate").Id("OperationKind").Op("=").Lit("UPDATE"),
		Id("OperationKindUnchanged").Id("OperationKind").Op("=").Lit("UNCHANGED"),
	)

	return s
}

func (s *EngineFactory) writeEngine() *EngineFactory {

	s.file.Type().Id("Engine").Struct(
		Id("State").Id("*State"),
		Id("Patch").Id("*State"),
		Id("Tree").Id("*Tree"),
		Id("planner").Id("*assemblePlanner"),
		Id("IDgen").Int(),
	)

	s.file.Func().Id("newEngine").Params().Id("*Engine").Block(
		Return(Id("&Engine").Values(Dict{
			Id("State"):   Id("newState").Call(),
			Id("Patch"):   Id("newState").Call(),
			Id("Tree"):    Id("newTree").Call(),
			Id("planner"): Id("newAssemblePlanner").Call(),
			Id("IDgen"):   Lit(1),
		})),
	)

	return s
}

func (s *EngineFactory) writeGenerateID() *EngineFactory {

	s.file.Func().Params(Id("engine").Id("*Engine")).Id("GenerateID").Params().Int().Block(
		Id("newID").Op(":=").Id("engine").Dot("IDgen"),
		Id("engine").Dot("IDgen").Op("=").Id("engine").Dot("IDgen").Op("+").Lit(1),
		Return(Id("newID")),
	)

	return s
}

func (s *EngineFactory) writeUpdateState() *EngineFactory {

	u := updateStateWriter{}

	s.file.Func().Params(u.receiverParams()).Id("UpdateState").Params().Block(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			u.typeName = func() string {
				return configType.Name
			}
			return writeUpdateElement(u)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return field.ValueTypeName
			}
			return writeUpdateElement(u)
		}),
		ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return anyNameByField(field)
			}
			return writeUpdateElement(u)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			u.typeName = func() string {
				return configType.Name
			}
			return writeClearPatch(u)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return field.ValueTypeName
			}
			return writeClearPatch(u)
		}),
		ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return anyNameByField(field)
			}
			return writeClearPatch(u)
		}),
	)

	return s
}

func writeUpdateElement(u updateStateWriter) *Statement {
	return For(u.loopPatchElementsConditions()).Block(
		If(u.isOperationKindDelete()).Block(
			u.deleteElement(),
		).Else().Block(
			u.setOperationKindUnchanged(),
			u.updateElement(),
		),
	)
}

func writeClearPatch(u updateStateWriter) *Statement {
	return For(u.loopPatchKeysConditions()).Block(
		u.clearElementFromPatch(),
	)
}