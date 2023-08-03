package version

// Version info.
var (
	Version = "0.14.0"
	Meta    = "dev"
)

// Git commit/date info, set via linker flags.
var (
	GitCommit = ""
	GitDate   = ""
)

// VersionWithCommit returns a textual version string including Git commit/date
// information.
func VersionWithCommit() string {
	vsn := Version + "-" + Meta
	if len(GitCommit) >= 8 {
		vsn += "-" + GitCommit[:8]
	}
	if (Meta != "stable") && (GitDate != "") {
		vsn += "-" + GitDate
	}
	return vsn
}
