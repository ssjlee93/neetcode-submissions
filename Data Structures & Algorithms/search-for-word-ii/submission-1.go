type TrieNode struct {
    children map[byte]*TrieNode
    word bool
}

type Trie struct {
    root *TrieNode
}

func Constructor() Trie {
    return Trie{&TrieNode{map[byte]*TrieNode{}, false}}
}

func (trie *Trie) AddWords(words []string) {
    for _, w := range words {
        curr := trie.root
        for i := 0; i < len(w); i++ {
            c := w[i]
            if curr.children[c] == nil {
                curr.children[c] = &TrieNode{children : make(map[byte]*TrieNode)}
            }
            curr = curr.children[c]
        }
        curr.word = true
    } 
}

func findWords(board [][]byte, words []string) []string {
    // trie lesson
    // copied AI solution
    
    // build trie from the words
    // as we walk through the trie, we perform DFS on the grid

    trie := Constructor()
    trie.AddWords(words)

    ROWS, COLS := len(board), len(board[0])
    ans := make([]string, 0)
    visit := make([][]bool, ROWS)

    for i := range visit {
        visit[i] = make([]bool, COLS)
    }

    var dfs func(r, c int, node *TrieNode, word string)
    dfs = func(r, c int, node *TrieNode, word string) {
        // 1. boundary check
        if r < 0 || c < 0 || r >= ROWS || c >= COLS || visit[r][c] {
            return
        }
        // 2. dead end check
        char := board[r][c]
        next, ok := node.children[char]
        if !ok {
            return
        }
        // 3. extend
        word += string(char)
        if next.word {
            ans = append(ans, word)
            next.word = false // avoid recording duplicates
        }
        // 4. recurse
        visit[r][c] = true
        dfs(r+1, c, next, word)
        dfs(r-1, c, next, word)
        dfs(r, c+1, next, word)
        dfs(r, c-1, next, word)
        visit[r][c] = false // backtrack
    }

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            dfs(r, c, trie.root, "")
        }
    }
    return ans
}
