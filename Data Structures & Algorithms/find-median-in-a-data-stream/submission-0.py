import heapq

class MedianFinder:
    # based on Video solution

    def __init__(self):
        # two heaps
        self.small = [] # maxheap of nums below median
        self.large = [] # minheap of nums above median
        # idea : by keeping two heaps, we can sort in log n time per insert
        # median retrieval : O(1) = root of max heap and root of min heap

    def addNum(self, num: int) -> None:
        # push to samll and see if we need to move it
        heapq.heappush_max(self.small, num)
        if self.small and self.large and self.small[0] > self.large[0]:
            val = heapq.heappop_max(self.small)
            heapq.heappush(self.large, val)
        
        # balance the two heaps
        if len(self.small) > len(self.large) + 1:
            val = heapq.heappop_max(self.small)
            heapq.heappush(self.large, val)
        elif len(self.large) > len(self.small) + 1:
            val = heapq.heappop(self.large)
            heapq.heappush_max(self.small, val)


    def findMedian(self) -> float:
        if len(self.small) > len(self.large):
            return self.small[0]
        elif len(self.large) > len(self.small):
            return self.large[0]
        
        return (self.small[0] + self.large[0]) / 2
        
        