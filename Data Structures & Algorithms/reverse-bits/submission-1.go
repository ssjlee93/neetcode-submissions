func reverseBits(n int) int {
    ans := 0
    for i := 0; i < 32; i++ {
        ans = ans << 1
        if n % 2 == 1 {
            ans++
        }
        n = n >> 1

    }
    return ans
}
