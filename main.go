package main

import (
	"fmt"
	"stock-processor/internal/consumer"
	"strings"

	ext "github.com/reugn/go-streams/extension"
	"github.com/reugn/go-streams/flow"
)

func main() {
	feeder := consumer.NewFeeder("localhost:9093", "test-client")
	source := ext.NewChanSource(feeder.FeedChannel)

	toUpperMapFlow := flow.NewMap(toUpper, 1)
	for i := range source.Via(toUpperMapFlow).Out() {
		fmt.Println(i)
	}

}

var toUpper = func(msg string) string {
	return strings.ToUpper(msg)
}
