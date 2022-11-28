package core

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/depocket/near-lake-framework-go/types"
)

type LakeConfig struct {
	s3BucketName          string
	s3RegionName          string
	startBlockHeight      uint64
	s3Config              *client.ConfigProvider
	blocksPreloadPoolSize uint64
}

func Streamer(config LakeConfig) chan types.StreamMessage {
	messageChannel := make(chan types.StreamMessage, 1)
	go func(cfg LakeConfig, mc chan types.StreamMessage) {
		start(cfg, mc)
	}(config, messageChannel)
	return messageChannel
}

func start(config LakeConfig, messageChannel chan types.StreamMessage) {
	var startFromBlockHeight = config.startBlockHeight
	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	s3Client := s3.New(awsSession)
	if config.s3Config != nil {
		s3Client = s3.New(*config.s3Config)
	}
	for {
		s3Fetcher := S3Fetcher{}
		blockHeightsPrefixes, err := s3Fetcher.ListBlocks(s3Client, config.s3BucketName, startFromBlockHeight, config.blocksPreloadPoolSize*2)
		if err != nil {
			fmt.Println(err)
		}
		if len(blockHeightsPrefixes) <= 0 {
			continue
		}
		startFromBlockHeight = blockHeightsPrefixes[len(blockHeightsPrefixes)-1] + 1

		for _, blockHeight := range blockHeightsPrefixes {
			message, err := s3Fetcher.FetchStreamerMessage(s3Client, config.s3BucketName, blockHeight)
			if err != nil {
				fmt.Println(err)
				continue
			}
			messageChannel <- *message
		}
	}
}
