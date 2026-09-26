type TrieNode struct {
    children map[byte]*TrieNode
    isWord bool
}

func (trie *TrieNode) addWord(word string) {
    curr := trie
    for i := 0; i < len(word); i++ {
        c := word[i]
        if _, ok := curr.children[c]; !ok {
            curr.children[c] = &TrieNode{ children : make(map[byte]*TrieNode)}
        }
        curr = curr.children[c]
    }
    curr.isWord = true
}

func findWords(board [][]byte, words []string) []string {
    // form trie from words
    root := &TrieNode{children : make(map[byte]*TrieNode)}
    for _, w := range words {
        root.addWord(w)
    }

    // define constants
    ROWS, COLS := len(board), len(board[0])
    found := make(map[string]bool)
    visit := make(map[[2]int]bool)

    // define dfs
    var dfs func(r, c int, node *TrieNode, word string)
    dfs = func(r, c int, node *TrieNode, word string) {
        point := [2]int{r, c}
        if r < 0 || c < 0 || r >= ROWS || c >= COLS || visit[point] {
            return
        }
        // base case : check if our children has curr char
        char := board[r][c]
        next, ok := node.children[char]
        if !ok {
            return
        }

        // mark visited
        visit[point] = true
        // add to current prefix we found so far
        word += string(char)

        if next.isWord {
            found[word] = true
        }

        dfs(r+1, c, next, word)
        dfs(r-1, c, next, word)
        dfs(r, c+1, next, word)
        dfs(r, c-1, next, word)
        // backtrack
        visit[point] = false
    }

    // scan grid
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            dfs(r, c, root, "")
        }
    }

    // form ans
    ans := make([]string, 0, len(found))
    for k, _ := range found {
        ans = append(ans, k)
    }

    return ans
}
