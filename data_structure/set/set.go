// golang set operate

package set

import (
	"sort"
)

// Set : 集合类型
type Set []interface{}

// IntersectionTwo : 两个集合取交集
// 空间换时间，O(m*n) -> O(m + n)
func IntersectionTwo(s1, s2 Set) Set {
	if len(s1) == 0 || len(s2) == 0 {
		return Set{}
	}

	tmp := make(map[interface{}]bool, len(s1))
	for _, v := range s1 {
		if _, ok := tmp[v]; !ok {
			tmp[v] = true
		}
	}

	rs := make(Set, 0)
	for _, v := range s2 {
		if val, ok := tmp[v]; ok && val {
			rs = append(rs, v)
		}
	}
	return rs
}

// Intersection : 交集操作
// 递归实现多个集合的交集
func Intersection(ss ...Set) Set {
	if ss == nil || len(ss) == 0 {
		return Set{}
	}
	if len(ss) == 1 {
		return ss[0]
	}
	set := ss[0]
	for _, s := range ss[1:] {
		set = IntersectionTwo(set, s)
	}
	return set
}

// UnionTwo : 取两个集合的并集
func UnionTwo(s1, s2 Set) Set {
	tmp := make(map[interface{}]bool, len(s1))
	for _, v := range s1 {
		if _, ok := tmp[v]; !ok {
			tmp[v] = true
		}
	}
	for _, v := range s2 {
		if _, ok := tmp[v]; !ok {
			tmp[v] = true
		}
	}
	set := Set{}
	for k, v := range tmp {
		if v {
			set = append(set, k)
		}
	}
	return set
}

// Union : 并集
func Union(ss ...Set) Set {
	set := Set{}
	for _, s := range ss {
		set = UnionTwo(set, s)
	}
	return set
}

// DiffTwo : 计算两个集合差集
func DiffTwo(s1, s2 Set) Set {
	// 先取交集
	inter := IntersectionTwo(s1, s2)
	interMap := make(map[interface{}]bool)
	for _, v := range inter {
		interMap[v] = true
	}
	set := Set{}
	for _, v := range s1 {
		if _, ok := interMap[v]; !ok {
			set = append(set, v)
		}
	}
	for _, v := range s2 {
		if _, ok := interMap[v]; !ok {
			set = append(set, v)
		}
	}

	return set
}

// Diff : 差集
// TODO 有点复杂，没想明白，暂不做。
func Diff(ss ...Set) Set {
	// 1. 两两取交集，得到交集集合X

	// 2. 交集集合X取并集D

	// 3. 返回不在并集D中的所有元素
	return nil
}

// Sort : 排序
func Sort(s Set) Set {
	sort.Slice(s, func(i, j int) bool {
		switch v := s[i].(type) {
		case int:
			return v < s[j].(int)
		case int16:
			return v < s[j].(int16)
		case int8:
			return v < s[j].(int8)
		case int32:
			return v < s[j].(int32)
		case int64:
			return v < s[j].(int64)
		case uint:
			return v < s[j].(uint)
		case uint8:
			return v < s[j].(uint8)
		case uint16:
			return v < s[j].(uint16)
		case uint32:
			return v < s[j].(uint32)
		case uint64:
			return v < s[j].(uint64)
		case float32:
			return v < s[j].(float32)
		case float64:
			return v < s[j].(float64)
		case string:
			return v < s[j].(string)
		default:
			return true
		}
	})
	return s
}
