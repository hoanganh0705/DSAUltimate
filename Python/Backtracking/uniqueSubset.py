def subsetsWithDup(nums):
    #write code here
    res = []
    def helper(i,curr):
        #base
        if i == len(nums):
            res.append(curr[:])
            return
        #include
        curr.append(nums[i])
        helper(i+1,curr)
        curr.pop() #backtracking
        
        #exclude
        while i<len(nums)-1 and nums[i]==nums[i+1]:
            i+=1
        helper(i+1) 
        
    helper(0,[])
    return res