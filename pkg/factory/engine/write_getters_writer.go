package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

type existsGetterWriter struct {
	t ast.ConfigType
}

func (e existsGetterWriter) receiverName() string {
	return "_" + e.t.Name
}

func (e existsGetterWriter) receiverParams() *Statement {
	return Id(e.receiverName()).Id(Title(e.t.Name))
}

func (e existsGetterWriter) returnTypes() (*Statement, *Statement) {
	return Id(Title(e.t.Name)), Bool()
}

func (e existsGetterWriter) isNotOperationKindDelete() *Statement {
	return Id(e.t.Name).Dot(e.t.Name).Dot("OperationKind").Op("!=").Id("OperationKindDelete")
}

func (e existsGetterWriter) reassignElement() *Statement {
	return Id(e.t.Name).Op(":=").Id(e.receiverName()).Dot(e.t.Name).Dot("engine").Dot(Title(e.t.Name)).Call(Id(e.receiverName()).Dot(e.t.Name).Dot("ID"))
}

type everyTypeGetterWriter struct {
	t ast.ConfigType
}

func (t everyTypeGetterWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (t everyTypeGetterWriter) returns() *Statement {
	return Id("[]" + Title(t.t.Name))
}

func (t everyTypeGetterWriter) allIDs() *Statement {
	return Id(t.t.Name + "IDs").Op(":=").Id("engine").Dot("all" + Title(t.t.Name) + "IDs").Call()
}

func (t everyTypeGetterWriter) sortIDs() *Statement {
	return Id("sort").Dot("Slice").Params(Id(t.t.Name+"IDs"), Func().Params(Id("i"), Id("j").Int()).Bool().Block(
		Return(Id(t.t.Name+"IDs").Index(Id("i")).Op("<").Id(t.t.Name+"IDs").Index(Id("j"))),
	))
}

func (t everyTypeGetterWriter) sliceName() string {
	return t.t.Name + "s"
}

func (t everyTypeGetterWriter) declareSlice() *Statement {
	return Var().Id(t.sliceName()).Id("[]" + Title(t.t.Name))
}

func (t everyTypeGetterWriter) loopConditions() *Statement {
	return List(Id("_"), Id(t.t.Name+"ID")).Op(":=").Range().Id(t.t.Name + "IDs")
}

func (t everyTypeGetterWriter) declareElement() *Statement {
	return Id(t.t.Name).Op(":=").Id("engine").Dot(Title(t.t.Name)).Call(Id(t.t.Name + "ID"))
}

func (t everyTypeGetterWriter) elementHasParent() *Statement {
	return Id(t.t.Name).Dot(t.t.Name).Dot("HasParent")
}

func (t everyTypeGetterWriter) returnToIdsSliceToPool() *Statement {
	return Id(t.t.Name + "IDSlicePool").Dot("Put").Call(Id(t.t.Name + "IDs"))
}

func (t everyTypeGetterWriter) appendElement() *Statement {
	return Id(t.sliceName()).Op("=").Append(Id(t.sliceName()), Id(t.t.Name))
}

type typeGetterWriter struct {
	name     func() string
	typeName func() string
}

func (t typeGetterWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (t typeGetterWriter) idParam() string {
	return t.typeName() + "ID"
}

func (t typeGetterWriter) params() *Statement {
	return Id(t.idParam()).Id(Title(t.typeName()) + "ID")
}

func (t typeGetterWriter) returns() string {
	return Title(t.typeName())
}

func (t typeGetterWriter) definePatchingElement() *Statement {
	return List(Id("patching"+Title(t.typeName())), Id("ok")).Op(":=").Id("engine").Dot("Patch").Dot(Title(t.typeName())).Index(Id(t.idParam()))
}

func (t typeGetterWriter) earlyReturnPatching() *Statement {
	return Id(Title(t.typeName())).Values(Dict{Id(t.typeName()): Id("patching" + Title(t.typeName()))})
}

func (t typeGetterWriter) defineCurrentElement() *Statement {
	return List(Id("current"+Title(t.typeName())), Id("ok")).Op(":=").Id("engine").Dot("State").Dot(Title(t.typeName())).Index(Id(t.idParam()))
}

func (t typeGetterWriter) earlyReturnCurrent() *Statement {
	return Id(Title(t.typeName())).Values(Dict{Id(t.typeName()): Id("current" + Title(t.typeName()))})
}

func (t typeGetterWriter) finalReturn() *Statement {
	return Id(Title(t.typeName())).Values(Dict{Id(t.typeName()): Id(t.typeName() + "Core").Values(Dict{Id("OperationKind"): Id("OperationKindDelete"), Id("engine"): Id("engine")})})
}

type idGetterWriter struct {
	typeName        func() string
	returns         func() string
	idFieldToReturn func() string
}

func (i idGetterWriter) receiverName() string {
	return "_" + i.typeName()
}

func (i idGetterWriter) receiverParams() *Statement {
	return Id(i.receiverName()).Id(Title(i.typeName()))
}

func (i idGetterWriter) name() string {
	return "ID"
}

func (i idGetterWriter) returnID() *Statement {
	return Id(i.receiverName()).Dot(i.typeName()).Dot(i.idFieldToReturn())
}

type fieldGetterWriter struct {
	t ast.ConfigType
	f ast.Field
}

func (f fieldGetterWriter) receiverName() string {
	return "_" + f.t.Name
}

func (f fieldGetterWriter) receiverParams() *Statement {
	return Id(f.receiverName()).Id(Title(f.t.Name))
}

func (f fieldGetterWriter) name() string {
	return Title(f.f.Name)
}

func (f fieldGetterWriter) returnedType() string {

	if f.f.ValueType().IsBasicType {
		return f.f.ValueType().Name
	}
	if f.f.HasPointerValue {
		return Title(f.f.Parent.Name) + Title(Singular(f.f.Name)) + "Ref"
	}
	if f.f.HasAnyValue {
		return Title(anyNameByField(f.f))
	}
	return Title(f.f.ValueType().Name)
}

func (f fieldGetterWriter) returns() string {
	returnedLiteral := f.returnedType()
	if f.f.HasSliceValue {
		return "[]" + returnedLiteral
	}
	return returnedLiteral
}

func (f fieldGetterWriter) reassignElement() *Statement {
	return Id(f.t.Name).Op(":=").Id(f.receiverName()).Dot(f.t.Name).Dot("engine").Dot(Title(f.t.Name)).Call(Id(f.receiverName()).Dot(f.t.Name).Dot("ID"))
}

func (f fieldGetterWriter) declareSliceOfElements() *Statement {
	returnedType := f.returnedType()
	if f.f.HasSliceValue {
		returnedType = "[]" + returnedType
	}
	return Var().Id(f.f.Name).Id(returnedType)
}

func (f fieldGetterWriter) loopedElementIdentifier() string {
	if f.f.ValueType().IsBasicType {
		return "element"
	}
	if f.f.HasPointerValue {
		return "refID"
	}
	if f.f.HasAnyValue {
		return anyNameByField(f.f) + "ID"
	}
	return f.f.ValueType().Name + "ID"
}

func (f fieldGetterWriter) loopConditions() *Statement {
	identifier := f.loopedElementIdentifier()
	return List(Id("_"), Id(identifier)).Op(":=").Range().Id(f.t.Name).Dot(f.t.Name).Dot(Title(f.f.Name))
}

func (f fieldGetterWriter) appendedItem() *Statement {
	if f.f.ValueType().IsBasicType {
		return Id(f.loopedElementIdentifier())
	}
	returnedType := Lower(f.returnedType())
	if !f.f.HasPointerValue && !f.f.HasAnyValue {
		returnedType = Title(returnedType)
	}
	return Id(f.t.Name).Dot(f.t.Name).Dot("engine").Dot(returnedType).Call(Id(f.loopedElementIdentifier()))
}

func (f fieldGetterWriter) appendElement() *Statement {
	return Id(f.f.Name).Op("=").Append(Id(f.f.Name), f.appendedItem())
}

func (f fieldGetterWriter) returnSliceOfType() *Statement {
	return Id(f.f.Name)
}

func (f fieldGetterWriter) returnBasicType() *Statement {
	return Id(f.t.Name).Dot(f.t.Name).Dot(Title(f.f.Name))
}

func (f fieldGetterWriter) returnNamedType() *Statement {
	engine := Id(f.t.Name).Dot(f.t.Name).Dot("engine")
	returnedType := Lower(f.returnedType())
	if !f.f.HasAnyValue && !f.f.HasPointerValue {
		returnedType = Title(returnedType)
	}
	return engine.Dot(returnedType).Call(Id(f.t.Name).Dot(f.t.Name).Dot(Title(f.f.Name)))
}

func (f fieldGetterWriter) returnSingleType() *Statement {
	if f.f.ValueType().IsBasicType {
		return f.returnBasicType()
	}
	return f.returnNamedType()
}

type pathGetterWriter struct {
	t ast.ConfigType
}

func (p pathGetterWriter) receiverName() string {
	return "_" + p.t.Name
}

func (p pathGetterWriter) receiverParams() *Statement {
	return Id(p.receiverName()).Id(Title(p.t.Name))
}

func (p pathGetterWriter) returnPath() *Statement {
	return Id(p.receiverName()).Dot(p.t.Name).Dot("Path")
}