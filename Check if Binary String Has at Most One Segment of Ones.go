func checkOnesSegment(s string) bool { 
    if s=="1" {
        return true
    }
    sss := strings.Split(s,"")

    j:=0
    for i:=1; i<len(sss); i++{
        if sss[0]=="1"&&sss[i]=="0"{
            j++
        }
    }
    fmt.Println(j)
    if j==len(sss)-1 {
        return true
    }
    ss := strings.Split(s,"0")
    fmt.Println(ss)


    k := 0
    for _,v := range ss {
        if len(v)>=1 {
            k++
        }
    }
    fmt.Println(k)
    if k == 1 && ss[0]!= "1" && ss[len(ss)-1]!= "1" {
        return true
    } else {
        return false
    }

}
