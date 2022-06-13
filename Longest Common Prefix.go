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
    for _,v := range str {
        j := 1
        if v[j-1:j] != shortest[j] {
            break
            k = j
        } else{
            //k = j
            continue
        }
        j++
    }
    fmt.Println(k)
    strreturn := strs[m][0:k]
    return strreturn
}
