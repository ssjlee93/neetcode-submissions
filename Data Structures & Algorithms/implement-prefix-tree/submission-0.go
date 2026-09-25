type TrieNode struct {
    Children map[rune]*TrieNode
    Word bool
}

type PrefixTree struct {
    Root *TrieNode
}

func Constructor() PrefixTree {
    return PrefixTree{&TrieNode{map[rune]*TrieNode{}, false}}
}

func (this *PrefixTree) Insert(word string) {
    curr := this.Root
    for _, r := range word {
        if _, ok := curr.Children[r]; !ok {
            curr.Children[r] = &TrieNode{Children: map[rune]*TrieNode{}}
        }
        curr = curr.Children[r]
    }
    curr.Word = true
}

func (this *PrefixTree) Search(word string) bool {
    curr := this.Root
    for _, r := range word {
        if _, ok := curr.Children[r]; !ok {
            return false
        }
        curr = curr.Children[r]
    }
    return curr.Word
}

func (this *PrefixTree) StartsWith(prefix string) bool {
    curr := this.Root
    for _, r := range prefix {
        if _, ok := curr.Children[r]; !ok {
            return false
        }
        curr = curr.Children[r]
    }
    return true
}
