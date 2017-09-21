package rcu_table

import (
	"sync"
)

const bucket_num uint = 8
type Table_rcu struct {
	ht_bkt [bucket_num]ht_bucket
}

type ht_bucket struct {
	head *ht_bucket_entry
	readers uint
	locker *sync.Mutex
}

type ht_bucket_entry struct {
	data int
	next *ht_bucket_entry
}

func (t *Table_rcu)Num() uint {
	var ret uint
	for i := uint(0); i < bucket_num; i++ {
		t.ht_bkt[i].locker.Lock()		
		for cur := t.ht_bkt[i].head; cur != nil; cur = cur.next {
			ret++
		}
		t.ht_bkt[i].locker.Unlock()
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
	t.ht_bkt[i].locker.Lock()
	for cur := t.ht_bkt[i].head; cur != nil; cur = cur.next {
		if cur.data == k {
			t.ht_bkt[i].locker.Unlock()
			return true
		}
	}
	t.ht_bkt[i].locker.Unlock()	
	return false
}
func (t *Table_rcu)Insert(k int) {
	i := uint(k) % bucket_num
	var bucket ht_bucket_entry
	bucket.data = k

	t.ht_bkt[i].locker.Lock()	
	bucket.next = t.ht_bkt[i].head
	t.ht_bkt[i].head = &bucket
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
