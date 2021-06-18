package enginefactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeWalkElement() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {

		w := walkElementWriter{
			t: configType,
		}

		if configType.IsLeafType {
			decls.File.Func().Params(w.receiverParams()).Id(w.name()).Params(w.params()).Block(
				w.updatePath(),
			)
			return
		}

		decls.File.Func().Params(w.receiverParams()).Id(w.name()).Params(w.params()).Block(
			w.getElementFromPatch(),
			If(Id("!hasUpdated")).Block(
				w.getElementFromState(),
			),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				if field.ValueType().IsBasicType || field.HasPointerValue {
					return Empty()
				}

				w.f = &field
				w.v = field.ValueType()

				if !field.HasAnyValue {
					if !field.HasSliceValue {
						return writeEvaluateChildPath(w)
					}
					return For(w.childrenLoopConditions()).Block(
						writeEvaluateChildPath(w),
					)
				}

				if !field.HasSliceValue {
					return &Statement{
						w.declareAnyContainer().Line(),
						ForEachFieldValueComparison(field, *Id(field.Name + "Container").Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
							w.v = valueType
							return writeEvaluateChildPath(w)

						}),
					}
				}
				return For(w.anyChildLoopConditions()).Block(
					w.declareAnyContainer(),
					ForEachFieldValueComparison(field, *Id(field.Name + "Container").Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
						w.v = valueType
						return writeEvaluateChildPath(w)

					}),
				)
			}),
			w.updatePath(),
		)
	})

	decls.Render(s.buf)
	return s
}

func writeEvaluateChildPath(w walkElementWriter) *Statement {
	return &Statement{
		w.declarePathVar().Line(),
		If(w.getChildPath(), w.pathNeedsUpdate()).Block(
			w.setChildPathNew(),
		).Else().Block(
			w.setChildPathExisting(),
		).Line(),
		w.walkChild(),
	}
}

func (s *EngineFactory) writeWalkTree() *EngineFactory {
	decls := NewDeclSet()

	w := walkTreeWriter{
		t: nil,
	}

	decls.File.Func().Params(Id("engine").Id("*Engine")).Id("walkTree").Params().Block(
		Id("walkedCheck").Op(":=").Id("newRecursionCheck").Call(),

		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			w.t = &configType

			if configType.IsRootType {
				return For(w.patchLoopConditions()).Block(
					w.walkElement(),
					w.checkWalked(),
				)
			}
			return For(w.patchLoopConditions()).Block(
				If(w.elementDoesNotHaveParent()).Block(
					w.walkElement(),
					w.checkWalked(),
				),
			)
		}),

		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			w.t = &configType
			if configType.IsRootType {
				return For(w.stateLoopConditions()).Block(
					If(w.hasNotBeenWalked()).Block(
						w.walkElement(),
					),
				)
			}
			return For(w.stateLoopConditions()).Block(
				If(w.elementDoesNotHaveParent()).Block(
					If(w.hasNotBeenWalked()).Block(
						w.walkElement(),
					),
				),
			)
		}),

		Id("engine").Dot("PathTrack").Dot("_iterations").Op("+=").Lit(1),
		If(Id("engine").Dot("PathTrack").Dot("_iterations").Op("==").Lit(100)).Block(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				w.t = &configType
				return w.clearPathTrack()
			}),
		),
	)

	decls.Render(s.buf)
	return s
}
