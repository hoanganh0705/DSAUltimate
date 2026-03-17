/*
	Count down

Print numbers from n to 1 (no loops allowed)
*/

function countDown(n: number): void {
    if (n <= 0) {
        return;
    }
    console.log(n);
    countDown(n - 1);
}