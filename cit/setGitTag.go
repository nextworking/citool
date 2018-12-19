package cit

import (
	"fmt"
	"os"
	"os/exec"
)

func SetGitTag(d string, ver string) string {

	cmd := exec.Command("git", "tag", "-a", ver, "-m", "\"gitlab ci tag\"")
	cmd.Dir = d
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in setTag: %v\n", string(out))
		os.Exit(1)
	}

	return string(out)
}
