package queue

import (
	"math/rand"
	"sync"
	"unsafe"
)

const minQueueLen = 32

type Queue struct {
	items             map[int64]unsafe.Pointer
	ids               map[unsafe.Pointer]int64
	buf               []int64
	head, tail, count int
	mutex             *sync.Mutex
	notEmpty          *sync.Cond
	//订阅该通道判断队列是否是空的还是非空的
	NotEmpty chan struct{}
}

func New() *Queue {
	q := &Queue{
		items:    make(map[int64]unsafe.Pointer),
		ids:      make(map[unsafe.Pointer]int64),
		buf:      make([]int64, minQueueLen),
		mutex:    &sync.Mutex{},
		NotEmpty: make(chan struct{}, 1),
	}

	q.notEmpty = sync.NewCond(q.mutex)

	return q
}

//清除队列中所有的数据
func (q *Queue) Clean() {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = make(map[int64]unsafe.Pointer)
	q.ids = make(map[unsafe.Pointer]int64)
	q.buf = make([]int64, minQueueLen)
	q.tail = 0
	q.head = 0
	q.count = 0
}

//返回队列中数据长度
func (q *Queue) Length() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return len(q.items)
}

//扩充队列容量
func (q *Queue) resize() {
	newCount := q.count << 1

	if q.count < 2<<18 {
		newCount = newCount << 2
	}

	newBuf := make([]int64, newCount)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

//写入订阅队列
func (q *Queue) notify() {
	if len(q.items) > 0 {
		select {
		case q.NotEmpty <- struct{}{}:
		default:
		}
	}
}

//在队列尾部加入一个数据
func (q *Queue) Append(elem unsafe.Pointer) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.count == len(q.buf) {
		q.resize()
	}

	id := q.newId()
	q.items[id] = elem
	q.ids[elem] = id
	q.buf[q.tail] = id
	// bitwise modulus
	//取余数
	q.tail = (q.tail + 1) & (len(q.buf) - 1)
	q.count++

	q.notify()

	if q.count == 1 {
		q.notEmpty.Broadcast()
	}
}

func (q *Queue) newId() int64 {
	for {
		id := rand.Int63()
		_, ok := q.items[id]
		if id != 0 && !ok {
			return id
		}
	}
}

//在队列头部加入数据
func (q *Queue) Prepend(elem unsafe.Pointer) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.count == len(q.buf) {
		q.resize()
	}

	q.head = (q.head - 1) & (len(q.buf) - 1)
	id := q.newId()
	q.items[id] = elem
	q.ids[elem] = id
	q.buf[q.head] = id
	// bitwise modulus
	q.count++

	q.notify()

	if q.count == 1 {
		q.notEmpty.Broadcast()
	}
}

//获取头部数据但是不出队列
func (q *Queue) Front() unsafe.Pointer {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	id := q.buf[q.head]
	if id != 0 {
		return q.items[id]
	}
	return nil
}

//查看队尾数据
func (q *Queue) Back() unsafe.Pointer {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	id := q.buf[(q.tail-1)&(len(q.buf)-1)]
	if id != 0 {
		return q.items[id]
	}
	return nil
}

func (q *Queue) pop() int64 {
	for {
		if q.count <= 0 {
			q.notEmpty.Wait()
		}

		// I have no idea why, but sometimes it's less than 0
		if q.count > 0 {
			break
		}
	}

	id := q.buf[q.head]
	q.buf[q.head] = 0

	// bitwise modulus
	q.head = (q.head + 1) & (len(q.buf) - 1)
	q.count--
	if len(q.buf) > minQueueLen && (q.count<<1) == len(q.buf) {
		q.resize()
	}

	return id
}

//出队，如果队列为空那么会被阻塞掉
func (q *Queue) Pop() unsafe.Pointer {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for {
		id := q.pop()

		item, ok := q.items[id]

		if ok {
			delete(q.ids, item)
			delete(q.items, id)
			q.notify()
			return item
		}
	}
}

//删除一个队列中的元素
func (q *Queue) Remove(elem unsafe.Pointer) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	id, ok := q.ids[elem]
	if !ok {
		return false
	}
	delete(q.ids, elem)
	delete(q.items, id)
	return true
}
