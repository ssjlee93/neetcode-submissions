type SegmentTree struct {
    left *SegmentTree
    right *SegmentTree
    l int
    r int
    sum int
}

func NewSegmentTree(nums []int) *SegmentTree {
    st := build(0, len(nums)-1, nums)
    return st
}

func build(L, R int, nums []int) *SegmentTree {
    // base case : we reached a leaf node which has L == R
    if L == R {
        return &SegmentTree{l : L, r : R, sum : nums[L]}
    }

    M := (L + R) / 2
    root := &SegmentTree{sum : 0, l : L, r : R}
    root.left = build(L, M, nums)
    root.right = build(M+1, R, nums)
    root.sum = root.left.sum + root.right.sum
    return root
}

func (st *SegmentTree) Update(index, val int) {
    // base case : l and r index is the same
    if st.l == st.r {
        st.sum = val
        return
    }

    M := (st.l + st.r) / 2
    if index > M {
        st.right.Update(index, val)
    } else {
        st.left.Update(index, val)
    }
    st.sum = st.left.sum + st.right.sum
}

func (st *SegmentTree) Query(L, R int) int {
    // base case : we reached node that matches our range
    if L == st.l && st.r == R {
        return st.sum
    }

    M := (st.l + st.r) / 2
    // 3 cases
    if L > M {
        // query window all fits in right tree
        return st.right.Query(L, R)
    } else if R <= M {
        // query window all fits in left tree
        return st.left.Query(L, R)
    } else {
        // query window lies between left and right tree
        return st.left.Query(L, M) + st.right.Query(M+1, R)
    }
}
