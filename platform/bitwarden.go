package platform

// Bitwarden documentation can be found at
// https://bitwarden.com/help/

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/tobischo/gokeepasslib/v3"
)

// these 'bw***Key' consts must match directly to csv tags in 'BitwardenEntry for conversion purposes
const (
	bwFolderKey    = "folder"
	bwFavoriteKey  = "favorite"
	bwTypeKey      = "type"
	bwNameKey      = "name"
	bwNotesKey     = "notes"
	bwFieldsKey    = "fields"
	bwLoginURIKey  = "login_uri"
	bwLoginUserKey = "login_username"
	bwLoginPassKey = "login_password"
	bwLoginSeedKey = "login_totp"
)

type BitwardenExport struct {
	Timestamp time.Time
	Entries   []BitwardenEntry
}

type BitwardenEntry struct {
	Folder        string `csv:"folder"`
	Favorite      int    `csv:"favorite"`
	Type          string `csv:"type"`
	Name          string `csv:"name"`
	Notes         string `csv:"notes"`
	Fields        string `csv:"fields"`
	LoginURI      string `csv:"login_uri"`
	LoginUserName string `csv:"login_username"`
	LoginPassword string `csv:"login_password"`
	LoginTOTPSeed string `csv:"login_totp"`
}

func openBitwardenExport(filepath string) (*BitwardenExport, error) {

	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	entries := []BitwardenEntry{}

	err = gocsv.UnmarshalFile(csvFile, &entries)
	if err != nil {
		return nil, err
	}

	newBWExport := &BitwardenExport{
		Timestamp: time.Now(),
		Entries:   entries,
	}

	return newBWExport, nil
}

func unmarshalBitwardenExport(content string) (*BitwardenExport, error) {

	entries := []BitwardenEntry{}

	err := gocsv.UnmarshalBytes([]byte(content), &entries)
	if err != nil {
		return nil, err
	}

	newBWExport := &BitwardenExport{
		Timestamp: time.Now(),
		Entries:   entries,
	}

	return newBWExport, nil
}

func (this *BitwardenExport) getName() string {
	name := fmt.Sprintf("%s-%s", Bitwarden, this.Timestamp.Format("02.01.2006"))

	return name
}

func (this *BitwardenExport) toKeepass() (*KeepassExport, error) {

	rootGroup := gokeepasslib.NewGroup()
	rootGroup.Name = this.getName()

	folderGroup := make(map[string]gokeepasslib.Group)

	for _, bwEntry := range this.Entries {
		kpEntry := gokeepasslib.NewEntry()

		kpEntry.Values = append(kpEntry.Values, mkValue(bwNameKey, bwEntry.Name))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpTitleKey, bwEntry.Name))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwTypeKey, bwEntry.Type))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwFavoriteKey,
			fmt.Sprintf("%d", bwEntry.Favorite)))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwLoginURIKey, bwEntry.LoginURI))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpURLKey, bwEntry.LoginURI))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwLoginUserKey, bwEntry.LoginUserName))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpUserKey, bwEntry.LoginUserName))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwLoginPassKey, bwEntry.LoginPassword))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpPassKey, bwEntry.LoginPassword))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwLoginSeedKey, bwEntry.LoginTOTPSeed))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwNotesKey, bwEntry.Notes))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpNotesKey, bwEntry.Notes))

		if bwEntry.Fields != "" {
			kpEntry.Values = append(kpEntry.Values, mkValue(bwFieldsKey, bwEntry.Fields))
		}

		if bwEntry.Folder != "" {

			_, folderExists := folderGroup[bwEntry.Folder]
			if !folderExists {
				newKeepassGroup := gokeepasslib.NewGroup()
				newKeepassGroup.Name = bwEntry.Folder

				folderGroup[bwEntry.Folder] = newKeepassGroup
			}

			bwFolder := folderGroup[bwEntry.Folder]
			bwFolder.Entries = append(bwFolder.Entries, kpEntry)

			folderGroup[bwEntry.Folder] = bwFolder
		} else {
			rootGroup.Entries = append(rootGroup.Entries, kpEntry)
		}

	}

	for _, group := range folderGroup {
		rootGroup.Groups = append(rootGroup.Groups, group)
	}

	csvContent, err := gocsv.MarshalString(&this.Entries)
	if err != nil {
		return nil, err
	}

	csvGroup := gokeepasslib.NewGroup()
	csvGroup.Name = kpBackupName

	csvEntry := gokeepasslib.NewEntry()
	csvEntry.Values = append(csvEntry.Values, mkValue(kpTitleKey, kpBackupName))
	csvEntry.Values = append(csvEntry.Values, mkValue(kpNotesKey, csvContent))

	csvGroup.Entries = append(csvGroup.Entries, csvEntry)
	rootGroup.Groups = append(rootGroup.Groups, csvGroup)

	newExport := &KeepassExport{
		Platform:     Bitwarden,
		Timestamp:    this.Timestamp,
		KeepassGroup: rootGroup,
	}

	return newExport, nil
}
