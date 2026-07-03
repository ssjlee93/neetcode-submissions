class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> seen = new HashMap<>();
        
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            int iComplement = seen.getOrDefault(complement, -1);

            if (iComplement != -1) {
                int[] ans = new int[2];
                ans[0] = iComplement;
                ans[1] = i;
                return ans;
            } 
            seen.put(nums[i], i);
        }
        return null;
    }
}
