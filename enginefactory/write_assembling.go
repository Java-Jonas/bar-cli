package enginefactory

import (
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"

	"bar-cli/ast"
)

func (s *EngineFactory) writeAssembleTree() *EngineFactory {
	decls := NewDeclSet()

	a := assembleTreeWriter{}

	decls.File.Func().Params(a.receiverParams()).Id("assembleTree").Params().Id("Tree").Block(
		a.createConfig(),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			a.t = &configType

			if a.t.IsRootType {
				return &Statement{
					For(a.patchLoopConditions()).Block(
						a.assembleElement(),
						If(Id("include")).Block(
							a.setElementInTree(),
						),
					),
				}
			}

			return &Statement{
				For(a.patchLoopConditions()).Block(
					If(a.elementHasNoParent()).Block(
						a.assembleElement(),
						If(Id("include")).Block(
							a.setElementInTree(),
						),
					),
				),
			}

		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			a.t = &configType

			if a.t.IsRootType {
				return &Statement{
					For(a.stateLoopConditions()).Block(
						If(a.elementNonExistentInTree()).Block(
							a.assembleElement(),
							If(Id("include")).Block(
								a.setElementInTree(),
							),
						),
					),
				}
			}

			return &Statement{
				For(a.stateLoopConditions()).Block(
					If(a.elementHasNoParent()).Block(
						If(a.elementNonExistentInTree()).Block(
							a.assembleElement(),
							If(Id("include")).Block(
								a.setElementInTree(),
							),
						),
					),
				),
			}

		}),
		Return(Id("engine").Dot("Tree")),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeAssembleTreeElement() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {

		a := assembleElement{
			t: configType,
			f: nil,
		}

		decls.File.Func().Params(a.receiverParams()).Id(a.name()).Params(a.params()).Params(a.returns()).Block(
			If(a.checkIsDefined()).Block(
				If(a.elementExistsInCheck()).Block(
					Return(a.returnEmpty()),
				).Else().Block(
					a.checkElement(),
				),
			),
			a.getElementFromPatch(),
			If(Id("!hasUpdated")).Block(
				a.getElementFromState(),
			),
			a.declareTreeElement(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				a.f = &field

				if a.f.ValueType().IsBasicType {
					return Empty()
				}

				if field.HasSliceValue {
					if field.HasAnyValue && !field.HasPointerValue {
						return For(a.sliceFieldLoopConditions()).Block(
							onlyIf(!field.HasPointerValue, a.createAnyContainer().Line()),
							forEachFieldValueComparison(field, *Id(a.anyContainerName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
								return &Statement{
									Id(valueType.Name + "ID").Op(":=").Id(a.anyContainerName()).Dot(title(valueType.Name)).Line(),
									If(a.elementHasUpdated(valueType, a.usedAssembleID(configType, field, valueType))).Block(
										If(Id("childHasUpdated")).Block(
											a.setHasUpdatedTrue(),
										),
										a.appendToElementsInField(valueType),
									),
								}
							}),
						)
					} else {
						return For(a.sliceFieldLoopConditions()).Block(
							If(a.elementHasUpdated(field.ValueType(), a.usedAssembleID(configType, field, field.ValueType()))).Block(
								If(Id("childHasUpdated")).Block(
									a.setHasUpdatedTrue(),
								),
								a.appendToElementsInField(field.ValueType()),
							),
						)
					}
				}

				if field.HasAnyValue && !field.HasPointerValue {
					return &Statement{
						onlyIf(!field.HasPointerValue, a.createAnyContainer().Line()),
						forEachFieldValueComparison(field, *Id(a.anyContainerName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
							return &Statement{
								Id(valueType.Name + "ID").Op(":=").Id(a.anyContainerName()).Dot(title(valueType.Name)).Line(),
								If(a.elementHasUpdated(valueType, a.usedAssembleID(configType, field, valueType))).Block(
									If(Id("childHasUpdated")).Block(
										a.setHasUpdatedTrue(),
									),
									a.setFieldElement(valueType),
								),
							}

						}),
					}
				}
				return If(a.elementHasUpdated(field.ValueType(), a.usedAssembleID(configType, field, field.ValueType()))).Block(
					If(Id("childHasUpdated")).Block(
						a.setHasUpdatedTrue(),
					),
					a.setFieldElement(field.ValueType()),
				)
			}),
			a.setID(),
			a.setOperationKind(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				a.f = &field

				if !a.f.ValueType().IsBasicType {
					return Empty()
				}

				return a.setField()
			}),
			Return(a.finalReturn()),
		)
	})

	decls.Render(s.buf)
	return s

}
