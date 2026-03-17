'''
	Count down

Print numbers from n to 1 (no loops allowed)
'''
def count_down(n):
    if n <= 0:
        return
    print(n)
    count_down(n-1)