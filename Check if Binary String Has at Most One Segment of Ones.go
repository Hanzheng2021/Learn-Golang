func checkOnesSegment(s string) bool {
    if s=="1"||s=="10" {
        return true
    }
    ss := strings.Split(s,"0")
    fmt.Println(ss)
    k := 0
    for _,v := range ss {
        if len(v)>=2 {
            k++
        }
    }

    if k == 1|| ss == ["1"] {
        return true
    } else {
        return false
    }

}



func checkOnesSegment(s string) bool {
    var sint int
    var sstr string  
    if s=="1" {
        return true
    } else {
        sstr = s[1:]
        sint = int(str)
        if sint ==0 {
            return true
        }
    }
 
    ss := strings.Split(s,"0")
    fmt.Println(ss)
    k := 0
    for _,v := range ss {
        if len(v)>=2 {
            k++
        }
    }

    if k == 1 {
        return true
    } else {
        return false
    }

}
