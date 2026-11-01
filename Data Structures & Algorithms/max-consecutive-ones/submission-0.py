class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        # ans = 0 # record max consecutive 1s
        # keep track of curr consecutive
        l  = 0 
        r = 0
        count = 0
        while r < len(nums):
            if nums[l] == 1 and  nums[r] == 1:
                count = max(count, r-l+1)
                r += 1
            elif nums[r] == 0:
                r += 1
                l = r
        count = max(count, r-l)
        return count