func characterReplacement(s string, k int) int {
    // sliding window
    // copied solution : plain sliding window
    ans := 0
    chars := make(map[byte]int)

    l := 0
    for r := 0; r < len(s); r++ {
        // add chars count
        chars[s[r]]++

        // shrink window until window-maxf <= k
        for (r-l+1) - findMax(chars) > k {
            chars[s[l]]--
            l++
        }
        ans = max(ans, r-l+1)
    }
    return ans
}

func findMax(chars map[byte]int) int {
    // find max value in chars
    maxf := 0
    for _, v := range chars {
        maxf = max(maxf, v)
    }
    return maxf
}
