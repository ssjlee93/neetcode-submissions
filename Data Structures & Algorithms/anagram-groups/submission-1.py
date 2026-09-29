class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # brute force : sort each word and add to hashmap
        anagrams = {}
        
        for word in strs:
            ordered = "".join(sorted(word))
            if ordered not in anagrams:
                anagrams[ordered] = [word]
            else:
                anagrams[ordered].append(word)
        
        
        return [group for group in anagrams.values()]
