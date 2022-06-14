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
    fmt.Println(len(strs))

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
    /*
    test := make([][]string, len(strs)) // 二维切片，xx行
    for i := range test {
        test[i] = make([]string, len(shortest)) // 每一行xx列
    }
*/
    var test [][]string
    var a []string
    for _,v := range str {
        a = strings.Split(v, "")
        test = append(test, a)
    }
    fmt.Println(test)
//    fmt.Println(test[1][3])
    

    for i:=0; i<n-1; i++ {
        b := 0
        for j:=0; j<len(str); j++ {
            if test[j][i] == shortest[i] {                
                b++
            }
        }
        fmt.Println(b)
        if b==len(str) {
            k++
        }
    }

    fmt.Println(k)
    var strreturn string
    //strreturn := strs[m][0:k]
    if k==1 {
        strreturn = shortest[0]
    } else {
        strreturn = strs[m][0:k]
    }
    
    if len(strs)==1 {
        return strs[0]
    } else {
        return strreturn
    }

    //return strreturn
    
}
