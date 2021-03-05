package platform

type Export struct {
	Platform string
	Data     interface{}
}

func OpenCSV(platform string, filename string) (*Export, error) {

	switch platform {
	case Bitwarden:

	}

	return nil, nil
}
