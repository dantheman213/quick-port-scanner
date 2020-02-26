package scanner

import (
    "fmt"
    "net"
    "strconv"
    "time"
)

var CommonPorts *[]int = &[]int{21, 22, 23, 25, 26, 37, 53, 80, 110, 111, 135, 139, 143, 179, 389, 443, 445, 465, 548, 515, 554, 631, 990, 993, 995, 1025, 1027, 1139, 1433, 1723, 1900, 3306, 3389, 5432, 5800, 5900, 8000, 8008, 8080, 8443}

// TODO: support UDP (must be done differently)
// https://networkengineering.stackexchange.com/questions/26541/how-to-verify-that-a-udp-port-is-open
func ScanIP(host string, ports *[]int, timeout int) {
    expire := time.Millisecond * time.Duration(timeout)

    for _, port := range *ports {
        address := net.JoinHostPort(host, strconv.Itoa(port))
        conn, err := net.DialTimeout("tcp", address, expire)
        if err != nil {
            fmt.Print(".")
        }
        if conn != nil {
            defer conn.Close()
            fmt.Printf("\nPort opened: %s\n", address)
        }
    }
}

func MakePortRangeArray(start, end int) *[]int {
    arr := make([]int, end - start)
    j := -1
    for k := 0; k < (end - start); k++ {
        if j == -1 {
            j = start
        } else {
            j += 1
        }
        arr[k] = j
    }

    return &arr
}
