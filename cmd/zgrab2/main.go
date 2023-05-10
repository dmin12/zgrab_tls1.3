package main

import (
	"github.com/dmin12/zgrab_tls1.3/bin"
	_ "github.com/dmin12/zgrab_tls1.3/modules"
)

// main wraps the "true" main, bin.ZGrab2Main(), after importing all scan
// modules in ZGrab2.
func main() {
	bin.ZGrab2Main()
}
