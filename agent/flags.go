package agent

import "flag"

var IsServerUpdate *bool
var StandaloneUpdate *bool

func GetArguments() {
	IsServerUpdate = flag.Bool("server", false, "should update server or client")
	StandaloneUpdate = flag.Bool("standalone", false, "update game without prior run")

	flag.Parse()
}
