package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStateMachine(t *testing.T) {
	t.Run("should add state machine declaration", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.addStateMachineDeclaration())
		expected := []string{
			_personDeclaration,
			_nameDeclaration, `
type state struct {
	person map[personID]person
	name map[nameID]name
}`, `
type stateMachine struct {
	state state
	patch state
	patchReceiver chan state
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addStateMachineDeclaration() *stateMachine {
	return sm
}
