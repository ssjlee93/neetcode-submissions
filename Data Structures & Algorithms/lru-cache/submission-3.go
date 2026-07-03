type LRUCache struct {
    // keep hashmap of key values with capacity
	cache map[int]int
	// Go map has len tracks capacity
	// but eaiser to track as a value
	capacity int
	// LRU track with slice. Queue ideal, but Go is faster with slice.
	lru []int	// keep track of keys used. popleft when Put operation exceeds
}

func Constructor(capacity int) LRUCache {
    return LRUCache{ cache:make(map[int]int, capacity), capacity: capacity, lru: make([]int, 0)}
}

func (this *LRUCache) Get(key int) int {
	// use Go built in syntax to check existence
    val, ok := this.cache[key]
	if !ok {
		return -1
	}
	// remove existing key in lru
	this.remove(key)
	// add new key at the end
	this.lru = append(this.lru, key)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	// if value exists, update
	if _, ok := this.cache[key]; ok {
		this.cache[key] = value
		this.remove(key)
		this.lru = append(this.lru, key)
		return
	}
    // if within capacity, 
	// add and return
	if len(this.cache) < this.capacity {
		this.cache[key] = value
		this.lru = append(this.lru, key)
		return
	}
	// if outside capacity, 
	// popleft, which is least used
	popleft := this.lru[0]
	// remove this least used key from lru and cache
	this.lru = this.lru[1:]
	delete(this.cache, popleft)
	// add our value
	this.cache[key] = value
	this.lru = append(this.lru, key)
	return
}

func (this *LRUCache) remove(key int) {
	for i, v := range this.lru {
		if v == key {
			this.lru = append(this.lru[:i], this.lru[i+1:]...)
			break
		}
	}
}