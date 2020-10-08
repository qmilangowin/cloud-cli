package commands

import (
	"strings"

	"com.elpigo/cli/internal/amazon"
	"github.com/spf13/cobra"
)

var bucketName string
var maxKeys int64

var s3RootCmd = &cobra.Command{
	Use:   "s3",
	Short: "s3 command building root command",
	Long: `s3 command that includes a set of commands to issue to
		s3 bucket operations. Run click-cloud s3 --help for more info`,
}

var s3ListAllCmd = &cobra.Command{
	Use:   "all",
	Short: "List All s3 buckets",
	Long: `Print in JSON format a list of all s3 buckets for your
		defined configuration and region`,
	// Args: cobra.MinimumNArgs(1),
	RunE: S3ListAllBucketsCmdFunc,
}

var s3ListBucketObjectsCmd = &cobra.Command{
	Use:   "files",
	Short: "List objects in a specific s3 bucket",
	Long: `Prints in JSON format contents of an s3 bucket. Set maxkeys flag (int64)
	to retreive the amount of objects`,
	Example: "cloud-cli files --bucket <bucket-name>",
	//Args: cobra.MinimumNArgs(1),
	RunE: S3ListBucketObjectsCmdFunc,
}

var s3UploadFileSingleCmd = &cobra.Command{
	Use:     "upload",
	Short:   "Upload file to s3 bucket",
	Long:    `Upload a file to s3, uploads only one file per time`,
	Example: "cloud-cli upload <file-name> --bucket <bucket-name>",
	Args:    cobra.MinimumNArgs(1),
	RunE:    S3UploadFileSingleCmdFunc,
}

//Command Functions

//S3ListAllBucketsCmdFunc will list all s3 buckets in AWS
func S3ListAllBucketsCmdFunc(cmd *cobra.Command, args []string) error {
	json, _ := cmd.Flags().GetBool("json")
	amazon.S3ListAll(json)
	return nil
}

//S3ListBucketObjects list all content in an S3 bucket
func S3ListBucketObjectsCmdFunc(cmd *cobra.Command, args []string) error {
	json, _ := cmd.Flags().GetBool("json")
	bucket, _ := cmd.Flags().GetString("bucket")
	maxkeys, _ := cmd.Flags().GetInt64("maxkeys")
	amazon.S3ListBucketObject(bucket, maxkeys, json)
	return nil
}

//S3UploadFileSingleCmdFunc will upload a single file
func S3UploadFileSingleCmdFunc(cmd *cobra.Command, args []string) error {
	filename := strings.Join(args, " ")
	bucket, _ := cmd.Flags().GetString("bucket")
	amazon.S3UploadFileSingle(filename, bucket)
	return nil
}

func init() {

	//commands
	rootCmd.AddCommand(s3RootCmd)
	s3RootCmd.AddCommand(s3ListAllCmd, s3UploadFileSingleCmd)
	s3RootCmd.AddCommand(s3ListBucketObjectsCmd)

	//flags

	//list objects flags
	s3ListBucketObjectsCmd.Flags().StringVarP(&bucketName, "bucket", "b", "", "name of S3 bucket")
	s3ListBucketObjectsCmd.MarkFlagRequired("bucket")
	s3ListBucketObjectsCmd.Flags().Int64("maxkeys", 10, "Amount of objects to return")

	//upload single file flags
	s3UploadFileSingleCmd.Flags().StringVarP(&bucketName, "bucket", "b", "", "name of S3 bucket")
	s3UploadFileSingleCmd.MarkFlagRequired("bucket")

}
