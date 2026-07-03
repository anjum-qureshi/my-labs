# 169. Majority Element

## 1. Problem Understanding
The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

---

## 2. Brute Force Approach
Find the frequencies of element and find the max occurred frequency

```go
func majorityElement(nums []int) int {
    for i := 0; i < len(nums); i++ {
        count := 0
        for j := 0; j < len(nums); j++ {
            if i == j {
                continue;
            }
            if nums[i] == nums[j] {
                count++
            }
        }
        if count >= len(nums)/2 {
            return nums[i]
        }
    }
    return -1
}
```

Time Complexity: O(n2)
Space Complexity: O(1)

3. Why it is inefficient
Repeated check of the number occurrence
Two iterations over same array to get the frequency

4. Pattern Identification
The problem guarantees one element appears more than n/2 times.

This suggests:
1. frequency counting (HashMap)
2. or exploiting the majority property directly

Boyer-Moore works by canceling out different elements.
Since the majority element appears more than all others combined,
it survives the cancellation process.
HASHMAP ---->  Boyer-Moore

What pattern does this look like?
Hash Map

Why this pattern?
Because the info that a majority would be greater than n/2 which means every other elements is less than n/2.
A majority element survives pairwise cancellation.
If we remove one majority element and one non-majority element repeatedly,
the majority element can never be fully eliminated because it appears more than n/2 times.

Boyer-Moore simulates this cancellation process using a counter.

5. Optimized Solution (with hashMap)

```go
func majorityElement(nums []int) int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }

    largest := nums[0]
    for ele, f := range freq {
        if f > freq[largest] {
            largest = ele
        }
    }
    return largest
}```

With Boyer-Moore

```go
func majorityElement(nums []int) int {
    candidate,count := 0, 0
    for _,num := range nums {
        if count == 0 {
            candidate = num
        }
        if candidate == num {
            count++
        } else {
            count--
        }
    }

 return candidate
}
```

6. Key Insight

One or two lines:
Some optimal algorithms are not discovered by incremental optimization.

Boyer-Moore comes from reasoning about the mathematical property
of a majority element rather than improving frequency counting.

Need to learn deep on the Boyer-Moore algoritm

7. Mistake / Learning
You don't need to read the whole frequency.
Knows and constraints info in the problem are clue to optimal solutions
Thought HashMap was best optimal
