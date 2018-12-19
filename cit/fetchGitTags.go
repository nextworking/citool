package cit

import (
	"os/exec"
)


func FetchGetGitTags(d string) string {

	cmd := exec.Command("git", "fetch", "--tags")
	cmd.Dir = d
	out, _ := cmd.CombinedOutput()

	return string(out)
}

