package cmd

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/nextworking/citool/cit"
	"github.com/spf13/cobra"
	"os"
)

func init() {

	rootCmd.AddCommand(setVersion)
	setVersion.Flags().StringVar(&Path, "path",  "", "path")
	setVersion.MarkFlagRequired("path")

}

var GitVer semver.Version
var err error

var setVersion= &cobra.Command{
	Use:   "setVersion",
	Short: "Set puppet module version",
	Long:  `Set the puppet module version from metadata.json as a git tag. `,
	Run: func(cmd *cobra.Command, args []string) {

		p := fmt.Sprintf("%v/metadata.json", Path)
		metaVer := cit.GetMetadataVersion(p)
		gitTag := cit.GetGitTag(Path)

		if gitTag != "" {
			GitVer, err = semver.Parse(gitTag)
		} else {
			GitVer = metaVer
		}

		fmt.Printf("metadata version: %v\n", metaVer)
		fmt.Printf("git version: %v\n", GitVer)

		if metaVer.Compare(GitVer) <= 0 {
			if metaVer.Equals(GitVer) == true {
				fmt.Println("Probably WIP and/or no git tag. No actions taken")
			} else {
				fmt.Println("The metadata version is lower than the git tag. This must be fixed")
				os.Exit(1)
			}
		} else if metaVer.Compare(GitVer) == 1 {
			fmt.Println("metadata version > git tag. we are doing stuff now")
     		fmt.Printf("Adding tag %v\n", metaVer.String())
			fmt.Println(cit.FetchGetGitTags(Path))
			fmt.Println(cit.SetGitTag(Path,metaVer.String()))
			fmt.Println(cit.PushGitTag(Path))
		}

	},
}