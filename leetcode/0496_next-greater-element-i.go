package leetcode

// Map + Stack solution
// e.g. nums1 = [1,2,3,4] nums2 = [4,2,1,3,5]
// iteration      curr ele    stack        actions            map
//     0              4       [4]          push(4)            {1:-1,2:-1,3:-1,4:-1}
//     1              2       [4,2]        push(2)            {1:-1,2:-1,3:-1,4:-1}
//     2              1       [4,2,1]      push(1)            {1:-1,2:-1,3:-1,4:-1}
//     3              3       [4,3]        pop(1,2),push(3)   {1:3,2:3,3:-1,4:-1}
//     4              5       []           pop(3,4)           {1:3,2:3,3:5,4:5}

type Stack []int

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

func (s *Stack) Pop() int {
	m := len(*s)
	r := (*s)[m-1]
	*s = (*s)[:m-1]
	return r
}

func (s *Stack) Top() (int, bool) {
	m := len(*s)
	if m == 0 {
		return 0, false
	}
	return (*s)[m-1], true
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// Prepare map: value -> nextBiggest
	m := make(map[int]int)
	for i := range nums1 {
		m[nums1[i]] = -1
	}

	// Iterate using stack as auxiliary data structure
	s := make(Stack, 0)
	for _, v := range nums2 { // O(len(nums2))
		// Iterate each element in nums1 push and pop in stack within the outer for loop
		for t, ok := s.Top(); ok && t < v; t, ok = s.Top() {
			p := s.Pop()
			m[p] = v
		}

		if _, ok := m[v]; ok {
			s.Push(v)
		}
	}

	// Prepare result
	result := make([]int, len(nums1))
	for i := range nums1 {
		result[i] = m[nums1[i]]
	}

	return result
}

// Time complexity:
//    for pop and push = O(len(nums1))
//    for top = O(len(nums1) + len(nums2))

// Public func to debug from main.go
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	return nextGreaterElement(nums1, nums2)
}
