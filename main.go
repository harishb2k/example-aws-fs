package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime"
	"sync"
)

func main() {

	mySession := session.Must(session.NewSession())

	// Create a SageMakerFeatureStoreRuntime client with additional configuration
	svc := sagemakerfeaturestoreruntime.New(mySession, aws.NewConfig().WithRegion("ap-south-1"))

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10000; i++ {

				for j := 0; j < len(data); j++ {
					output, err := svc.GetRecord(&sagemakerfeaturestoreruntime.GetRecordInput{
						FeatureGroupName: aws.String("group_name"),
						FeatureNames: []*string{
							aws.String("name_1"),
							aws.String("name_2"),
							aws.String("name_3"),
						},
						RecordIdentifierValueAsString: aws.String(data[j]),
					})
					if err != nil {
						fmt.Println(">>>> Got problem for this feature get")
					} else {
						fmt.Println("Ok", output.String())
					}
				}

			}
			wg.Done()
		}()
	}

	wg.Wait()
}

var data = []string{
	"a",
}
