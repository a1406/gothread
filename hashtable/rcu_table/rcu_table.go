package rcu_table

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

//const bucket_num uint = 8
type Table_rcu struct {
	size uint
	ht_bkt []ht_bucket
}

type ht_bucket struct {
	head *ht_bucket_entry
//	readers int32
	locker *sync.Mutex
}

type ht_bucket_entry struct {
	data int
	next *ht_bucket_entry
}

func (t *Table_rcu)Num() uint {
	var ret uint
	for i := uint(0); i < t.size; i++ {
		head := t.ht_bkt[i].load_head()
		ret += head.len()
	}
	return ret
}

func (t *ht_bucket) load_head() *ht_bucket_entry {
	t1 := unsafe.Pointer(t.head)
	head := (*ht_bucket_entry)(atomic.LoadPointer(&t1))
	return head
}
func (t *ht_bucket) save_head(head *ht_bucket_entry) {
	t1 := unsafe.Pointer(t.head)	
	atomic.StorePointer(&t1, unsafe.Pointer(head))
}


func (t *Table_rcu)Init(size uint) {
	t.size = size
	t.ht_bkt = make([]ht_bucket, size)	
	for i := uint(0); i < t.size; i++ {
		t.ht_bkt[i].locker = new(sync.Mutex)
	}
}
func (t *Table_rcu)Lookup(k int) bool {
	i := uint(k) % t.size
	//	atomic.AddInt32(&t.ht_bkt[i].readers, 1)
	for cur := t.ht_bkt[i].load_head(); cur != nil; cur = cur.next {
		if cur.data == k {
			return true
		}
	}
//	atomic.AddInt32(&t.ht_bkt[i].readers, -1)		
	return false
}

func (t *ht_bucket_entry)len() uint {
	var ret uint
	for cur := t; cur != nil; cur = cur.next {
		ret++
	}
	return ret
}

func (t *ht_bucket_entry)copy_bucket() *ht_bucket_entry {
	if t == nil {
		return nil
	}
	
	var ret ht_bucket_entry
	pre := &ret
	ret.data = t.data
	
	for cur := t.next; cur != nil; cur = cur.next {
		var entry ht_bucket_entry
		entry.data = cur.data
		pre.next = &entry
		pre = &entry
	}
	return &ret
}

func (t *Table_rcu)Insert(k int) {
	i := uint(k) % t.size
	var bucket ht_bucket_entry
	bucket.data = k

	t.ht_bkt[i].locker.Lock()
	new_head := t.ht_bkt[i].load_head().copy_bucket()
	bucket.next = new_head
	t.ht_bkt[i].save_head(&bucket)
	t.ht_bkt[i].locker.Unlock()		
	return
}
func (t *Table_rcu)Delete(k int) bool {
	i := uint(k) % t.size
	
	t.ht_bkt[i].locker.Lock()
	new_head := t.ht_bkt[i].load_head().copy_bucket()
	
	pre := new_head
	if pre == nil {
		t.ht_bkt[i].locker.Unlock()				
		return false
	}
	if pre.data == k {
//		t.ht_bkt[i].head = pre.next
		t.ht_bkt[i].save_head(pre.next)		
		t.ht_bkt[i].locker.Unlock()				
		return true
	}
	
	for cur := pre.next; cur != nil; cur = cur.next {
		if cur.data == k {
			pre.next = cur.next
			//			t.ht_bkt[i].head = new_head
			t.ht_bkt[i].save_head(new_head)					
			t.ht_bkt[i].locker.Unlock()					
			return true
		}
		pre = cur
	}
	t.ht_bkt[i].locker.Unlock()			
	return false
}
