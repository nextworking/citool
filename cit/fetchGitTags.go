package cit

import (
	"fmt"
	"os"
	"os/exec"
)

func FetchGetGitTags(d string) string {

	fmt.Println("fetching tags")
	cmd := exec.Command("git", "fetch", "--tags")
	cmd.Dir = d
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in fetchGitTag: %v\n", string(out))
		os.Exit(1)
	}
	return string(out)
}
