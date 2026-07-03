type Entry struct {
    key      int
    val      int
    occupied bool
    deleted  bool // tombstone
}

type HashTable struct {
    capacity int
    size     int
    table    []Entry
}

func NewHashTable(capacity int) *HashTable {
    return &HashTable{
        capacity: capacity,
        size:     0,
        table:    make([]Entry, capacity),
    }
}

func (ht *HashTable) Insert(key, value int) {
    index := ht.hash(key)
    firstDeleted := -1

    for i := 0; i < ht.capacity; i++ {
        probe := (index + i) % ht.capacity
        e := &ht.table[probe]

        if e.occupied {
            if e.key == key {
                e.val = value // update existing
                return
            }
        } else if e.deleted {
            if firstDeleted == -1 {
                firstDeleted = probe // save earliest tombstone
            }
        } else {
            // truly empty — key is definitely not in table
            insertAt := probe
            if firstDeleted != -1 {
                insertAt = firstDeleted // reuse tombstone slot
            }
            ht.table[insertAt] = Entry{key: key, val: value, occupied: true}
            ht.size++
            if ht.size >= ht.capacity/2 {
                ht.Resize()
            }
            return
        }
    }

    // no empty slot found, but we have a tombstone
    if firstDeleted != -1 {
        ht.table[firstDeleted] = Entry{key: key, val: value, occupied: true}
        ht.size++
        if ht.size >= ht.capacity/2 {
            ht.Resize()
        }
    }
}

func (ht *HashTable) Get(key int) int {
    index := ht.hash(key)

    for i := 0; i < ht.capacity; i++ {
        probe := (index + i) % ht.capacity
        e := &ht.table[probe]

        if e.occupied {
            if e.key == key {
                return e.val
            }
        } else if !e.deleted {
            return -1 // truly empty slot — key can't be past here
        }
        // tombstone: keep probing
    }
    return -1
}

func (ht *HashTable) Remove(key int) bool {
    index := ht.hash(key)

    for i := 0; i < ht.capacity; i++ {
        probe := (index + i) % ht.capacity
        e := &ht.table[probe]

        if e.occupied {
            if e.key == key {
                e.occupied = false
                e.deleted = true // leave tombstone, don't break probe chain
                ht.size--
                return true
            }
        } else if !e.deleted {
            return false
        }
    }
    return false
}

func (ht *HashTable) Resize() {
    newCapacity := 2 * ht.capacity
    newTable := make([]Entry, newCapacity)

    for _, e := range ht.table {
        if !e.occupied {
            continue // skip empty and tombstone slots — natural cleanup
        }
        index := e.key % newCapacity
        for i := 0; i < newCapacity; i++ {
            probe := (index + i) % newCapacity
            if !newTable[probe].occupied {
                newTable[probe] = Entry{key: e.key, val: e.val, occupied: true}
                break
            }
        }
    }

    ht.capacity = newCapacity
    ht.table = newTable
}

func (ht *HashTable) GetSize() int     { return ht.size }
func (ht *HashTable) GetCapacity() int { return ht.capacity }
func (ht *HashTable) hash(key int) int { return key % ht.capacity }