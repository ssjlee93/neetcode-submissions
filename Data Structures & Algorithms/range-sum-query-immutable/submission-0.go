type NumArray struct {
    Prefix []int
}


func Constructor(nums []int) NumArray {
    init := NumArray{[]int{}}
    for i, n := range nums {
        if i == 0 {
            init.Prefix = append(init.Prefix, n)
            continue
        }
        
        init.Prefix = append(init.Prefix, init.Prefix[len(init.Prefix)-1] + n)
    }
    return init
}


func (this *NumArray) SumRange(left int, right int) int {
    prefixSum := this.Prefix[right]
    if left > 0 {
        prefixSum -= this.Prefix[left-1]
    }
    return prefixSum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */