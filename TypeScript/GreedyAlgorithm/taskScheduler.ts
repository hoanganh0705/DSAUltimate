function leastInterval(tasks: string[], n: number): number {
    const count: number[] = new Array(26).fill(0);
    let maxFreq = 0;
    let numberMaxFreq = 0;

    for (const task of tasks) {
        const index = task.charCodeAt(0) - 'A'.charCodeAt(0);
        count[index]++;

        if (maxFreq < count[index]) {
            maxFreq = count[index];
            numberMaxFreq = 1;
        } else if (maxFreq === count[index]) {
            numberMaxFreq++;
        }
    }

    const parts = maxFreq - 1;
    const slotsPerPart = n - (numberMaxFreq - 1);
    const totalSlots = parts * slotsPerPart;
    const remainingTasks = tasks.length - maxFreq * numberMaxFreq;

    const idles = Math.max(0, totalSlots - remainingTasks);

    return tasks.length + idles;
}