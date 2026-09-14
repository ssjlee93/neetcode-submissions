func isPalindrome(s string) bool {
    // two pointer
    L, R := 0, len(s)-1
    
    for L < R {
        // must skip all non-alphanumeric chars
        left := s[L]
        right := s[R]
        if !isAlphanumeric(left) {
            L++
            continue
        }
        if !isAlphanumeric(right) {
            R--
            continue
        }
        // must handle case
        if 'A' <= left && left <= 'Z' {
            left += 32
        }
        if 'A' <= right && right <= 'Z' {
            right += 32
        }
        // palindrome check
        if left != right {
            return false
        }
        L++
        R--
    }
    return true
}

func isAlphanumeric(b byte) bool {
    return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}
