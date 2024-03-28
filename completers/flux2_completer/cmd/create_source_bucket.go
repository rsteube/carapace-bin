package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_source_bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Create or update a Bucket source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_source_bucketCmd).Standalone()
	create_source_bucketCmd.Flags().String("access-key", "", "the bucket access key")
	create_source_bucketCmd.Flags().String("bucket-name", "", "the bucket name")
	create_source_bucketCmd.Flags().String("endpoint", "", "the bucket endpoint address")
	create_source_bucketCmd.Flags().Bool("insecure", false, "for when connecting to a non-TLS S3 HTTP endpoint")
	create_source_bucketCmd.Flags().SourceBucketProvider("provider", generic, "the S3 compatible storage provider name, available options are: (generic, aws)")
	create_source_bucketCmd.Flags().String("region", "", "the bucket region")
	create_source_bucketCmd.Flags().String("secret-key", "", "the bucket secret key")
	create_source_bucketCmd.Flags().String("secret-ref", "", "the name of an existing secret containing credentials")
	create_sourceCmd.AddCommand(create_source_bucketCmd)
}
