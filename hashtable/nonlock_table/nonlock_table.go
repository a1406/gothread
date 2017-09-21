package nonlock_table

//const bucket_num uint = 8
type Table_nonlock struct {
	size uint
	ht_bkt []ht_bucket
}

type ht_bucket struct {
	head *ht_bucket_entry
}
type ht_bucket_entry struct {
	data int
	next *ht_bucket_entry
}

func (t *Table_nonlock)Num() uint {
	var ret uint
	for i := uint(0); i < t.size; i++ {
		for cur := t.ht_bkt[i].head; cur != nil; cur = cur.next {
			ret++
		}
	}
	return ret
}

func (t *Table_nonlock)Init(size uint) {
	t.size = size
	t.ht_bkt = make([]ht_bucket, size)
}
func (t *Table_nonlock)Lookup(k int) bool {
	i := uint(k) % t.size
	for cur := t.ht_bkt[i].head; cur != nil; cur = cur.next {
		if cur.data == k {
			return true
		}
	}
	return false
}
func (t *Table_nonlock)Insert(k int) {
	i := uint(k) % t.size
	var bucket ht_bucket_entry
	bucket.data = k
	bucket.next = t.ht_bkt[i].head
	t.ht_bkt[i].head = &bucket
	return
}
func (t *Table_nonlock)Delete(k int) bool {
	i := uint(k) % t.size
	pre := t.ht_bkt[i].head
	if pre == nil {
		return false
	}
	if pre.data == k {
		t.ht_bkt[i].head = pre.next
		return true
	}
	
	for cur := pre.next; cur != nil; cur = cur.next {
		if cur.data == k {
			pre.next = cur.next
			return true
		}
		pre = cur
	}
	return false
}
