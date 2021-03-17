package actionsfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildActionsConfigAST(t *testing.T) {
	t.Run("builds expected actionsConfigAST", func(t *testing.T) {
		data := map[interface{}]interface{}{
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

		actual := buildActionsConfigAST(data)
		expected := &actionsConfigAST{
			Actions: map[string]action{
				"makeFoo": {
					Name: "makeFoo",
					Params: map[string]actionParameter{
						"entities": {
							Name:         "entities",
							TypeLiteral:  "entity",
							IsSliceValue: true,
							IsBasicType:  false,
						},
						"count": {
							Name:         "count",
							TypeLiteral:  "int",
							IsSliceValue: false,
							IsBasicType:  true,
						},
						"origins": {
							Name:         "origins",
							TypeLiteral:  "string",
							IsSliceValue: true,
							IsBasicType:  true,
						},
					},
				},
				"walkBar": {
					Name: "walkBar",
					Params: map[string]actionParameter{
						"distance": {
							Name:         "distance",
							TypeLiteral:  "float64",
							IsSliceValue: false,
							IsBasicType:  true,
						},
						"altitude": {
							Name:         "altitude",
							TypeLiteral:  "altitude",
							IsSliceValue: false,
							IsBasicType:  false,
						},
					},
				},
				"interactBaz": {
					Name: "interactBaz",
					Params: map[string]actionParameter{
						"target": {
							Name:         "target",
							TypeLiteral:  "bool",
							IsSliceValue: false,
							IsBasicType:  true,
						},
					},
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}
