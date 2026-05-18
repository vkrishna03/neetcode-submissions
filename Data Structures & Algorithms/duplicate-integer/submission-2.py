class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        nums_tup = tuple(nums)
        map = dict()
        map.fromkeys(nums_tup, 0)
        for i in nums:
            if i in map:
                map[i] += 1
            else:
                map[i] = 1
        
        for k,v in map.items():
            if v > 1:
                return True
        return False        