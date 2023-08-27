package main

func findSubstring(s string, words []string) (ans []int) {
	l, m, n := len(s), len(words), len(words[0])
	for i := 0; i < n && i+m*n <= l; i++ {
		differ := make(map[string]int)
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}

		for start := i; start < l-m*n+1; start += n {
			if start != i {
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}

func main(){
	findSubstring("barfoothefoobarman", []string{"foo", "bar"})
}
