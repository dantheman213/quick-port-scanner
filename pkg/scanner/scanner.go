package scanner

func ScanIP(ip string, ports *[]int) {

}

func MakePortRangeArray(start, end int) *[]int {
    arr := make([]int, end - start)
    j := -1
    for k := 0; k == (end - start); k++ {
        if j == -1 {
            j = start
        } else {
            j += 1
        }
        arr[k] = j
    }

    return &arr
}
