package pkg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSAuth struct {
	AccessKey string
	SecretKey string
	Token     string
}

type AwsS3Iterator struct {
	S3Svc       *s3.S3
	StorageTree *S3DataStorageTree
}

func NewS3Iterator(s3AccssObj *s3.S3) *AwsS3Iterator {
	return &AwsS3Iterator{
		S3Svc: s3AccssObj,
	}
}

func NewS3Service(awsAccessKey, awsSecret, token, region string) (*AwsS3Iterator, error) {
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)
	_, err := creds.Get()
	if err != nil {
		return nil, errors.New("Invalid Credentials")
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)
	return NewS3Iterator(svc), nil
}

type S3DataStorageTree struct {
	IsExist        bool
	MapToNextLevel map[string]*S3DataStorageTree
}

func (s3iter *AwsS3Iterator) GetAllObjectPath(bucket string) []string {
	allPaths := make([]string, 0)
	var token string
	response := &s3.ListObjectsV2Output{}
	var err error
	var continuationToken *string
	for {
		if token == "" {
			continuationToken = nil
		}
		response, err = s3iter.S3Svc.ListObjectsV2(&s3.ListObjectsV2Input{
			Bucket:            aws.String(bucket),
			ContinuationToken: continuationToken,
		})
		if err != nil {
			return nil
		}
		for _, currObj := range response.Contents {
			allPaths = append(allPaths, *currObj.Key)
		}
		if !*response.IsTruncated {
			continuationToken = response.ContinuationToken
			break
		}
	}
	return allPaths
}

func (s3iter *AwsS3Iterator) InsertInDataStore(path string, bucket string) {
	cleanPath := strings.TrimRight(path, "/")
	x := strings.Split(cleanPath, "/")
	pathSubSection := append([]string{bucket}, x...)
	tempPtr := (*s3iter.StorageTree)
	for _, path := range pathSubSection {
		if _, present := tempPtr.MapToNextLevel[path]; !present {
			store := S3DataStorageTree{IsExist: false, MapToNextLevel: make(map[string]*S3DataStorageTree)}
			tempPtr.MapToNextLevel[path] = &store
		}
		tempPtr = (*tempPtr.MapToNextLevel[path])
	}
}

func (s3iter *AwsS3Iterator) GenerateS3ObjectTreeForPath(bucketList []string, path []string) (*S3DataStorageTree, error) {
	s3iter.StorageTree = &S3DataStorageTree{IsExist: true, MapToNextLevel: make(map[string]*S3DataStorageTree)}
	tempPtr := (*s3iter.StorageTree)
	for _, b := range bucketList {
		newSubTree := &S3DataStorageTree{IsExist: true, MapToNextLevel: make(map[string]*S3DataStorageTree)}
		tempPtr.MapToNextLevel[b] = newSubTree
	}
	for _, bucket := range bucketList {
		if bucket == path[0] {
			keypathList := s3iter.GetAllObjectPath(bucket)
			for _, path := range keypathList {
				s3iter.InsertInDataStore(path, bucket)
			}
		}
	}
	for k, v := range s3iter.StorageTree.MapToNextLevel {
		fmt.Println(k, v)
	}

	return &S3DataStorageTree{}, nil
}

func (s3iter *AwsS3Iterator) ListDir(path string) ([]string, error) {
	result, err := s3iter.S3Svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		}
		return nil, errors.New(fmt.Sprintf("Failed to Fetch List of buckets %v", err.Error()))
	}
	var bucketList []string
	for _, rs := range result.Buckets {
		bucketList = append(bucketList, *rs.Name)
	}
	pathList := strings.Split(path, "/")
	if path != "" {
		s3iter.GenerateS3ObjectTreeForPath(bucketList, pathList)
		return []string{"hell"}, nil
	}
	return bucketList, nil
}
