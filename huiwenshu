func isPalindrome(x int) bool {
    var y,z, input int
    var i,m float64
    var n int
    z = 0
    input = x

    for i = 1; x > 0; i++ {
        m = math.Pow(10,i-1)
        n = int(m)
        y = x % 10// last
        x = (x - y)/10//deldete the last
        z = z*10 + y//new number
        fmt.Println(n, y, x, z)
    }
    if z == input {
        return true
    } else {
        return false
    }

}

//func mi(a,b int) int{//a de b cifang
//    for i := 0; i < b; i++ {
//        a = a*a
//}
