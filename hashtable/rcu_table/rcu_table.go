package rcu_table

import (
	"sync"
	"sync/atomic"	
)

const bucket_num uint = 8
type Table_rcu struct {
	ht_bkt [bucket_num]ht_bucket
}

type ht_bucket struct {
	head *ht_bucket_entry
	readers int32
	locker *sync.Mutex
}

type ht_bucket_entry struct {
	data int
	next *ht_bucket_entry
}

func (t *Table_rcu)Num() uint {
	var ret uint
	for i := uint(0); i < bucket_num; i++ {
		atomic.AddInt32(&t.ht_bkt[i].readers, 1)
		for cur := t.ht_bkt[i].head; cur != nil; cur = cur.next {
			ret++
		}
		atomic.AddInt32(&t.ht_bkt[i].readers, -1)		
	}
	return ret
}

func (t *Table_rcu)Init() {
	for i := uint(0); i < bucket_num; i++ {
		t.ht_bkt[i].locker = new(sync.Mutex)
	}
}
func (t *Table_rcu)Lookup(k int) bool {
	i := uint(k) % bucket_num
	atomic.AddInt32(&t.ht_bkt[i].readers, 1)
	for cur := t.ht_bkt[i].head; cur != nil; cur = cur.next {
		if cur.data == k {
			t.ht_bkt[i].locker.Unlock()
			return true
		}
	}
	atomic.AddInt32(&t.ht_bkt[i].readers, -1)		
	return false
}
func (t *Table_rcu)Insert(k int) {
	i := uint(k) % bucket_num
	var bucket ht_bucket_entry
	bucket.data = k

	t.ht_bkt[i].locker.Lock()
	var new_head *ht_bucket_entry
	new_head = t.ht_bkt[i].head
	
	bucket.next = new_head
	new_head = &bucket
	
	t.ht_bkt[i].locker.Unlock()		
	return
}
func (t *Table_rcu)Delete(k int) bool {
	i := uint(k) % bucket_num
	
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
