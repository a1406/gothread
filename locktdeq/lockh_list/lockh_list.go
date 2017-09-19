package lockh_list

import (
       "../node"
       "sync"
)

const MAX_HASH_SIZE int = 8

type list struct {
	locker *sync.Mutex
	head *node.Node
	tail *node.Node
}

