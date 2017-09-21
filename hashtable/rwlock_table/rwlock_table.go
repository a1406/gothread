package rwlock_table

import (
	"sync"
)

//const bucket_num uint = 8
type Table_lock struct {
	size uint	
	ht_bkt []ht_bucket
}

type ht_bucket struct {
	head *ht_bucket_entry
	locker *sync.RWMutex
}

type ht_bucket_entry struct {
	data int
	next *ht_bucket_entry
}

func (t *Table_lock)Num() uint {
	var ret uint
	for i := uint(0); i < t.size; i++ {
		t.ht_bkt[i].locker.RLock()
		for cur := t.ht_bkt[i].head; cur != nil; cur = cur.next {
			ret++
		}
		t.ht_bkt[i].locker.RUnlock()
	}
	return ret
}

func (t *Table_lock)Init(size uint) {
	t.size = size
	t.ht_bkt = make([]ht_bucket, size)	
	for i := uint(0); i < t.size; i++ {
		t.ht_bkt[i].locker = new(sync.RWMutex)
	}
}
func (t *Table_lock)Lookup(k int) bool {
	i := uint(k) % t.size
	t.ht_bkt[i].locker.RLock()
	for cur := t.ht_bkt[i].head; cur != nil; cur = cur.next {
		if cur.data == k {
			t.ht_bkt[i].locker.RUnlock()
			return true
		}
	}
	t.ht_bkt[i].locker.RUnlock()	
	return false
}
func (t *Table_lock)Insert(k int) {
	i := uint(k) % t.size
	var bucket ht_bucket_entry
	bucket.data = k

	t.ht_bkt[i].locker.Lock()
	bucket.next = t.ht_bkt[i].head
	t.ht_bkt[i].head = &bucket
	t.ht_bkt[i].locker.Unlock()		
	return
}
func (t *Table_lock)Delete(k int) bool {
	i := uint(k) % t.size
	
	t.ht_bkt[i].locker.Lock()		
	pre := t.ht_bkt[i].head
	if pre == nil {
		t.ht_bkt[i].locker.Unlock()				
		return false
	}
	if pre.data == k {
		t.ht_bkt[i].head = pre.next
		t.ht_bkt[i].locker.Unlock()				
		return true
	}
	
	for cur := pre.next; cur != nil; cur = cur.next {
		if cur.data == k {
			pre.next = cur.next
			t.ht_bkt[i].locker.Unlock()					
			return true
		}
		pre = cur
	}
	t.ht_bkt[i].locker.Unlock()			
	return false
}
