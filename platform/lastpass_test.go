package platform

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testLastpassCSV = "example/lastpass_example.csv"
)

func TestOpenLastpassExportFile(t *testing.T) {

	lpExport, err := openLastpassExport(testLastpassCSV)
	require.NoError(t, err)

	require.Equal(t, 6, len(lpExport.Entries))

	for _, bwEntry := range lpExport.Entries {
		require.NotEqual(t, bwEntry.Name, "")
	}

}

func TestConvertLastpass(t *testing.T) {

	lpExport, err := openLastpassExport(testLastpassCSV)
	require.NoError(t, err)

	require.Equal(t, 6, len(lpExport.Entries))

	exportName := lpExport.getName()

	kpExport, err := lpExport.toKeepass()
	require.NoError(t, err)

	require.Equal(t, kpExport.KeepassGroup.Name, exportName)
}
