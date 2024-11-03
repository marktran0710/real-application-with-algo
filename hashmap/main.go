package hashmap

type Node struct {
	Key   string
	Value interface{}
	Next  *Node
}

type HashMap struct {
	buckets []*Node
	size    int
}

func (hm *HashMap) Size(value int) {
	hm.buckets = make([]*Node, value)
	hm.size = value
}

func (hm *HashMap) hashFunc(key string) (index int) {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}

	return hash % hm.size
}

func (hm *HashMap) Set(key string, value interface{}) {
	var index = hm.hashFunc(key)
	var newNode = &Node{
		Key:   key,
		Value: value,
		Next:  nil,
	}

	if hm.buckets[index] == nil {
		hm.buckets[index] = newNode
	} else {
		list := hm.buckets[index]

		for list.Next != nil {
			list = list.Next
		}

		list.Next = newNode
	}
}

func (hm *HashMap) Get(key string) (value interface{}, found bool) {
	var index = hm.hashFunc(key)
	var head = hm.buckets[index]

	for head != nil {
		if head.Key == key {
			return head.Value, true
		}

		head = head.Next
	}

	var zeroValue interface{}
	return zeroValue, false
}

func remove() {}

func containsKey() {}

func size() {}

func isEmpty() {}

func keys() {}

func values() {}

func clear() {}
