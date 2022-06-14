func checkOnesSegment(s string) bool { 
    ss := strings.Split(s,"0")
    fmt.Println(ss)

    k := 0
    for _,v := range ss {
        if len(v)>=1 {
            k++
        }
    }

    if k == 1 {
        return true
    } else {
        return false
    }

}
