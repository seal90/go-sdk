/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"
	"google.golang.org/grpc/metadata"
)

var (
	// set the environment as instructions.
	pubsubName = "rocketmq-pubsub"
	topicName  = "virtualnamespace"
)

func main() {
	ctx := context.Background()
	publishEventData := []byte("ping33")
	// publishEventsData := []interface{}{"multi-ping", "multi-pong"}

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx = metadata.NewIncomingContext(ctx, map[string][]string{"virtual-namespace": []string{"gray"}})
	// ctx = metadata.AppendToOutgoingContext(ctx, "virtual-namespace", "gray1")

	// Publish a single event
	if err := client.PublishEvent(ctx, pubsubName, topicName, publishEventData); err != nil {
		panic(err)
	}

	// // Publish multiple events
	// if res := client.PublishEvents(ctx, pubsubName, topicName, publishEventsData); res.Error != nil {
	// 	panic(err)
	// }

	fmt.Println("data published")

	fmt.Println("Done (CTRL+C to Exit)")
}
