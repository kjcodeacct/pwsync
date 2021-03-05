package platform

// Bitwarden documentation can be found at
// https://bitwarden.com/help/

import (
	"os"

	"github.com/gocarina/gocsv"
)

type BitwardenExport struct {
	Folder        string `csv:"folder"`
	Favorite      int    `csv:"favorite"`
	Type          string `csv:"type"`
	Name          string `csv:"name"`
	Notes         string `csv:"notes"`
	Fields        string `csv:"fields"`
	LoginURI      string `csv:"login_uri"`
	LoginUserName string `csv:"login_username"`
	LoginPassword string `csv:"login_password"`
	LoginTTPSeed  string `csv:"login_totp"`
}

func openBitwardenExportFile(filepath string) ([]BitwardenExport, error) {

	newBitwardenExport := []BitwardenExport{}

	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	err = gocsv.UnmarshalFile(csvFile, &newBitwardenExport)
	if err != nil {
		return nil, err
	}

	return newBitwardenExport, nil
}
