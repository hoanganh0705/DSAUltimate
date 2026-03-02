def combine(n,k):
    # write code here
    res = []
    def helper(start,curr):
        # base
        if len(curr) == k:
            res.append(curr[:])
            return
        for j in range(start,n+1):
            curr.append(j)
            helper(j+1,curr)
            curr.pop()

    helper(1,[])
    return res