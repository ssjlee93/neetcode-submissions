import "slices"

func groupAnagrams(strs []string) [][]string {
    // sort each string
    // use sorted string as key of hashmap
    // add elements to it
    seen := make(map[string][]string)

    for _, s := range strs {
        
        rs := []rune(s)
        slices.Sort(rs)
        sorted_s := string(rs)
        
        if _, ok := seen[sorted_s]; !ok {
            seen[sorted_s] = []string{s}
        } else {
            seen[sorted_s] = append(seen[sorted_s], s)
        }
    }

    ans := make([][]string, 0, len(seen))
    
    for _, v := range seen {
        ans = append(ans, v)
    }

    return ans
}
