package id981

/*
 * @lc app=leetcode.cn id=981 lang=golang
 *
 * [981] Time Based Key-Value Store
 */

// @lc code=start
type TimeMap struct {
	data map[string][]timeValue
}

type timeValue struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]timeValue),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], timeValue{
		timestamp,
		value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.data[key]
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if arr[mid].timestamp > timestamp {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if lo > 0 {
		return arr[lo-1].value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end
