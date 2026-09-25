// trie lesson
// copied solution

type TrieNode struct {
	Children map[rune]*TrieNode
	Word bool
}

type WordDictionary struct {
    Root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{&TrieNode{ Children: map[rune]*TrieNode{} }}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this.Root
	for _, r := range word {
		if _, ok := curr.Children[r]; !ok {
			curr.Children[r] = &TrieNode{ Children : map[rune]*TrieNode{} }
		}
		curr = curr.Children[r]
	}
	curr.Word = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
    dfs = func(j int, node *TrieNode) bool {
		curr := node

		for i := j; i < len(word); i++ {
			r := rune(word[i])
			if r == '.' {
				for _, child := range curr.Children {
					if dfs(i+1, child) {
						return true
					}
				}
				return false
			} else {
				if _, ok := curr.Children[r]; !ok {
					return false
				}
				curr = curr.Children[r]
			}
		}
		return curr.Word
	}
	return dfs(0, this.Root)
}
