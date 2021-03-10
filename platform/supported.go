package platform

const Bitwarden = "bitwarden"

var SupportedPlatforms = make(map[string]bool)

func init() {
	SupportedPlatforms[Bitwarden] = true
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
