package platform

const Bitwarden = "bitwarden"

var SupportedPlatforms = make(map[string]bool)

func init() {
	SupportedPlatforms[Bitwarden] = true
}

func Check(input string) bool {
	return SupportedPlatforms[input]
}
