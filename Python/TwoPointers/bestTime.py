def maxProfit(prices):
    left = 0  # ngày mua (giá thấp nhất hiện tại)
    profit = 0

    for right in range(1, len(prices)):  # ngày bán
        if prices[right] < prices[left]:
            left = right  # tìm được giá thấp hơn -> cập nhật ngày mua
        else:
            profit = max(profit, prices[right] - prices[left])

    return profit