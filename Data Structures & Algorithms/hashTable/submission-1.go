type Node struct {
    key int
    val int
    next *Node
}

func NewNode(key, val int) *Node {
    return &Node{key: key, val : val, next: nil}
}

type HashTable struct {
    capacity int
    size int
    table []*Node
}

func NewHashTable(capacity int) *HashTable {
    return &HashTable{
        capacity : capacity,
        size : 0,
        table : make([]*Node, capacity),
    }
}

func (ht *HashTable) Insert(key, value int) {
    index := ht.hash(key)
    node := ht.table[index]

    // case : index not occupied
    if node == nil {
        ht.table[index] = NewNode(key, value)
        ht.size++
    } else {
        // case : index occupied with same key
        var prev *Node
        for node != nil {
            if node.key == key {
                node.val = value
                return
            }
            prev, node = node, node.next
        }
        // case : collision
        prev.next = NewNode(key, value)
        ht.size++
    }
    // Check if resizing is needed
	// if float64(ht.size)/float64(ht.capacity) >= 0.5 {
	// 	ht.Resize()
	// }

    if ht.size >= ht.capacity/2 {
        ht.Resize()
    }
}

func (ht *HashTable) Get(key int) int {
    index := ht.hash(key)
    node := ht.table[index]
    
    for node != nil {
        // check if a value exists for the key
        if node.key == key {
            return node.val
        }
        // keep looping
        node = node.next
    }
    return -1
}

func (ht *HashTable) Remove(key int) bool {
    index := ht.hash(key)
    node := ht.table[index]

    var prev *Node
    for node != nil {
        // find key
        if node.key == key {
            // remove node
            if prev != nil {
                prev.next = node.next // pop middle node
            } else {
                ht.table[index] = node.next // pop head
            }
            ht.size--
            return true
        }
        // keep looping
        prev, node = node, node.next
    }
    return false
}

func (ht *HashTable) GetSize() int {
    return ht.size
}

func (ht *HashTable) GetCapacity() int {
    return ht.capacity
}

func (ht *HashTable) Resize() {
    newCapacity := 2 * ht.capacity
    newTable := make([]*Node, newCapacity)

    // migrate all elements to the new map
    for _, node := range ht.table {
        // resolve each LinkedList
        for node != nil {
            index := node.key % newCapacity
            if newTable[index] == nil {
                newTable[index] = NewNode(node.key, node.val)
            } else {
                newNode := newTable[index]
                // move to the end of the LinkedList
                for newNode.next != nil {
                    newNode = newNode.next
                }
                newNode.next = NewNode(node.key, node.val)
            }
            node = node.next
        }
    }
    ht.capacity = newCapacity
    ht.table = newTable
}

func (ht *HashTable) hash(key int) int {
    // simple modulo 
    return key % ht.capacity
}
