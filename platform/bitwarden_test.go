package platform

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testBitwardenCSV = "example/bitwarden_example.csv"
)

func TestOpenBitwardenExportFile(t *testing.T) {

	bwExport, err := openBitwardenExport(testBitwardenCSV)
	require.NoError(t, err)

	require.Equal(t, 4, len(bwExport.Entries))

	for _, bwEntry := range bwExport.Entries {
		require.NotEqual(t, bwEntry.Name, "")
	}
}

func TestConvertBitwarden(t *testing.T) {

	bwExport, err := openBitwardenExport(testBitwardenCSV)
	require.NoError(t, err)

	require.Equal(t, 4, len(bwExport.Entries))

	exportName := bwExport.getName()

	kpExport, err := bwExport.toKeepass()
	require.NoError(t, err)

	require.Equal(t, kpExport.KeepassGroup.Name, exportName)

}
