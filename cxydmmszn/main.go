package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

func main() {
	bitree := &common.BiTree{}
	bitree.Init()
	root := common.NewBiTreeNode(1)
	bitree.SetRoot(root)
	bitree.Ins_left(root, 2)
	bitree.Ins_right(root, 3)



	fmt.Println("preorder-----------------")
	pl := new(common.List)
	bitree.Root().Preorder(pl)

	pl.Traverse(common.Tf) //1 2 3

	fmt.Println("preorderno-----------------")

	pls := new(common.List)
	bitree.Root().PreorderNo(pls)

	pls.Traverse(common.Tf) //1 2 3

	fmt.Println("inorder-----------------")

	il := new(common.List)
	bitree.Root().Inorder(il)

	il.Traverse(common.Tf) //2 1 3

	fmt.Println("inordreno-----------------")

	ils := new(common.List)
	bitree.Root().InorderrNo(ils)

	ils.Traverse(common.Tf) //2 1 3

	fmt.Println("postordre-----------------")

	pol := new(common.List)
	bitree.Root().Postorder(pol)

	pol.Traverse(common.Tf) //2 3 1

	fmt.Println("postorderno-----------------")

	pols := new(common.List)
	bitree.Root().PostorderNo(pols)

	pols.Traverse(common.Tf) //2 3 1

	fmt.Println("leverorder-----------------")

	ll := new(common.List)
	bitree.Root().LevelorderNo(ll)

	ll.Traverse(common.Tf) //1 2 3

}
