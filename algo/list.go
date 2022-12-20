/*
   Copyright 2022 Josh Deprez

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package algo

// Sometimes all you need is a linked list...

// ListNode implements a node in a doubly-linked list.
type ListNode[E any] struct {
	Prev, Next *ListNode[E]
	Value      E
}

// ListFromSlice creates a circular list from a slice, and returns a slice
// containing all the nodes in order. (Why? Often the original order is
// important for some reason.)
func ListFromSlice[E any, S ~[]E](s S) []*ListNode[E] {
	if len(s) == 0 {
		return nil
	}
	N := len(s)
	out := make([]*ListNode[E], N)
	for i, e := range s {
		out[i] = &ListNode[E]{
			Value: e,
		}
	}
	for i, n := range out[1:] {
		out[i].Next = n
		n.Prev = out[i]
	}
	out[0].Prev = out[N-1]
	out[N-1].Next = out[0]
	return out
}

// ToSlice returns a new slice containing the items in list order.
func (n *ListNode[E]) ToSlice() []E {
	if n == nil {
		return nil
	}
	out := []E{n.Value}
	for p := n.Next; p != nil && p != n; p = p.Next {
		out = append(out, p.Value)
	}
	return out
}

// Succ returns the mth item from n. For example, n.Succ(1) == n.Next.
// m can be negative (n.Succ(-1) == n.Prev). Runs in time O(|m|).
func (n *ListNode[E]) Succ(m int) *ListNode[E] {
	if n == nil {
		return nil
	}
	for j := m; j < 0 && n != nil; j++ {
		n = n.Prev
	}
	for j := m; j > 0 && n != nil; j-- {
		n = n.Next
	}
	return n
}

// Remove removes the node from the list (its neighbors point to one another).
// The node's Prev and Next remain unaltered.
func (n *ListNode[E]) Remove() {
	// p -> n -> q becomes p -> q
	p, q := n.Prev, n.Next
	p.Next = q
	q.Prev = p
}

// InsertAfter inserts n after p.
func (n *ListNode[E]) InsertAfter(p *ListNode[E]) {
	if p == n {
		return
	}
	q := p.Next
	// p -> q becomes p -> n -> q
	p.Next = n
	n.Prev = p
	n.Next = q
	q.Prev = n
}

// InsertBefore inserts n before p.
func (n *ListNode[E]) InsertBefore(p *ListNode[E]) {
	if p == n {
		return
	}
	q := p.Prev
	// q -> p becomes q -> n -> p
	q.Next = n
	n.Prev = q
	n.Next = p
	p.Prev = n
}
