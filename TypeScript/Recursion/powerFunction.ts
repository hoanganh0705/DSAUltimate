/*
Power function

Compute xⁿ recursively

Example: power(2, 5) → 32

*/

const powerFunc = (base: number, exponent: number): number => {
  if (exponent === 0) {
    return 1;
  }
  return powerFunc(base, exponent - 1) * base;
}