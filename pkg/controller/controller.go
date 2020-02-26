package controller

import (
    "fmt"
    "github.com/dantheman213/quick-port-scanner/pkg/scanner"
    "os"
)

func Start() {
    o := parseOptions()
    if len(os.Args) == 1 || *o.Help == true {
        printHelpSheet()
        os.Exit(0)
    }

    if err := checkOptionSanity(o); err != nil {
        fmt.Printf("Parse exception: %s", err)
        os.Exit(1)
    }

    if *o.CommandQuickScanIP {
        scanner.ScanIP(*o.IPAddress, scanner.MakePortRangeArray(1, 1024))
    }
}
