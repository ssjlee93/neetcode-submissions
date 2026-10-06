class UnionFind:
    def __init__(self, nums: List[int]):
        self.parents = {}
        self.ranks = {}
        self.longest = 1

        for n in nums:
            self.parents[n] = n
            self.ranks[n] = 1
    

    def find(self, x: int):
        p = self.parents[x]
        while p != self.parents[p]:
            self.parents[p] = self.parents[self.parents[p]]
            p = self.parents[p]
        return p
    

    def union(self, x, y):
        p1, p2 = self.find(x), self.find(y)
        if p1 == p2:
            return False
        
        if self.ranks[p1] > self.ranks[p2]:
            self.parents[p2] = p1
            self.ranks[p1] += self.ranks[p2]
        elif self.ranks[p2] > self.ranks[p2]:
            self.parents[p1] = p2
            self.ranks[p2] += self.ranks[p1]
        else:
            self.parents[p1] = p2
            self.ranks[p2] += self.ranks[p1]
        self.longest = max(self.longest, self.ranks[p1], self.ranks[p2])
        return True
    

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        # python version
        n = len(nums)
        # edge case : nums too short
        if n <= 1:
            return n
        
        union_find = UnionFind(nums)

        # loop and cur num - 1 joins union
        for curr in nums:
            if curr-1 not in union_find.parents:
                continue
            union_find.union(curr, curr-1)

        return union_find.longest
        