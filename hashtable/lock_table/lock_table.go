package lock_table

import (
	"sync"
)

const bucket_num uint = 8
type Table_lock struct {
	ht_bkt [bucket_num]*ht_bucket
}

type ht_bucket struct {
	data int
	next *ht_bucket
	locker *sync.Mutex
}

func (t *Table_lock)Init() {
	for i := uint(0); i < bucket_num; i++ {
		t.ht_bkt[i].locker = new(sync.Mutex)
	}
}
func (t *Table_lock)Lookup(k int) bool {
	i := uint(k) % bucket_num
	t.ht_bkt[i].locker.Lock()
	for cur := t.ht_bkt[i]; cur != nil; cur = cur.next {
		if cur.data == k {
			t.ht_bkt[i].locker.Unlock()
			return true
		}
	}
	t.ht_bkt[i].locker.Unlock()	
	return false
}
func (t *Table_lock)Insert(k int) {
	i := uint(k) % bucket_num
	var bucket ht_bucket
	bucket.data = k

	t.ht_bkt[i].locker.Lock()	
	bucket.next = t.ht_bkt[i]
	t.ht_bkt[i] = &bucket
	t.ht_bkt[i].locker.Unlock()		
	return
}
func (t *Table_lock)Delete(k int) bool {
	i := uint(k) % bucket_num
	
	t.ht_bkt[i].locker.Lock()		
	pre := t.ht_bkt[i]
	if pre == nil {
		t.ht_bkt[i].locker.Unlock()				
		return false
	}
	if pre.data == k {
		t.ht_bkt[i] = pre.next
		t.ht_bkt[i].locker.Unlock()				
		return true
	}
	
	for cur := pre.next; cur != nil; cur = cur.next {
		if cur.data == k {
			pre.next = cur.next
			t.ht_bkt[i].locker.Unlock()					
			return true
		}
	}
	t.ht_bkt[i].locker.Unlock()			
	return false
}
