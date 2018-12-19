package cit

import (
	"fmt"
	"os/exec"
)


func FetchGetGitTags(d string) string {

	fmt.Println("fetching tags")
	cmd := exec.Command("git", "fetch", "--tags")
	cmd.Dir = d
	out, _ := cmd.CombinedOutput()

	return string(out)
}

