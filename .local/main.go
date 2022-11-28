package main

import (
	"fmt"
	"github.com/depocket/near-lake-framework-go/core"
)

func main() {
	config := core.DefaultLakeConfigBuilder().
		Mainnet().SetStartBlockHeight(79075963).SetBlocksPreloadPoolSize(100).Build()
	channel := core.Streamer(*config)
	for {
		select {
		case message := <-channel:
			fmt.Println(message.Shards[0].Chunk.Receipts)
		}
	}
	//s3Fetcher := core.S3Fetcher{}
	//awsSession, _ := session.NewSession(&aws.Config{
	//	Region: aws.String("eu-central-1"),
	//})
	//objects, err := s3Fetcher.ListBlocks(s3.New(awsSession), "near-lake-data-mainnet", 79075963, 100)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(objects[0])
	//result, err := s3Fetcher.FetchStreamerMessage(s3.New(awsSession), "near-lake-data-mainnet", objects[0])
	//if err != nil {
	//	return
	//}
	//fmt.Println(result)
}
