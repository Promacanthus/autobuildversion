package autobuildversion

import (
	"fmt"
	"runtime"

	"github.com/Promacanthus/autobuildversion/internel/git"
	"github.com/Promacanthus/autobuildversion/internel/version"
)

func init() {
	if !git.IsGitRepo(".") {
		panic("FAILED: not a Git repo.")
	}

	v := version.Info{
		GoVersion:      runtime.Version(),
		Platform:       fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		CommitHash:     git.GetGitRepoInfo(git.CommitHash),
		CommitterName:  git.GetGitRepoInfo(git.CommitterName),
		CommitterEmail: git.GetGitRepoInfo(git.CommitterEmail),
		CommitterDate:  git.GetGitRepoInfo(git.CommitterDate),
	}

	fmt.Printf("App Info:\n\n%s\n\n", v)
}

