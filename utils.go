package main

func copySlice[T interface{}](v []T) []T {
	tmp := make([]T, len(v))
	copy(tmp, v)
	return tmp
}

// TODO optimization
func sliceDiff[T comparable](a []T, b []T) []T {
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

//TODO replace with binary search
func searchInSlice[T comparable](needle T, haystack []T) int {
	for i := range haystack {
		if haystack[i] == needle {
			return i
		}
	}
	return -1
}

func isIndexExist[T interface{}](index int, slice []T) bool {
	return index >= 0 && index < len(slice)
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
