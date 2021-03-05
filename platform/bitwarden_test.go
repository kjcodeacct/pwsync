package platform

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpenBitwardenExportFile(t *testing.T) {

	testFile := "bitwarden_example.csv"

	bwExport, err := openBitwardenExportFile(testFile)
	require.NoError(t, err)

	require.Equal(t, 4, len(bwExport))

	for _, bwEntry := range bwExport {
		require.NotEqual(t, bwEntry.Name, "")
	}
}
