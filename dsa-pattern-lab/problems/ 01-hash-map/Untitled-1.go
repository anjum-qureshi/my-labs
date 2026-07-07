// 347. Top K Frequent Elements

type Pair struct {
	Num   int
	Count int
}

//. brute force
func topKFrequent(nums []int, k int) []int {
	result := []Pair{}

	for _, num := range nums {
		// Check if already processed
		seen := false
		for _, p := range result {
			if p.Num == num {
				seen = true
				break
			}
		}

		if seen {
			continue
		}

		// Count frequency
		count := 0
		for _, x := range nums {
			if x == num {
				count++
			}
		}

		result = append(result, Pair{
			Num:   num,
			Count: count,
		})

	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

    output := []int{}
    for i := 0; i < k; i++ {
        output = append(output,result[i].Num)
    }
	return output
}

// hashmap + sort
type Pair struct {
	Num   int
	Count int
}

func topKFrequent(nums []int, k int) []int {
	frequencies := map[int]int{}

    for i := 0; i < len(nums); i++ {
        frequencies[nums[i]]++ 
    }

    result := []Pair{}

    for num, count := range frequencies {
        result = append(result, Pair{Num:num,Count:count})
    }
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	output := []int{}
	for i := 0; i < k; i++ {
		output = append(output, result[i].Num)
	}
	return output
}


// heashmap + min heap
package main

import (
	"container/heap"
)

type Pair struct {
	Num   int
	Count int
}

// MinHeap implements heap.Interface. O(n log k)
type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

// Min Heap based on frequency
func (h MinHeap) Less(i, j int) bool {
	return h[i].Count < h[j].Count
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

func topKFrequent(nums []int, k int) []int {
	// Step 1: Count frequencies
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	// Step 2: Create Min Heap
	h := &MinHeap{}
	heap.Init(h)

	// Step 3: Keep only top k elements in the heap
	for num, count := range freq {
		heap.Push(h, Pair{
			Num:   num,
			Count: count,
		})

		if h.Len() > k {
			heap.Pop(h) // Removes the smallest frequency
		}
	}

	// Step 4: Extract answer
	result := make([]int, 0, k)

	for h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		result = append(result, p.Num)
	}

	return result
}

// hashmap + bucket sort O(n)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int,len(nums))

	for _, num := range nums {
		freq[num]++
	}

    buckets := make([][]int,len(nums))
	for num, count := range freq {
		buckets[count-1] = append(buckets[count-1],num)
	}

    final := []int{}
    for i := len(buckets)-1; i >= 0; i-- {
        if len(final) == k {
            break;
        }
        if len(buckets[i]) > 0 {
            final = append(final,buckets[i]...)
        } 
    }
    return final

}
