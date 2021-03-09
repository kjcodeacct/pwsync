package platform

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tobischo/gokeepasslib/v3"
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

	file, err := os.Open(filename)
	require.NoError(t, err)

	db := gokeepasslib.NewDatabase()

	db.Credentials = gokeepasslib.NewPasswordCredentials(testPassword)
	err = gokeepasslib.NewDecoder(file).Decode(db)
	require.NoError(t, err)

	err = db.UnlockProtectedEntries()
	require.NoError(t, err)

	bwExport, err := openBitwardenExport(testBitwardenCSV)
	require.NoError(t, err)

	var backupContent string

	for _, rootGroup := range db.Content.Root.Groups {
		for _, group := range rootGroup.Groups {
			if group.Name == kpBackupName {
				for _, entry := range group.Entries {
					if entry.GetTitle() == kpBackupName {
						temp := entry.Get(kpNotesKey)
						backupContent = temp.Value.Content

					}
				}

			}
		}
	}

	compareBwExport, err := unmarshalBitwardenExport(backupContent)
	require.NoError(t, err)

	require.Equal(t, len(bwExport.Entries), len(compareBwExport.Entries))
	require.Equal(t, bwExport.Entries, compareBwExport.Entries)

	err = os.Remove(filename)
	t.Log("cleaned up test files")
}
