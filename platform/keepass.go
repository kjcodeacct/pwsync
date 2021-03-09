package platform

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/tobischo/gokeepasslib/v3"
)

const (
	kpExtension = ".kdbx"
	kpTitleKey  = "Title"
	kpUserKey   = "UserName"
	kpPassKey   = "Password"
	kpNotesKey  = "Notes"
	kpURLKey    = "URL"

	kpBackupName = "pwsync_backup"
)

type KeepassExport struct {
	Platform     string
	Timestamp    time.Time
	ShaHash      string
	KeepassGroup gokeepasslib.Group
}

func ConvertCSV(platform string, filepath string) (*KeepassExport, error) {

	newExport := &KeepassExport{}

	switch platform {
	case Bitwarden:
		bwExport, err := openBitwardenExport(filepath)
		if err != nil {
			return nil, err
		}

		newExport, err = bwExport.toKeepass()
		if err != nil {
			return nil, err
		}

	}

	return newExport, nil
}

func (this *KeepassExport) Write(password string) (string, error) {

	if password == "" {
		return "", fmt.Errorf("password must be provided")
	}

	keepassDB := &gokeepasslib.Database{
		Header:      gokeepasslib.NewHeader(),
		Credentials: gokeepasslib.NewPasswordCredentials(password),
		Content: &gokeepasslib.DBContent{
			Meta: gokeepasslib.NewMetaData(),
			Root: &gokeepasslib.RootData{
				Groups: []gokeepasslib.Group{this.KeepassGroup},
			},
		},
	}

	keepassDB.LockProtectedEntries()

	buff := new(bytes.Buffer)
	keepassEncoder := gokeepasslib.NewEncoder(buff)
	err := keepassEncoder.Encode(keepassDB)
	if err != nil {
		return "", err
	}

	hash := sha256.New()

	_, err = hash.Write(buff.Bytes())
	if err != nil {
		return "", err
	}

	shaHash := hex.EncodeToString(hash.Sum(nil))
	this.ShaHash = shaHash

	filename := this.getName()
	filename = filename + kpExtension

	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = file.Write(buff.Bytes())
	if err != nil {
		return "", err
	}

	return filename, nil
}

func mkValue(key string, value string) gokeepasslib.ValueData {
	return gokeepasslib.ValueData{Key: key, Value: gokeepasslib.V{Content: value}}
}

func (this *KeepassExport) getName() string {

	shortHash := this.ShaHash[0:7]
	name := fmt.Sprintf("%s-%s-%s", this.Platform, this.Timestamp.Format("02.01.2006"), shortHash)

	return name
}
