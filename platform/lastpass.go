package platform

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/tobischo/gokeepasslib/v3"
)

// Lastpass documentation can be found at
// https://support.logmeininc.com/lastpass

// these 'lp***Key' consts must match directly to csv tags in LastpassEntry for conversion purposes
const (
	lpFolderKey     = "grouping"
	lpFavoriteKey   = "fav"
	lpNameKey       = "name"
	lpNotesKey      = "extra"
	lpLoginURLKey   = "url"
	lpLoginUserKey  = "username"
	lpLoginPassKey  = "password"
	LastpassProcess = "lpass"
)

type LastpassExport struct {
	Timestamp time.Time
	Entries   []LastpassEntry
}

type LastpassEntry struct {
	Folder        string `csv:"grouping"`
	Favorite      int    `csv:"fav"`
	Name          string `csv:"name"`
	Notes         string `csv:"extra"`
	LoginURI      string `csv:"url"`
	LoginUsername string `csv:"username"`
	LoginPassword string `csv:"password"`
}

func openLastpassExport(filepath string) (*LastpassExport, error) {

	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	entries := []LastpassEntry{}

	err = gocsv.UnmarshalFile(csvFile, &entries)
	if err != nil {
		return nil, err
	}

	newLPExport := &LastpassExport{
		Timestamp: time.Now(),
		Entries:   entries,
	}

	return newLPExport, nil
}

func unmarshalLastpassExport(content string) (*LastpassExport, error) {

	entries := []LastpassEntry{}

	err := gocsv.UnmarshalBytes([]byte(content), &entries)
	if err != nil {
		return nil, err
	}

	newLPExport := &LastpassExport{
		Timestamp: time.Now(),
		Entries:   entries,
	}

	return newLPExport, nil
}

func (this *LastpassExport) getName() string {
	name := fmt.Sprintf("%s-%s", Lastpass, this.Timestamp.Format("02.01.2006"))

	return name
}

func (this *LastpassExport) toKeepass() (*KeepassExport, error) {

	rootGroup := gokeepasslib.NewGroup()
	rootGroup.Name = this.getName()

	folderGroup := make(map[string]gokeepasslib.Group)

	for _, lpEntry := range this.Entries {
		kpEntry := gokeepasslib.NewEntry()
		kpEntry.Values = append(kpEntry.Values, mkValue(lpNameKey, lpEntry.Name))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpTitleKey, lpEntry.Name))

		kpEntry.Values = append(kpEntry.Values, mkValue(lpFavoriteKey,
			fmt.Sprintf("%d", lpEntry.Favorite)))

		kpEntry.Values = append(kpEntry.Values, mkValue(lpLoginURLKey, lpEntry.LoginURI))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpURLKey, lpEntry.LoginURI))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwLoginUserKey, lpEntry.LoginUsername))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpUserKey, lpEntry.LoginUsername))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwLoginPassKey, lpEntry.LoginPassword))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpPassKey, lpEntry.LoginPassword))

		kpEntry.Values = append(kpEntry.Values, mkValue(bwNotesKey, lpEntry.Notes))
		kpEntry.Values = append(kpEntry.Values, mkValue(kpNotesKey, lpEntry.Notes))

		if lpEntry.Folder != "" {

			_, folderExists := folderGroup[lpEntry.Folder]
			if !folderExists {
				newKeepassGroup := gokeepasslib.NewGroup()
				newKeepassGroup.Name = lpEntry.Folder

				folderGroup[lpEntry.Folder] = newKeepassGroup
			}

			bwFolder := folderGroup[lpEntry.Folder]
			bwFolder.Entries = append(bwFolder.Entries, kpEntry)

			folderGroup[lpEntry.Folder] = bwFolder
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
		Platform:     Lastpass,
		Timestamp:    this.Timestamp,
		KeepassGroup: rootGroup,
	}

	return newExport, nil
}
