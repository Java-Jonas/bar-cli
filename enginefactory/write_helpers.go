package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeDeduplicate() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {
			if !field.HasPointerValue {
				return
			}

			d := deduplicateWriter{
				f: field,
			}

			decls.File.Func().Id(d.name()).Params(d.params()).Id(d.returns()).Block(
				d.defineCheck(),
				d.defineDeduped(),
				For(d.loopConditions("a")).Block(
					d.checkValue(),
				),
				For(d.loopConditions("b")).Block(
					d.checkValue(),
				),
				d.loopCheck(),
				Return(Id("deduped")),
			)

		})
	})

	decls.Render(s.buf)
	return s
}

type deduplicateWriter struct {
	f ast.Field
}

func (d deduplicateWriter) idType() string {
	return title(d.f.Parent.Name) + title(pluralizeClient.Singular(d.f.Name)) + "RefID"
}

func (d deduplicateWriter) name() string {
	return "deduplicate" + d.idType() + "s"
}

func (d deduplicateWriter) params() *Statement {
	return List(Id("a").Id("[]"+d.idType()), Id("b").Id("[]"+d.idType()))
}

func (d deduplicateWriter) returns() string {
	return "[]" + d.idType()
}

func (d deduplicateWriter) defineCheck() *Statement {
	return Id("check").Op(":=").Make(Map(Id(d.idType())).Bool())
}

func (d deduplicateWriter) defineDeduped() *Statement {
	return Id("deduped").Op(":=").Make(Id(d.returns()), Lit(0))

}

func (d deduplicateWriter) loopConditions(getsLooped string) *Statement {
	return List(Id("_"), Id("val")).Op(":=").Range().Id(getsLooped)
}

func (d deduplicateWriter) checkValue() *Statement {
	return Id("check").Index(Id("val")).Op("=").Id("true")
}

func (d deduplicateWriter) loopCheck() *Statement {
	loop := For(Id("val").Op(":=").Range().Id("check")).Block(
		Id("deduped").Op("=").Append(Id("deduped"), Id("val")),
	)
	return loop
}

func (s *EngineFactory) writeAllIDsMethod() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {
			if !field.HasPointerValue {
				return
			}

			a := allIDsMehtodWriter{
				f: field,
			}

			decls.File.Func().Params(a.receiverParams()).Id(a.name()).Params().Id(a.returns()).Block(
				a.declareStateIDsSlice(),
				For(a.stateIDsLoopConditions()).Block(
					a.appendStateID(),
				),
				a.declarePatchIDsSlice(),
				For(a.patchIDsLoopConditions()).Block(
					a.appendPatchID(),
				),
				Return(a.deduplicatedIDs()),
			)
		})
	})

	decls.Render(s.buf)
	return s
}

type allIDsMehtodWriter struct {
	f ast.Field
}

func (a allIDsMehtodWriter) typeName() string {
	return title(a.f.Parent.Name) + title(pluralizeClient.Singular(a.f.Name)) + "Ref"
}

func (a allIDsMehtodWriter) idType() string {
	return a.typeName() + "ID"
}

func (a allIDsMehtodWriter) name() string {
	return "all" + a.idType() + "s"
}

func (a allIDsMehtodWriter) receiverParams() *Statement {
	return Id("engine").Id("Engine")
}

func (a allIDsMehtodWriter) returns() string {
	return "[]" + a.idType()
}

func (a allIDsMehtodWriter) idSliceName(prefix string) string {
	return prefix + a.idType() + "s"
}

func (a allIDsMehtodWriter) declareStateIDsSlice() *Statement {
	return Var().Id(a.idSliceName("state")).Id("[]" + a.idType())
}

func (a allIDsMehtodWriter) stateIDsLoopConditions() *Statement {
	return Id(lower(a.idType())).Op(":=").Range().Id("engine").Dot("State").Dot(a.typeName())
}

func (a allIDsMehtodWriter) appendStateID() *Statement {
	return Id(a.idSliceName("state")).Op("=").Append(Id(a.idSliceName("state")), Id(lower(a.idType())))
}

