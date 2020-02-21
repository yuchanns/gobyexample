package misc

import "testing"

func TestUserAges(t *testing.T) {
	ua := UserAges{ages: make(map[string]int)}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			ua.Add("yuchanns", i)
		}(i)
	}
	for i := 0; i < 10000; i++ {
		go func() {
			println(ua.Get("yuchanns"))
		}()
	}
}
