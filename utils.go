package utils

func CopySlice[T interface{}](v []T) []T {
	tmp := make([]T, len(v))
	copy(tmp, v)
	return tmp
}

// TODO optimization
func SliceDiff[T comparable](a []T, b []T) []T {
	set := make(map[T]struct{})
	for i, max := 0, len(b); i < max; i++ {
		set[b[i]] = struct{}{}
	}

	res := make([]T, 0)
	for i, max := 0, len(a); i < max; i++ {
		if _, ok := set[a[i]]; !ok {
			res = append(res, a[i])
		}
	}

	return res
}

// TODO replace with binary search
func SearchInSlice[T comparable](needle T, haystack []T) int {
	for i := range haystack {
		if haystack[i] == needle {
			return i
		}
	}
	return -1
}

func IsIndexExist[T interface{}](index int, slice []T) bool {
	return index >= 0 && index < len(slice)
}

// TODO generic
func unique(input []string) []string {
	set := make(map[string]struct{})
	for i := range input {
		set[input[i]] = struct{}{}
	}

	res := make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	return res
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SetOf[V any, K comparable](v []V, fn func(V) K) map[K]V {
	set := make(map[K]V, len(v))
	for _, val := range v {
		set[fn(val)] = val
	}
	return set
}

func Mapping[IN any, OUT any](input []IN, fn func(IN) OUT) []OUT {
	res := make([]OUT, 0, len(input))
	for _, val := range input {
		res = append(res, fn(val))
	}
	return res
}

// func printTree(t *TreeNode) {
// 	if t == nil {
// 		fmt.Print("nil", ",")
// 		return
// 	} else {
// 		fmt.Print(t.Val, ",")
// 	}

// 	printTree(t.Left)
// 	printTree(t.Right)
// }
