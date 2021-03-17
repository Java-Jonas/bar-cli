package statefactory

func newSimpleASTExample() *stateConfigAST {
	data := map[interface{}]interface{}{
		"player": map[interface{}]interface{}{
			"items":     "[]item",
			"gearScore": "gearScore",
			"position":  "position",
		},
		"zone": map[interface{}]interface{}{
			"items":   "[]zoneItem",
			"players": "[]player",
			"tags":    "[]string",
		},
		"zoneItem": map[interface{}]interface{}{
			"position": "position",
			"item":     "item",
		},
		"position": map[interface{}]interface{}{
			"x": "float64",
			"y": "float64",
		},
		"item": map[interface{}]interface{}{
			"gearScore": "gearScore",
		},
		"gearScore": map[interface{}]interface{}{
			"level": "int",
			"score": "int",
		},
	}

	// TODO: make prettier
	simpleAST := buildRudimentaryStateConfigAST(data)
	simpleAST.fillInReferences().fillInParentalInfo()

	return simpleAST
}
