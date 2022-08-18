package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes the installed configs",
	Long:  "Removes the config files for installed programs",
	Run: func(cmd *cobra.Command, args []string) {
		removeInstalled(args)
	},
}

func removeInstalled(packages []string) {
	fmt.Println("Checking if the packages are installed.")
	for i := 0; i < len(packages); i++ {
		fmt.Printf("%s found!\n", packages[i])
	}
	stageDir := shareDir() // Get the staging directory
	for i := 0; i < len(packages); i++ {
		dir := path.Join(stageDir, packages[i]) // Join staging directory and the package
		finalDir := handleDir(dir)              // Handle "%APPDATA%" and "~"
		fmt.Printf("Removing %s\n", finalDir)
		_, err := os.Lstat(finalDir)
		if err == nil {
			err := os.RemoveAll(finalDir)
			if err != nil {
				fmt.Printf("\nCould not remove %s", finalDir)
				os.Exit(1)
			}
		}
	}
}
