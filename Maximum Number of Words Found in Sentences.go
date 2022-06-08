func mostWordsFound(sentences []string) int {
    var ss [] string
    var n [100] (int)
    var slen int
    var maxn int
    maxn = 0
    slen = len(sentences)
    fmt.Println(slen)
    for i:=0; i < slen; i++ {
        ss = strings.Fields(strings.TrimSpace(sentences[i]))
        fmt.Println(ss)
        //fmt.Println(ss)
        n[i] = len(ss)
        //fmt.Println(n[i])
        if n[i] >= maxn {
            maxn = n[i]
        }
    }
    return maxn
}
