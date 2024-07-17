package arraysstrings

func mergeAlternately(word1 string, word2 string) string {
	var merged string
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		merged += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		merged += string(word1[i])
		i++
	}

	for j < len(word2) {
		merged += string(word2[j])
		j++
	}

	return merged
}

func mergeAlternately2(word1 string, word2 string) string {
	var merged string

	for len(word1) > 0 && len(word2) > 0 {
		merged += string(word1[0]) + string(word2[0])
		word1 = word1[1:]
		word2 = word2[1:]
	}

	if len(word1) > 0 {
		merged += word1
	}

	if len(word2) > 0 {
		merged += word2
	}

	return merged
}

/*
	goos: linux
	goarch: amd64
	pkg: github.com/AydenHex/leetcode_75/go/arrays_strings
	cpu: 12th Gen Intel(R) Core(TM) i9-12900KF
	BenchmarkSimpleMerge/input_size_10-24            2879959               350.1 ns/op
	BenchmarkSimpleMerge/input_size_100-24            200137              5127 ns/op
	BenchmarkSimpleMerge/input_size_1000-24             4374            255160 ns/op
	BenchmarkSimpleMerge/input_size_10000-24              68          18574179 ns/op
	BenchmarkSimpleMerge2/input_size_10-24           3093956               401.7 ns/op
	BenchmarkSimpleMerge2/input_size_100-24           233101              5794 ns/op
	BenchmarkSimpleMerge2/input_size_1000-24            4688            259578 ns/op
	BenchmarkSimpleMerge2/input_size_10000-24             80          16555701 ns/op
*/
