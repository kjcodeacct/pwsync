package platform

const (
	Bitwarden = "bitwarden"
	Lastpass  = "lastpass"
)

var SupportedPlatforms = make(map[string]bool)

func init() {
	SupportedPlatforms[Bitwarden] = true
	SupportedPlatforms[Lastpass] = true
}

func Check(input string) bool {
	return SupportedPlatforms[input]
}

func GetSupportedPlatforms() []string {
	var supportedList []string

	for platform, _ := range SupportedPlatforms {
		supportedList = append(supportedList, platform)
	}

	return supportedList
}
