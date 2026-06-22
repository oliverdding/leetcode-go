package id1146

/*
 * @lc app=leetcode.cn id=1146 lang=golang
 *
 * [1146] Snapshot Array
 */

// @lc code=start
type SnapshotArray struct {
	data   [][]SnapVal
	snapID int
}

type SnapVal struct {
	snapID int
	val    int
}

func Constructor(length int) SnapshotArray {
	data := make([][]SnapVal, length)
	return SnapshotArray{
		data,
		0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.data[index] = append(this.data[index], SnapVal{
		snapID: this.snapID,
		val:    val,
	})
}

func (this *SnapshotArray) Snap() int {
	this.snapID++
	return this.snapID - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	arr := this.data[index]
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if arr[mid].snapID > snap_id {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if lo > 0 {
		return arr[lo-1].val
	}
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
// @lc code=end
