package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Formatter uint

const (
	CommitHash Formatter = iota + 1
	CommitterName
	CommitterEmail
	CommitterDate
)

var (
	FormatterName = map[Formatter]string{
		CommitHash:     "%H",
		CommitterName:  "%cn",
		CommitterEmail: "%ce",
		CommitterDate:  "%ci",
	}
	FormatterValue = map[string]Formatter{}
)

func (f Formatter) String() string {
	return FormatterName[f]
}

func GetGitRepoInfo(f Formatter) string {
	formatter := fmt.Sprintf("--format=format:%s", f.String())
	return chomp(run(".", "git", "log", "-n", "1", formatter))
}

func IsGitRepo(wd string) bool {
	gitDir := chomp(run(wd, "git", "rev-parse", "--git-dir"))
	if !filepath.IsAbs(gitDir) {
		wd, err := os.Getwd()
		if err != nil {
			fatalf("FAILED: %v: %v", "os.Getwd()", err)
		}
		gitDir = filepath.Join(wd, gitDir)
	}
	return isDir(gitDir)
}

func chomp(s string) string {
	return strings.TrimRight(s, " \t\r\n")
}

func isDir(p string) bool {
	fi, err := os.Stat(p)
	return err == nil && fi.IsDir()
}

func run(dir string, cmd ...string) string {
	xcmd := exec.Command(cmd[0], cmd[1:]...)
	xcmd.Dir = dir
	data, err := xcmd.CombinedOutput()
	if err != nil {
		fatalf("FAILED: %v: %v", strings.Join(cmd, " "), err)
	}
	return string(data)
}

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "auto build version: %s\n", fmt.Sprintf(format, args...))
	os.Exit(2)
}
