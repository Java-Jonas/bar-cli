package serverfactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeProcessClientMessage() *ServerFactory {
	decls := NewDeclSet()

	p := processClientMessageWriter{}

	decls.File.Func().Params(p.receiverParams()).Id("processClientMessage").Params(p.params()).Params(Id("message"), Id("error")).Block(
		Switch().Id("messageKind").Call(Id("msg").Dot("Kind")).Block(
			ForEachActionInAST(s.config, func(action ast.Action) *Statement {
				p.a = &action
				return Case(Id(p.actionMessageKind())).Block(
					p.declareParams(),
					p.unmarshalMessageContent(),
					If(Id("err").Op("!=").Nil()).Block(
						Return(Id("message").Values(), Id("err")),
					),
					p.callAction(),
					OnlyIf(action.Response != nil, p.marshalResponseContent()),
					OnlyIf(action.Response != nil, p.returnMarshallingError()),
					p.returnResponse(),
				)
			}),
			Default().Block(
				Return(Id("message").Values(), Id("fmt").Dot("Errorf").Call(Lit("unknown message kind in: %s"), Id("printMessage").Call(Id("msg")))),
			),
		),
	)

	decls.Render(s.buf)
	return s
}

type processClientMessageWriter struct {
	a *ast.Action
	p *ast.Field
}

func (p processClientMessageWriter) receiverParams() *Statement {
	return Id("r").Id("*Room")
}

func (p processClientMessageWriter) params() *Statement {
	return Id("msg").Id("message")
}

func (p processClientMessageWriter) actionMessageKind() string {
	return "messageKindAction_" + p.a.Name
}

func (p processClientMessageWriter) declareParams() *Statement {
	return Var().Id("params").Id(Title(p.a.Name) + "Params")
}

func (p processClientMessageWriter) unmarshalMessageContent() *Statement {
	return Id("err").Op(":=").Id("params").Dot("UnmarshalJSON").Call(Id("msg").Dot("Content"))
}

func (p processClientMessageWriter) callAction() *Statement {
	call := Id("r").Dot("actions").Dot(p.a.Name).Call(Id("params"), Id("r").Dot("state"))
	if p.a.Response != nil {
		return Id("res").Op(":=").Add(call)
	}
	return call
}

func (p processClientMessageWriter) marshalResponseContent() *Statement {
	return List(Id("resContent"), Id("err")).Op(":=").Id("res").Dot("MarshalJSON").Call()
}

func (p processClientMessageWriter) returnMarshallingError() *Statement {
	return If(Id("err").Op("!=").Nil()).Block(
		Return(Id("message").Values(), Id("err")),
	)
}

func (p processClientMessageWriter) returnResponse() *Statement {
	if p.a.Response == nil {
		return Return(Id("message").Values(), Nil())
	}
	return Return(Id("message").Values(Id("msg").Dot("Kind"), Id("resContent"), Id("msg").Dot("client")), Nil())
}
