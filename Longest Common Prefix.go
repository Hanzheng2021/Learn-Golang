func longestCommonPrefix(strs []string) string {
    //fmt.Println(strs[0])
    //fmt.Println(strs[1])
    var length, n, m int
    length = 0
    n = 200
    m = 0

    var text [][]string
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
    /*
    for i,v := range strs {
        length = len(v)
        row := strings.Split(strs[i], "")
        for j := n-1; j>0; j--{
            if shortest[j] != row[j] {
                n = n-1
            }
        } 

        }
        text = append(text, row)

    }
*/
    return strs[0]
}
