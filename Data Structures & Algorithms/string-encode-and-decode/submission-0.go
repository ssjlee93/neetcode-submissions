type Solution struct{}

var secret rune = rune(256)

func (s *Solution) Encode(strs []string) string {
    // use a very simple byte shift.
    // since the input limit is 256, let's shift each byte of a char by 256.
    encoded := ""
    for i := range len(strs) {
        s := strs[i]
        encoded += s + string(secret)
    }
    return encoded
}

func (s *Solution) Decode(encoded string) []string {
    decoded := make([]string, 0)
    curr := ""
    for _, r := range encoded {
        if r != secret {
            curr += string(r)
        } else {
            decoded = append(decoded, curr)
            curr = ""
            continue
        }
    }
    return decoded
}
