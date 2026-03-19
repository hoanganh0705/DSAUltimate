function fractionalKnapsack(
  W: number,
  arr: [number, number][],
  n: number
): number {

  // sort by value/weight ratio (descending)
  arr.sort((a, b) => (b[0] / b[1]) - (a[0] / a[1]));

  let remainingWeight: number = W;
  let value: number = 0;
  n = arr.length;

  for (let i = 0; i < n; i++) {
    if (remainingWeight === 0) break;

    const weight: number = Math.min(remainingWeight, arr[i][1]);
    remainingWeight -= weight;

    value += (arr[i][0] / arr[i][1]) * weight;
  }

  return value;
}