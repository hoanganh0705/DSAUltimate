'''
Coding Exercise(Tower of Hanoi)
The tower of Hanoi is a famous puzzle where we have three rods and N disks. The objective of the puzzle is to move the entire stack to another rod.
You are given the number of discs N. Initially, these discs are in the rod 1.
You need to print all the steps of discs movement so that all the discs reach the 3rd rod. Also, you need to find the total moves.



Note: The discs are arranged such that the top disc is numbered 1 and the bottom-most disc is numbered N.
Also, all the discs have different sizes and a bigger disc cannot be put on the top of a smaller disc.
Refer the provided link to get a better clarity about the puzzle.



EXAMPLE:
For Input (N = 2)
Output:
move disk 1 from rod 1 to rod 2
move disk 2 from rod 1 to rod 3
move disk 1 from rod 2 to rod 3
3

'''

def toh(n,  fromm, to, aux):
    count = 0 
    def helper(n, fromm, to, aux):
        nonlocal count
        #base case 
        if n==1:
            print(f"move disk 1 from rod {fromm} to rod {to}")
            count += 1
            return
        #recursive case
        helper(n-1, fromm, aux, to)
        print(f"move disk {n} from rod {fromm} to rod {to}")
        count += 1
        helper(n-1, aux, to, fromm)  
    helper(n, fromm, to, aux)
    return count

n = int(input())
res = toh(n, 1, 3, 2)
print(res)