# 121. Best Time to Buy and Sell Stock

## Problem Understanding

Given an array `prices`, where `prices[i]` represents the stock price on day `i`, determine the maximum profit that can be achieved by buying once and selling once.

Rules:

- Buy before you sell.
- You may complete at most one transaction.
- If no profit is possible, return `0`.

Example:

```text
Input:
[7,1,5,3,6,4]

Output:
5

Buy at 1
Sell at 6
Profit = 5
```

---

## Brute Force Approach

### Idea

Try every possible buy day.

For each buy day, try every possible sell day after it.

Calculate the profit and keep track of the maximum.

```go
for sell := 1; sell < len(prices); sell++ {
    for buy := 0; buy < sell; buy++ {
        profit := prices[sell] - prices[buy]
    }
}
```

---

## Complexity

**Time Complexity:** `O(N²)`

- Every pair of buy/sell days is examined.

**Space Complexity:** `O(1)`

---

## Why It Is Inefficient

For every selling day, we repeatedly compare against all previous buying days.

Example:

```text
Sell Day 5

↓

Compare with

Day 1

Day 2

Day 3

Day 4
```

The minimum price seen so far is recalculated repeatedly.

---

## Pattern

**Running Minimum**

Maintain the minimum stock price seen so far while scanning the array once.

---

## Key Invariant

> **At every index, `minPrice` stores the minimum stock price seen before the current day.**

This means:

- We already know the best day to buy.
- We only need to decide whether selling today gives a better profit.

---

## Approach

For every price:

1. Calculate the profit if we sell today.

```text
profit = currentPrice - minPrice
```

2. Update the maximum profit.

3. Update the running minimum price if today's price is lower.

---

## Optimized Solution

```go
func maxProfit(prices []int) int {
    minPrice := prices[0]
    maxProfit := 0

    for i := 1; i < len(prices); i++ {

        profit := prices[i] - minPrice

        if profit > maxProfit {
            maxProfit = profit
        }

        if prices[i] < minPrice {
            minPrice = prices[i]
        }
    }

    return maxProfit
}
```

---

## Complexity

**Time Complexity:** `O(N)`

- Every price is processed exactly once.

**Space Complexity:** `O(1)`

- Only two variables are maintained.

---

## Why This Works

At each day:

- `minPrice` already represents the best buying opportunity seen so far.
- The only remaining decision is whether selling today yields a better profit.

Instead of comparing today's price with every previous day, we compare it only with the minimum price encountered so far.

---

## Mistakes / Learnings

- We don't need to compare every previous price.
- The only previous value that matters is the **minimum price seen so far**.
- Calculate profit **before** updating `minPrice`.
- If today's price becomes the new minimum, it can only be used for future selling days.

---

## Mental Model

> **For every day, ask only one question:**

```text
"If I sell today,

what is the cheapest price

I could have bought at earlier?"
```

That cheapest price is always stored in `minPrice`.