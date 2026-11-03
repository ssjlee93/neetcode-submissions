class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # form a graph
        graph = {}
        for c, p in prerequisites:
            if c not in graph:
                graph[c] = []
            graph[c].append(p)
        
        # store visited
        visited = set()

        def dfs(node):
            if node in visited:
                return False
            
            if node not in graph:
                return True
            
            visited.add(node)
            # process neighbors
            for nei in graph[node]:
                if not dfs(nei):
                    return False
            visited.remove(node)
            del graph[node]
            return True
        
        for c in range(numCourses):
            if not dfs(c):
                return False
        return True