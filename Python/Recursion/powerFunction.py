'''
Power function

Compute xⁿ recursively

Example: power(2, 5) → 32

'''

def power_func(base, exponent):
    if exponent == 0:
        return 1
    return power_func(base, exponent - 1) * base
