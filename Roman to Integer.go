func romanToInt(s string) int {
    var ss [] string
    var n, sslen, m int
    ss = strings.Split(s, "")
    sslen = len(ss)
    n = 0
    fmt.Println(ss)
    fmt.Println(sslen)
    for i := sslen - 1; i > 0; i-- {
        switch {
            case ss[i] == "I":
            n= n + 1
            case ss[i] == "V"&&ss[i-1] != "I":
            n= n + 5
            case ss[i] == "V"&&ss[i-1] == "I":
            n= n + 4
            i--
            case ss[i] == "X"&&ss[i-1] != "I":
            n= n + 10
            case ss[i] == "X"&&ss[i-1] == "I":
            n= n + 9
            i--
            case ss[i] == "L"&&ss[i-1] != "X":
            n= n + 50
            case ss[i] == "L"&&ss[i-1] == "X":
            n= n + 40
            i--
            case ss[i] == "C"&&ss[i-1] != "X":
            n= n + 100
            case ss[i] == "C"&&ss[i-1] == "X":
            n= n + 90
            i--
            case ss[i] == "D"&&ss[i-1] != "C":
            n= n + 500
            case ss[i] == "D"&&ss[i-1] == "C":
            n= n + 400
            i--
            case ss[i] == "M"&&ss[i-1] != "C":
            n= n + 1000
            case ss[i] == "M"&&ss[i-1] == "C":
            n= n + 900
            i--
        }
        m = i
        //fmt.Println(m)
    }
    //m == i
    //fmt.Println(m)
    // i= -1 or 0
    if m == 1 || sslen ==1 {
        switch {
            case ss[0] == "I":
            n= n + 1
            case ss[0] == "V":
            n= n + 5
            case ss[0] == "X" :
            n= n + 10
            case ss[0] == "L" :
            n= n + 50
            case ss[0] == "C" :
            n= n + 100
            case ss[0] == "D" :
            n= n + 500
            case ss[0] == "M" :
            n= n + 1000

    }

    }
    
    fmt.Println(n)
    return n

}

/*
    switch {
        case ss[0] == "I":
        n=1
        case ss[0] == "V"&&ss[1] != "I":
        n=5
        case ss[0] == "V"&&ss[1] == "I":
        n=4
        //fmt.Println(n)
        case ss[0] == "X"&&ss[1] != "I":
        n=10
        case ss[0] == "X"&&ss[1] == "I":
        n=9
        case ss[0] == "L"&&ss[1] != "X":
        n=50
        case ss[0] == "L"&&ss[1] == "X":
        n=40
        case ss[0] == "C"&&ss[1] != "X":
        n=100
        case ss[0] == "C"&&ss[1] == "X":
        n=90
        case ss[0] == "D"&&ss[1] != "C":
        n=500
        case ss[0] == "D"&&ss[1] != "C":
        n=400
        case ss[0] == "M"&&ss[1] != "C":
        n=1000
        case ss[0] == "M"&&ss[1] == "C":
        n=900
*/
