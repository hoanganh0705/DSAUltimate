def leastInterval(tasks, n):
    # Đếm tần suất của 26 chữ cái A-Z
    count = [0] * 26
    max_freq = 0
    number_max_freq = 0

    for task in tasks:
        index = ord(task) - ord('A')
        count[index] += 1

        if max_freq < count[index]:
            max_freq = count[index]
            number_max_freq = 1
        elif max_freq == count[index]:
            number_max_freq += 1

    # Tính toán số khoảng trống
    parts = max_freq - 1
    slots_per_part = n - (number_max_freq - 1)
    total_slots_in_parts = parts * slots_per_part
    tasks_remaining = len(tasks) - max_freq * number_max_freq

    idles = max(0, total_slots_in_parts - tasks_remaining)

    return len(tasks) + idles