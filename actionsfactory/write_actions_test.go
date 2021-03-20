package actionsfactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

var testActionsConfig = map[interface{}]interface{}{
	"makeFoo": map[interface{}]interface{}{
		"entities": "[]entity",
		"count":    "int",
		"origins":  "[]string",
	},
	"walkBar": map[interface{}]interface{}{
		"distance": "float64",
		"altitude": "altitude",
	},
	"interactBaz": map[interface{}]interface{}{
		"target": "bool",
	},
}

func TestWriteActions(t *testing.T) {
	t.Run("writes actions", func(t *testing.T) {

		ast := buildActionsConfigAST(testActionsConfig)
		af := newActionsFactory(ast)

		actual := utils.NormalizeWhitespace(string(af.writeActions().writtenSourceCode()))
		expected := utils.NormalizeWhitespace(strings.TrimSpace(`
func interactBaz(target bool, sm *statemachine.StateMachine) {}
func makeFoo(count int, entities []statemachine.Entity, origins []string, sm *statemachine.StateMachine) {}
func walkBar(altitude statemachine.Altitude, distance float64, sm *statemachine.StateMachine) {}
		`))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}