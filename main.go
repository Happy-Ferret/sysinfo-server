package main

import "flag"

var (
	netPort = flag.String("netPort", "9000", "The sysinfo-server netPort.")
	proto   = flag.String("proto", "tcp", "UDP or TCP.")
	webPort = flag.String("webPort", "8088", "The sysinfo-server webPort.")
	a       = App{}
)

func init() {
	a.Initialize("sysinfo", "sysinfo", "sysinfo", "127.0.0.1")
	ensureTableExists()
	flag.Parse()
}
func main() {
	go startNetServer()
	a.Run(":" + *webPort)
}
