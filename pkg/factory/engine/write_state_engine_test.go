package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteEngine(t *testing.T) {
	t.Run("writes operationKind", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeOperationKind()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_OperationKind_type,
			_OperationKindDelete_type,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes Engine", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeEngine()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_Engine_type,
			newEngine_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes generateID method", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeGenerateID()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_GenerateID_Engine_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes updateState method", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeUpdateState()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_UpdateState_Engine_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes importPatch method", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeImportPatch()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			importPatch_Engine_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
