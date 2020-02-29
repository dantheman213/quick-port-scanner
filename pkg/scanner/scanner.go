package scanner

import (
    "fmt"
    "net"
    "strconv"
    "time"
)

// TODO: support UDP (must be done differently)
// https://networkengineering.stackexchange.com/questions/26541/how-to-verify-that-a-udp-port-is-open
func ScanIP(host string, ports *[]int, timeout int) {
    fmt.Println("Scanning ports...")
    expire := time.Millisecond * time.Duration(timeout)

    for _, port := range *ports {
        address := net.JoinHostPort(host, strconv.Itoa(port))
        conn, err := net.DialTimeout("tcp", address, expire)
        if err != nil {
            // TODO
        }
        if conn != nil {
            defer conn.Close()
            fmt.Printf("OPEN PORT: %s %s\n", address, Ports[port])
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
