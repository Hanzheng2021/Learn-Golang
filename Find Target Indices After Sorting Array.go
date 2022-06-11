func targetIndices(nums []int, target int) []int {
    sort.Ints(nums)
    fmt.Println(nums)
    var numbers []int
    for i, v := range nums {
        if v == target {
            numbers = append(numbers, i)
        }
    }
    return numbers
}
