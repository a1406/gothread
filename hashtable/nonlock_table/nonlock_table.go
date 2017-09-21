package nonlock_table

const bucket_num uint = 8
type Table_nonlock struct {
	ht_bkt [bucket_num]ht_bucket
}

type ht_bucket struct {
	data int
	next *ht_bucket
}

func (t *Table_nonlock)Init() {
}
func (t *Table_nonlock)Lookup(k int) bool{
	return false
}
func (t *Table_nonlock)Insert(k int) {
}
func (t *Table_nonlock)Delete(k int) {
}
