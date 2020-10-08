package commands

import (
	"com.elpigo/cli/internal/helpers"
	"github.com/spf13/cobra"
)

var zipRootCmd = &cobra.Command{
	Use:     "zip",
	Short:   "zip root command",
	Long:    `Allows you to compress multiple files. `,
	Example: "cloud-cli zip <file-name(s)>",
	Args:    cobra.MinimumNArgs(1),
	Run:     zipRootCmdFunc,
}

//Command Functions

//S3UploadFileSingleCmdFunc will upload a single file
func zipRootCmdFunc(cmd *cobra.Command, args []string) {
	json, _ := cmd.Flags().GetBool("json")
	helpers.Zip(args, json)
}

func init() {

	//commands
	rootCmd.AddCommand(zipRootCmd)

}
