package main

import "github.com/gogo/protobuf/version"

func main() {
	print(version.AtLeast("v1.2.3"))
}
