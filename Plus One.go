func plusOne(digits []int) []int {
    for i:=len(digits)-1; i>=0;i-- {
        if digits[i] == 9 {
            digits[i] = 0
            continue
        } else {
            digits[i]++
            break
        }
    
    }

    var s []int
    for _,v := range digits {
        s = append(s, v)
    }
    sort.Ints(s)
    fmt.Println(digits)

    if s[len(s)-1]==0 {
        arr1 := []int {1}
        arr2 := append(arr1,digits...)
        return arr2 
    }
    return digits
}
