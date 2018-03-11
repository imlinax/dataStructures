/*
double linked list implement
*/
package dlist

// Element element in a double linked list
type Element struct {
	prev, next *Element
	list       *DList
	value      interface{}
}

// DList a double linked list
type DList struct {
	root Element // a dummy head
	len  int
}

func (e *Element) Next() *Element {
	if e.list != nil {
		return e.next
	}
	return nil
}
func (e *Element) Prev() *Element {
	if e.list != nil {
		return e.prev
	}
	return nil
}

// Init initialize or reset the list
func (dl *DList) Init() *DList {
	dl.root.prev = &dl.root
	dl.root.next = &dl.root
	dl.len = 0
	return dl
}

// Len list length
func (dl *DList) Len() int {
	return dl.len
}

// IsEmpty judge list is empty or not
func (dl *DList) IsEmpty() bool {
	return dl.len == 0
}

// InsertValue insert an element with value v after 'at'
func (dl *DList) InsertValue(v interface{}, at *Element) *Element {
	newElement := &Element{
		prev:  at,
		next:  at.next,
		list:  dl,
		value: v}

	at.next.prev = newElement
	at.next = newElement
	dl.len++

	return newElement
}

// PushFront insert an element with value v at head
func (dl *DList) PushFront(v interface{}) *Element {
	return dl.InsertValue(v, &dl.root)
}

// PushBack insert an element with value v at end
func (dl *DList) PushBack(v interface{}) *Element {
	return dl.InsertValue(v, dl.root.prev)
}

// Remove remove element from list
func (dl *DList) Remove(e *Element) interface{} {
	if e.list == dl {
		e.prev.next = e.next
		e.next.prev = e.prev

		// avoid memory leak
		e.next = nil
		e.prev = nil
		e.list = nil

		dl.len--
	}
	return e.value
}

func (dl *DList) Front() *Element {
	if dl.len == 0 {
		return nil
	}
	return dl.root.next
}
func (dl *DList) Back() *Element {
	if dl.len == 0 {
		return nil
	}
	return dl.root.prev
}
