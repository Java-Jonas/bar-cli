package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

type deleteTypeWrapperWriter struct {
	t ast.ConfigType
}

func (d deleteTypeWrapperWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (d deleteTypeWrapperWriter) name() string {
	return "Delete" + title(d.t.Name)
}

func (d deleteTypeWrapperWriter) idParam() string {
	return d.t.Name + "ID"
}

func (d deleteTypeWrapperWriter) params() *Statement {
	return Id(d.idParam()).Id(title(d.t.Name) + "ID")
}

func (d deleteTypeWrapperWriter) getElement() *Statement {
	return Id(d.t.Name).Op(":=").Id("engine").Dot(title(d.t.Name)).Call(Id(d.idParam())).Dot(d.t.Name)
}

func (d deleteTypeWrapperWriter) hasParent() *Statement {
	return Id(d.t.Name).Dot("HasParent")
}

func (d deleteTypeWrapperWriter) deleteElement() *Statement {
	return Id("engine").Dot("delete" + title(d.t.Name)).Call(Id(d.idParam()))
}

type deleteTypeWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (d deleteTypeWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (d deleteTypeWriter) name() string {
	return "delete" + title(d.t.Name)
}

func (d deleteTypeWriter) idParam() string {
	return d.t.Name + "ID"
}

func (d deleteTypeWriter) params() *Statement {
	return Id(d.idParam()).Id(title(d.t.Name) + "ID")
}

func (d deleteTypeWriter) getElement() *Statement {
	return Id(d.t.Name).Op(":=").Id("engine").Dot(title(d.t.Name)).Call(Id(d.idParam())).Dot(d.t.Name)
}

func (d deleteTypeWriter) setOperationKind() *Statement {
	return Id(d.t.Name).Dot("OperationKind").Op("=").Id("OperationKindDelete")
}

func (d deleteTypeWriter) updateElementInPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(title(d.t.Name)).Index(Id(d.t.Name).Dot("ID")).Op("=").Id(d.t.Name)
}

func (d deleteTypeWriter) loopedElementIdentifier() string {
	return pluralizeClient.Singular(d.f.Name) + "ID"
}

func (d deleteTypeWriter) loopConditions() *Statement {
	return List(Id("_"), Id(d.loopedElementIdentifier())).Op(":=").Range().Id(d.t.Name).Dot(title(d.f.Name))
}

func (d deleteTypeWriter) deleteElementInLoop() *Statement {
	deleteFunc := Id("engine").Dot("delete" + title(d.f.ValueTypeName))
	if !d.f.HasPointerValue && d.f.HasAnyValue {
		return deleteFunc.Call(Id(d.loopedElementIdentifier()), True())
	}
	return deleteFunc.Call(Id(d.loopedElementIdentifier()))
}

func (d deleteTypeWriter) deleteElement() *Statement {
	deleteFunc := Id("engine").Dot("delete" + title(d.f.ValueTypeName))
	if !d.f.HasPointerValue && d.f.HasAnyValue {
		return deleteFunc.Call(Id(d.t.Name).Dot(title(d.f.Name)), True())
	}
	return deleteFunc.Call(Id(d.t.Name).Dot(title(d.f.Name)))
}

func (d deleteTypeWriter) existsInState() *Statement {
	return List(Id("_"), Id("ok")).Op(":=").Id("engine").Dot("State").Dot(title(d.t.Name)).Index(Id(d.idParam())).Id(";").Id("ok")
}

func (d deleteTypeWriter) deleteFromPatch() *Statement {
	return Delete(Id("engine").Dot("Patch").Dot(title(d.t.Name)), Id(d.idParam()))
}

func (d deleteTypeWriter) dereferenceField(field *ast.Field) *Statement {
	var suffix string
	if field.HasAnyValue {
		suffix = title(d.t.Name)
	}
	return Id("engine").Dot("dereference" + title(field.Parent.Name) + title(pluralizeClient.Singular(field.Name)) + "Refs" + suffix).Call(Id(d.idParam()))
}

type deleteGeneratedTypeWriter struct {
	f             ast.Field
	valueTypeName func() string
}

func (d deleteGeneratedTypeWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (d deleteGeneratedTypeWriter) name() string {
	return "delete" + title(d.valueTypeName())
}

func (d deleteGeneratedTypeWriter) idParam() string {
	return d.valueTypeName() + "ID"
}

func (d deleteGeneratedTypeWriter) params() *Statement {
	return Id(d.idParam()).Id(title(d.valueTypeName()) + "ID")
}

func (d deleteGeneratedTypeWriter) getElement() *Statement {
	return Id(d.valueTypeName()).Op(":=").Id("engine").Dot(d.valueTypeName()).Call(Id(d.idParam())).Dot(d.valueTypeName())
}

func (d deleteGeneratedTypeWriter) deleteChild() *Statement {
	return Id(d.valueTypeName()).Dot("deleteChild").Call()
}

func (d deleteGeneratedTypeWriter) deleteAnyContainer() *Statement {
	return Id("engine").Dot("delete"+title(anyNameByField(d.f))).Call(Id(d.valueTypeName()).Dot("ReferencedElementID"), False())
}

func (d deleteGeneratedTypeWriter) setOperationKind() *Statement {
	return Id(d.valueTypeName()).Dot("OperationKind").Op("=").Id("OperationKindDelete")
}

func (d deleteGeneratedTypeWriter) updateElementInPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(title(d.valueTypeName())).Index(Id(d.valueTypeName()).Dot("ID")).Op("=").Id(d.valueTypeName())
}

func (d deleteGeneratedTypeWriter) existsInState() *Statement {
	return List(Id("_"), Id("ok")).Op(":=").Id("engine").Dot("State").Dot(title(d.valueTypeName())).Index(Id(d.idParam())).Id(";").Id("ok")
}

func (d deleteGeneratedTypeWriter) deleteFromPatch() *Statement {
	return Delete(Id("engine").Dot("Patch").Dot(title(d.valueTypeName())), Id(d.idParam()))
}