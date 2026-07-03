import (
    "slices"
    "cmp"
)

func topKFrequent(nums []int, k int) []int {
    // use hashmap for frequency
    freq := make(map[int]int)
    for _, n := range nums {
        freq[n]++
    }

    // extract values
    values := make([][2]int, 0, len(freq))
    for k, v := range freq {
        pair := [2]int{v, k}
        values = append(values, pair)
    }

    // sort in descending order
    slices.SortFunc(values, func(a, b [2]int) int {
        return cmp.Compare(b[0], a[0])
    })

    // match the value
    ans := make([]int, k, k)
    for i := 0; i < k; i ++ {
        ans[i] = values[i][1]
    }

    return ans
}
