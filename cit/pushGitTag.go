package cit

import (
	"fmt"
	"os"
	"os/exec"
)

func PushGitTag(d string) string {

	cmd := exec.Command("git", "push", "--tags")
	cmd.Dir = d
	out, err := cmd.CombinedOutput()
	fmt.Printf("path: %v\n", string(out))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in pushTag: %v\n", string(out))
		os.Exit(1)
	}

	return string(out)


}