func (a allIDsMehtodWriter) declarePatchIDsSlice() *Statement {
	return Var().Id(a.idSliceName("patch")).Id("[]" + a.idType())
}

func (a allIDsMehtodWriter) patchIDsLoopConditions() *Statement {
	return Id(lower(a.idType())).Op(":=").Range().Id("engine").Dot("Patch").Dot(a.typeName())
}

func (a allIDsMehtodWriter) appendPatchID() *Statement {
	return Id(a.idSliceName("patch")).Op("=").Append(Id(a.idSliceName("patch")), Id(lower(a.idType())))
}

func (a allIDsMehtodWriter) deduplicatedIDs() *Statement {
	dedu := deduplicateWriter{a.f}
	return Id(dedu.name()).Call(Id(a.idSliceName("state")), Id(a.idSliceName("patch")))
}

func (s *EngineFactory) writeMergeIDs() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		m := mergeIDsWriter{
			idType: func() string {
				return title(configType.Name) + "ID"
			},
		}

		decls.File.Func().Id(m.name()).Params(m.params()).Id(m.returns()).Block(
			m.declareIDs(),
			m.copyIDs(),
			m.declareCounter(),
			For(m.currentIDsLoopConditions()).Block(
				If(m.idDoesNotMatch()).Block(
					Continue(),
				),
				m.incrementCounter(),
			),
			For(m.nextIDsLoopConditions()).Block(
				m.appendID(),
			),
			Return(Id("ids")),
		)

	})
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {
			if !field.HasPointerValue {
				return
			}

			m := mergeIDsWriter{
				idType: func() string {
					return title(field.Parent.Name) + title(pluralizeClient.Singular(field.Name)) + "RefID"
				},
			}

			decls.File.Func().Id(m.name()).Params(m.params()).Id(m.returns()).Block(
				m.declareIDs(),
				m.copyIDs(),
				m.declareCounter(),
				For(m.currentIDsLoopConditions()).Block(
					If(m.idDoesNotMatch()).Block(
						Continue(),
					),
					m.incrementCounter(),
				),
				For(m.nextIDsLoopConditions()).Block(
					m.appendID(),
				),
				Return(Id("ids")),
			)

		})
	})

	decls.Render(s.buf)
	return s
}

type mergeIDsWriter struct {
	idType func() string
}

func (m mergeIDsWriter) name() string {
	return "merge" + m.idType() + "s"
}

func (m mergeIDsWriter) params() *Statement {
	return List(Id("currentIDs, nextIDs").Id("[]" + m.idType()))
}

func (m mergeIDsWriter) returns() string {
	return "[]" + m.idType()
}

func (m mergeIDsWriter) declareIDs() *Statement {
	return Id("ids").Op(":=").Make(Id("[]"+m.idType()), Len(Id("currentIDs")))
}

func (m mergeIDsWriter) copyIDs() *Statement {
	return Copy(Id("ids"), Id("currentIDs"))
}

func (m mergeIDsWriter) declareCounter() *Statement {
	return Var().Id("j").Int()
}

func (m mergeIDsWriter) currentIDsLoopConditions() *Statement {
	return List(Id("_"), Id("currentID")).Op(":=").Range().Id("currentIDs")
}

func (m mergeIDsWriter) idDoesNotMatch() *Statement {
	return Len(Id("nextIDs")).Op("<=").Id("j").Op("||").Id("currentID").Op("!=").Id("nextIDs").Index(Id("j"))
}

func (m mergeIDsWriter) incrementCounter() *Statement {
	return Id("j").Op("+=").Lit(1)
}

func (m mergeIDsWriter) nextIDsLoopConditions() *Statement {
	return List(Id("_"), Id("nextID")).Op(":=").Range().Id("nextIDs").Index(Id("j:"))
}

func (m mergeIDsWriter) appendID() *Statement {
	return Id("ids").Op("=").Append(Id("ids"), Id("nextID"))
}
