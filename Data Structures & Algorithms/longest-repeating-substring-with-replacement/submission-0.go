func characterReplacement(s string, k int) int {
    // sliding window
    // copied solution
    ans := 0
    chars := make(map[byte]int)
    mostf := 0

    l := 0
    for r := 0; r < len(s); r++ {
        // add chars count
        chars[s[r]]++
        mostf = max(mostf, chars[s[r]])
        for (r-l+1) - mostf > k {
            chars[s[l]]--
            l++
        }
        ans = max(ans, r-l+1)
    }
    return ans
}
