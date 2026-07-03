func isAnagram(s string, t string) bool {
    // hashmap
    // assumption : 
    // if all chars in s exists in t in equal frequency, 
    // then it can be an anagram

    // edge : if length doesn't match, cannot be an anagram
    if len(s) != len(t) {
        return false
    }

    // count s 
    freqS := make(map[rune]int)
    for _, r := range s {
        freqS[r]++
    }

    // check from t iteration
    for _, r := range t {
        v, ok := freqS[r]
        if !ok {
            return false
        }
        
        if v <= 0 {
            return false
        }

        freqS[r]--
    }
    
    
    return true
}
