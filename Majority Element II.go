func majorityElement(nums []int) []int {
    sort.Ints(nums)
    fmt.Println(nums)
    var l int
    var x float64
    x = float64(len(nums)/3)
    x = math.Floor(x)
    l = int(x)
    var result []int
    for i:=0; i+l<len(nums);i++  {
            if nums[i]==nums[i+l] {
                result = append(result, nums[i])
                i=i+l
    }
    }

    if len(result)==2&&result[0]==result[1] {
        result = result[0:1]
    }
    return result


}
