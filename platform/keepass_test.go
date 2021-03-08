package platform

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertAndWrite(t *testing.T) {

	kpExport, err := ConvertCSV(Bitwarden, testBitwardenCSV)
	require.NoError(t, err)

	noPassword := ""

	filename, err := kpExport.Write(noPassword)
	require.Error(t, err)

	testPassword := "test"

	filename, err = kpExport.Write(testPassword)
	require.NoError(t, err)

	t.Log("created keepass db", filename)
}
