package amazon

import (
	"fmt"
	"log"
	"os"
	"strings"

	"com.elpigo/cli/internal/authentication"
	"com.elpigo/cli/internal/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cheggaaa/pb"
)

//S3ListAll will list all s3 buckets

func s3Session() (*session.Session, error) {

	session, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: authentication.AwsAuth.AwsProfile,
		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String(authentication.AwsAuth.AwsRegion),
		},
		// // Force enable Shared Config support
		// SharedConfigState: session.SharedConfigEnable,
	})

	return session, err

}

func S3ListAll(json bool) {
	session, err := s3Session()
	svc := s3.New(session)
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	if json {
		fmt.Println(result)

	} else {
		data := []string{}
		for _, v := range result.Buckets {
			data = append(data, *v.Name)
		}
		fmt.Println("--- S3 Buckets ---")
		fmt.Println(strings.Join(data, "\n"))
	}

}

func S3ListBucketObject(bucket string, maxkeys int64, json bool) {
	session, err := s3Session()
	svc := s3.New(session)

	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		MaxKeys: aws.Int64(maxkeys),
	}

	result, err := svc.ListObjectsV2(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(fmt.Errorf("No such bucket - please check name or credentials"))
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	if json {
		fmt.Println(*result)

	} else {
		for _, v := range result.Contents {
			fmt.Println(*v.Key)
		}
	}
}

func S3UploadFileSingle(filename string, bucket string) {
	session, err := s3Session()
	// svc := s3.New(session)
	stat, err := os.Stat(filename)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fp, err := os.Open(filename)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	progBar := pb.New64(stat.Size()).SetUnits(pb.U_BYTES)
	fmt.Println("Uploading: ", filename)
	progBar.Start()
	defer progBar.Finish()

	r := &helpers.ProgressReader{
		ProgBar: progBar,
		Fp:      fp,
		Size:    stat.Size(),
		Reads:   -stat.Size(),
	}

	uploader := s3manager.NewUploader(session, func(u *s3manager.Uploader) {
		u.PartSize = 200 * 1024 * 1024
		u.LeavePartsOnError = true
		u.Concurrency = 10
	})

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   r,
	})

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

}
