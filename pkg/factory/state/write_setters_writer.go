package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type setterWriter struct {
	t ast.ConfigType
	f ast.Field
}

func (s setterWriter) receiverParams() *Statement {
	return Id(s.receiverName()).Id(Title(s.t.Name))
}

func (s setterWriter) name() string {
	if s.f.ValueType().IsBasicType {
		return "Set" + Title(s.f.Name)
	}
	return "Add" + Title(Singular(s.f.Name))
}

func (s setterWriter) newValueParam() string {
	return "new" + Title(s.f.Name)
}

func (s setterWriter) params() *Statement {
	return Id(s.newValueParam()).Id(s.f.ValueTypeName)
}

func (s setterWriter) returns() string {
	return Title(s.t.Name)
}

func (s setterWriter) reassignElement() *Statement {
	return Id(s.t.Name).Op(":=").Id(s.receiverName()).Dot(s.t.Name).Dot("engine").Dot(Title(s.t.Name)).Call(Id(s.receiverName()).Dot(s.t.Name).Dot("ID"))
}

func (s setterWriter) isOperationKindDelete() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (s setterWriter) setAttribute() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("engine").Dot("set"+Title(BasicTypes[s.f.ValueTypeName])).Call(Id(s.t.Name).Dot(s.t.Name).Dot(Title(s.f.Name)), Id(s.newValueParam()))
}

func (s setterWriter) setOperationKind() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (s setterWriter) setOperationKindUpdate() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (s setterWriter) updateElementInPatch() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("engine").Dot("Patch").Dot(Title(s.t.Name)).Index(s.elementID()).Op("=").Id(s.t.Name).Dot(s.t.Name)
}

func (s setterWriter) elementID() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("ID")
}

func (s setterWriter) receiverName() string {
	return "_" + s.t.Name
}

type setRefFieldWeiter struct {
	f ast.Field
	v *ast.ConfigType
}

func (s setRefFieldWeiter) receiverParams() *Statement {
	return Id("_" + s.f.Parent.Name).Id(Title(s.f.Parent.Name))
}

func (s setRefFieldWeiter) name() string {
	var optionalSuffix string
	if s.f.HasAnyValue {
		optionalSuffix = Title(s.v.Name)
	}
	return "Set" + Title(s.f.Name) + optionalSuffix
}

func (s setRefFieldWeiter) idParam() string {
	return s.v.Name + "ID"
}

func (s setRefFieldWeiter) params() *Statement {
	return Id(s.idParam()).Id(Title(s.v.Name) + "ID")
}

func (s setRefFieldWeiter) returns() string {
	return Title(s.f.Parent.Name)
}

func (s setRefFieldWeiter) reassignElement() *Statement {
	return Id(s.f.Parent.Name).Op(":=").Id("_" + s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot(Title(s.f.Parent.Name)).Call(Id("_" + s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("ID"))
}

func (s setRefFieldWeiter) isOperationKindDelete() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (s setRefFieldWeiter) isReferencedElementDeleted() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot(Title(s.v.Name)).Call(Id(s.idParam())).Dot(s.v.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (s setRefFieldWeiter) isRefAlreadyAssigned() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot(Title(s.f.Name)).Op("!=").Params(Id(Title(s.f.ValueTypeName) + "ID").Values())
}

func (s setRefFieldWeiter) isSameID() *Statement {
	return Id(Title(s.v.Name) + "ID").Call(Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot(Title(s.f.Name)).Dot("ChildID")).Op("==").Add(Id(s.idParam()))
}

func (s setRefFieldWeiter) deleteExistingRef() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot("delete" + Title(s.f.ValueTypeName)).Call(Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot(Title(s.f.Name)))
}

func (s setRefFieldWeiter) createAnyContainerCall() *Statement {
	return Call(Int().Call(Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("ID")), Int().Call(Id(s.idParam())), Id("ElementKind"+Title(s.v.Name)), Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("Path"), Id(FieldPathIdentifier(s.f)))
}

func (s setRefFieldWeiter) createAnyContainer() *Statement {
	return Id("anyContainer").Op(":=").Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot("create" + Title(anyNameByField(s.f))).Add(s.createAnyContainerCall())
}

func (s setRefFieldWeiter) referenceID() *Statement {
	switch {
	case s.f.HasAnyValue:
		return Id("anyContainer").Dot(anyNameByField(s.f)).Dot("ID")
	default:
		return Id(s.v.Name + "ID")
	}
}

func (s setRefFieldWeiter) createNewRefCall() *Statement {
	return Call(
		Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("Path"),
		Id(FieldPathIdentifier(s.f)),
		(s.referenceID()),
		Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("ID"),
		OnlyIf(s.f.HasAnyValue, List(
			Id("ElementKind"+Title(s.v.Name)),
			Int().Call(Id(s.idParam())),
		)),
	)
}

func (s setRefFieldWeiter) createNewRef() *Statement {
	return Id("ref").Op(":=").Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot("create" + Title(s.f.ValueTypeName)).Add(s.createNewRefCall())
}

func (s setRefFieldWeiter) setNewRef() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot(Title(s.f.Name)).Op("=").Id("ref").Dot("ID")
}

func (s setRefFieldWeiter) setOperationKind() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (s setRefFieldWeiter) signElement() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("Meta").Dot("sign").Call(Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot("broadcastingClientID"))
}

func (s setRefFieldWeiter) setItemInPatch() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot("Patch").Dot(Title(s.f.Parent.Name)).Index(Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("ID")).Op("=").Id(s.f.Parent.Name).Dot(s.f.Parent.Name)
}