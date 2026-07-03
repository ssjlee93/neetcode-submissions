class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # hahs map
        seen = defaultdict(int) # only one solution guaranteed.  

        for i, n in enumerate(nums): 
            # print(f"i: {i} n: {n}")
            complement = target - n

            if complement in seen:
                print([seen[complement], i])
                return [seen[complement], i]
            
            seen[n] = i
        # no need for other returns.
        # solution guaranteed. 
        return