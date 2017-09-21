package nonlock_table

const bucket_num uint = 8
type Table_nonlock struct {
	ht_bkt [bucket_num]*ht_bucket
}

type ht_bucket struct {
	data int
	next *ht_bucket
}

func (t *Table_nonlock)Init() {
}
func (t *Table_nonlock)Lookup(k int) bool {
	i := uint(k) % bucket_num
	for cur := t.ht_bkt[i]; cur != nil; cur = cur.next {
		if cur.data == k {
			return true
		}
	}
	return false
}
func (t *Table_nonlock)Insert(k int) {
	i := uint(k) % bucket_num
	var bucket ht_bucket
	bucket.data = k
	bucket.next = t.ht_bkt[i]
	t.ht_bkt[i] = &bucket
	return
}
func (t *Table_nonlock)Delete(k int) bool {
	i := uint(k) % bucket_num
	pre := t.ht_bkt[i]
	if pre == nil {
		return false
	}
	if pre.data == k {
		t.ht_bkt[i] = pre.next
		return true
	}
	
	for cur := pre.next; cur != nil; cur = cur.next {
		if cur.data == k {
			pre.next = cur.next
			return true
		}
	}
	return false
}
