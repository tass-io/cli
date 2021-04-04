package version

import (
	"log"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	AppVersion            = ""
	LocalSchedulerVersion = ""
	BuildTime             = ""
	GitCommit             = ""
)

type VersionInfo struct {
	AppVersion            string
	LocalSchedulerVersion string
	GoVersion             string
	Compiler              string
	Platform              string
	BuildTime             string
	GitCommit             string
}

func (info *VersionInfo) String() string {
	return "{Tass Version: " + info.AppVersion + ", Local scheduler version: " + info.LocalSchedulerVersion +
		", Go version: " + info.GoVersion + ", Compiler version: " + info.Compiler + ", Platform: " + info.Platform +
		", Build time: " + info.BuildTime + ", Git commit: " + info.GitCommit + "}"
}

func GetVersionInfo(cmd *cobra.Command, args []string) {
	info := &VersionInfo{
		AppVersion:            AppVersion,
		LocalSchedulerVersion: LocalSchedulerVersion,
		GoVersion:             runtime.Version(),
		Compiler:              runtime.Compiler,
		Platform:              runtime.GOOS + "/" + runtime.GOARCH,
		BuildTime:             BuildTime,
		GitCommit:             GitCommit,
	}
	log.Println(info.String())
}
