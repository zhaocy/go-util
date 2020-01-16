package go_util

import "sync"

func WithLock(locker sync.Locker, fn func()) {
	locker.Lock()
	defer locker.Unlock()
	fn()
}
