
func isValid(s string) bool {
    m := map[rune]rune{'(':')', '{':'}', '[':']'}
    // edge case : uneven num == cann't fulfill pairs
    if len(s) % 2 != 0 {
        return false
    }
    // edge case : s begins with closing bracket
    if _, ok := m[rune(s[0])]; !ok {
        return false
    }

    // keep a stack of chars
    stack := make([]rune, 0)
    
    // loop through and peek stack in each time
    for _, r := range s {
        // case 1 : open bracket - always goes in
        // case 2 : closing bracket - check with stack
        if _, ok := m[r]; ok {
            stack = append(stack, r)
            continue
        }
        // no opening bracket
        if len(stack) == 0 {
            return false
        }
        // pull out pair to compare
        pair := m[stack[len(stack)-1]]
        if r != pair {
            return false
        }
        // pop if r matches pair of top of stack
        stack = stack[:len(stack)-1]
    }
    return len(stack) == 0 // checks against all opening brackets in stack
}
