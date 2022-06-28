func kidsWithCandies(candies []int, extraCandies int) []bool {
    var result []bool
    maxcandies := 0
    for _,v := range candies {
        if v>maxcandies {
            maxcandies = v
        }
    }

    for _,v := range candies {
        if v+extraCandies>=maxcandies {
            result = append(result,true) 
        } else {
            result = append(result,false)
        }
    }

    return result
}
