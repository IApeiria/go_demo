package main

func main() {
	s := "abc"
	// 忽略 2nd value，支持 string/array/slice/map。
	for i := range s {
		println(s[i])
	}
	// 忽略 index。
	for _, c := range s {
		println(c)
	}
	// 忽略全部返回值，仅迭代。
	for range s {

	}

	m := map[string]int{"a": 1, "b": 2}
	// 返回 (key, value)。
	for k, v := range m {
		println(k, v)
	}
}
