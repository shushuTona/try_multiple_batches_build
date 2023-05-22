package main

import (
	"try_multiple_batches_build/batch-b/gateway"
	"try_multiple_batches_build/common"
)

func main() {
	name := gateway.GetName()

	common.Hello(name)
}
