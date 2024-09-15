package agent

import "flag"

var IsServerUpdate *bool

func GetArguments() {
	IsServerUpdate = flag.Bool("server", false, "should update server or client")

	flag.Parse()
}
