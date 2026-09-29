import "slices"

type UnionFind struct {
    parents []int
    ranks []int
}

func NewUF(n int) *UnionFind {
    uf := UnionFind{
        parents : make([]int, n, n),
        ranks : make([]int, n, n),
    }
    for i := 0; i < n; i++ {
        uf.parents[i] = i
        uf.ranks[i] = 1
    }
    return &uf
}

func (uf *UnionFind) find(x int) int {
    p := uf.parents[x]
    for p != uf.parents[p] {
        uf.parents[p] = uf.parents[uf.parents[p]]
        p = uf.parents[p]
    }
    return p
}

func (uf *UnionFind) union(x, y int) bool {
    p1, p2 := uf.find(x), uf.find(y)
    if p1 == p2 {
        return false
    }

    if uf.ranks[p1] > uf.ranks[p2] {
        uf.parents[p2] = p1
        uf.ranks[p1] += uf.ranks[p2]
    } else if uf.ranks[p2] > uf.ranks[p1] {
        uf.parents[p1] = p2
        uf.ranks[p2] = uf.ranks[p1]
    } else {
        uf.parents[p1] = p2
        uf.ranks[p2] += uf.ranks[p1]
    }
    return true
}

func accountsMerge(accounts [][]string) [][]string {
    // trie lesson

    // copied video solution

    // initialize UnionFind
    uf := NewUF(len(accounts))
    
    emailToAccount := make(map[string]int) // map email to account idx

    // form disjoin sets
    // loop through accounts
    for i, a := range accounts {
        // loop through emails
        for j := 1; j < len(a); j++ {
            email := a[j]
            // case 1 : email already exists
            if acc, ok := emailToAccount[email]; ok {
                // join the curr idx and logged acc idx
                uf.union(i, acc)
            } else {
                // case 2 : email doesn't exist
                // add to our map
                emailToAccount[email] = i
            }
        }
    }

    emailGroup := make(map[int][]string) // acc idx -> emails

    for email, account := range emailToAccount {
        leader := uf.find(account)
        if _, ok := emailGroup[leader]; !ok {
            emailGroup[leader] = make([]string, 0)
        }
        emailGroup[leader] = append(emailGroup[leader], email)
    }

    // form ans
    ans := make([][]string, 0, len(accounts))
    for idx, emails := range emailGroup {
        name := accounts[idx][0]
        person := make([]string, 1, len(emails)+1)
        person[0] = name
        slices.Sort(emails)
        person = append(person, emails...)
        ans = append(ans, person)
    }
    return ans

}
