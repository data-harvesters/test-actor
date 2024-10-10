package main

import (
	"fmt"
	"os"

	"github.com/data-harvesters/goapify"
)

type testInput struct {
	Input string `json:"input"`
}

type output struct {
	Test string `json:"test"`
}

func main() {
	fmt.Println("HELLO FROM TEST ACTOR")

	a := goapify.NewActor(os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"),
		os.Getenv("APIFY_TOKEN"),
		os.Getenv("APIFY_DEFAULT_DATASET_ID"))

	input := new(testInput)

	err := a.Input(input)
	if err != nil {
		panic(err)

	}

	fmt.Println(input.Input)

	a.Output(&output{
		Test: fmt.Sprintf("Hello from actor! %s", input.Input),
	})
}
