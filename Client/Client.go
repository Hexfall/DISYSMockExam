package main

import "flag"

var target = flag.String("target", ":5080", "target")

func main() {
	flag.Parse()
	fe := FrontEnd{}
	fe.Connect(*target)
}
