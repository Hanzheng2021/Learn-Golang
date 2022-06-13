func longestCommonPrefix(strs []string) string {
    //fmt.Println(strs[0])
    //fmt.Println(strs[1])
    var length, n, m int
    length = 0
    n = 200
    m = 0

    //var text [][]string
    for i,v := range strs {
        length = len(v)
        //row[i] := strings.Split(strs[i], "")
        //text = append(text, row[i])
        if length < n {
            n = length
            m = i
        }
    }
    fmt.Println(n)
    fmt.Println(m)

    var shortest []string
    shortest = strings.Split(strs[m], "")
    fmt.Println(shortest)
    
    var str []string 
    for _,v := range strs {
        str = append(str, v[0:n])      
    }
    //str1 := shortest[0:n]
    fmt.Println(str)
    
    
    var k int
    k = 0
    test := make([][]int, m) // 二维切片，m行
    for i := range test {
        test[i] = make([]int, len(str)) // 每一行4列
    }
    var a string
    var stri string
    for i:=0; i<n; i++ {
        for j:=0; j<len(str); j++ {
            stri = str[i]
            a = stri[j:j+1]
            test[i][j] = append(test, a...)
        }

    }
    /*
    // var test int
    //test = 0
    for _,v := range str {
        for j := 1; j < m; {
            if v[j-1:j] != shortest[j-1] {
                k = j-1
                break
            } else{                
                //test = j
                j++
                continue
            }
        }
        //if test == m - 1 {

        //}

        //j++
    }
    */
    fmt.Println(k)
    strreturn := strs[m][0:k]
    return strreturn
}
