package controller

import (
	"flag"
	"fmt"
)

type Options struct {
	CommandQuickScanIP *bool
	CommandNormalScanIP *bool
	IPAddress *string
	Help *bool
	Timeout *int
	Verbose *bool
}

func parseOptions() *Options {
	o := Options{}

	o.CommandQuickScanIP = flag.Bool("a", false, "RUN ACTION - Quick Scan Single IP Address.")
	o.CommandNormalScanIP = flag.Bool("b", false, "RUN ACTION - Normal Scan Single IP Address.")

	o.IPAddress = flag.String("ip", "", "IP address for action(s).")
	o.Help = flag.Bool ("help", false, "Print help sheet.")
	o.Timeout = flag.Int("timeout", 60, "Set the timeout (milleseconds) before disconnecting on error or inactivity.")
	o.Verbose = flag.Bool("verbose", false, "Extra information provided in standard out.")

	flag.Parse()
	return &o
}

func printHelpSheet() {
	fmt.Println("GPS Atlas / quick-port-scanner")
	fmt.Println("Auto-detect, plot, and map with common GPS USB serial devices")
	fmt.Print("\nARGUMENTS:\n\n")
	flag.PrintDefaults()
}

func checkOptionSanity(o *Options) error {

	return nil
}
