class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        ans = 0 # record max consecutive 1s
        # keep track of curr consecutive
        # when we see a 0, we update ans and reset curr
        curr = 0
        for num in nums:
            if num == 0:
                curr = 0
                continue
            curr += 1
            ans = max(ans, curr)
        return ans
