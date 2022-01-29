package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteCreators(t *testing.T) {
	t.Run("writes creators", func(t *testing.T) {
		sf := newStateFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeCreators()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_CreateEquipmentSet_Engine_func,
			createEquipmentSet_Engine_func,
			_CreateGearScore_Engine_func,
			createGearScore_Engine_func,
			_CreateItem_Engine_func,
			createItem_Engine_func,
			_CreatePlayer_Engine_func,
			createPlayer_Engine_func,
			_CreatePosition_Engine_func,
			createPosition_Engine_func,
			_CreateZone_Engine_func,
			createZone_Engine_func,
			_CreateZoneItem_Engine_func,
			createZoneItem_Engine_func,
			createEquipmentSetEquipmentRef_Engine_func,
			createItemBoundToRef_Engine_func,
			createPlayerEquipmentSetRef_Engine_func,
			createPlayerGuildMemberRef_Engine_func,
			createPlayerTargetRef_Engine_func,
			createPlayerTargetedByRef_Engine_func,
			createAnyOfPlayer_Position_Engine_func,
			createAnyOfPlayer_ZoneItem_Engine_func,
			createAnyOfItem_Player_ZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}