package main

import (
	"fmt"
)

type nodeBt struct {
	value uint
	left  *nodeBt
	right *nodeBt
}

func NewNodeBt(value uint) *nodeBt {
	return &nodeBt{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (b *nodeBt) Insert(value uint) {
	if value < b.value {
		if b.left == nil {
			b.left = NewNodeBt(value)
		} else {
			b.left.Insert(value)
		}
	} else {
		if b.right == nil {
			b.right = NewNodeBt(value)
		} else {
			b.right.Insert(value)
		}
	}
}

func (b *nodeBt) InOrderTvl() {
	if b == nil {
		return
	}

	b.left.InOrderTvl()
	fmt.Println("Value :", b.value)
	b.right.InOrderTvl()
}

func (b *nodeBt) PreOrderTvs() {
	if b == nil {
		return
	}

	fmt.Println("Value :", b.value)
	b.left.PreOrderTvs()
	b.right.PreOrderTvs()
}

func (b *nodeBt) PostOrderTvs() {
	if b == nil {
		return
	}

	b.left.PostOrderTvs()
	b.right.PostOrderTvs()
	fmt.Println("Value :", b.value)
}

func (b *nodeBt) findMin() *nodeBt {
	curent := b

	if curent.left != nil {
		return curent.left
	}

	return curent
}

func (b *nodeBt) Remove(value uint) *nodeBt {
	if b == nil {
		return nil
	}

	if value < b.value {
		b.left = b.left.Remove(value)
	} else if value > b.value {
		b.right = b.right.Remove(value)
	} else {

		// kasus 1 : node adalah leaf
		if b.left == nil && b.right == nil {
			return nil
		}
		// kasus 2 : node memiliki 1 anak
		if b.left == nil {
			return b.right
		}

		if b.right == nil {
			return b.left
		}

		// kasus 3 : node memiliki 2 anak
		minRight := b.right.findMin()
		b.value = minRight.value
		b.right = b.right.Remove(minRight.value)
	}

	return b
}

func (b *nodeBt) Search(value uint) bool {
	if b == nil {
		return false
	}

	if b.value == value {
		return true
	} else if value < b.value {
		return b.left.Search(value)
	} else {
		return b.right.Search(value)
	}

}

func main() {
	fmt.Println("----- Simple Binary Tree -----")

	myBt := NewNodeBt(10)

	// Menambahkan node
	myBt.Insert(9)
	myBt.Insert(11)
	myBt.Insert(8)
	myBt.Insert(12)
	myBt.Insert(1)

	fmt.Println("In-Order Traversal Sebelum Delete:")
	myBt.InOrderTvl()

	// Menghapus node
	fmt.Println("\nMenghapus node 11...")
	myBt.Remove(11)

	fmt.Println("In-Order Traversal Setelah Delete:")
	myBt.InOrderTvl()

}
