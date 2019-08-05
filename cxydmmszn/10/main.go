package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"strings"
	"time"
	"fmt"
)

type Pet interface {
	PetName() string
}

type Cat struct {
	name string
}

func (c *Cat) PetName() string {
	return "cat_" + c.name
}

type Dog struct {
	name string
}

func (c *Dog) PetName() string {
	return "dog_" + c.name
}

//构造新的结构
type PetEnqueueEle struct {
	pet   Pet
	count int64
}

func (c *PetEnqueueEle) GetPet() Pet {
	return c.pet
}
func (c *PetEnqueueEle) Count() int64 {
	return c.count
}
func (c *PetEnqueueEle) GetPetName() string {
	return c.pet.PetName()
}

type PetEnqueue struct {
	dogQ *common.Queue
	catQ *common.Queue
}

func (pq *PetEnqueue) Add(pet Pet) {
	if strings.HasPrefix(pet.PetName(), "dog") {
		pq.dogQ.EnQueue(&PetEnqueueEle{pet, time.Now().UnixNano()})
	} else if strings.HasPrefix(pet.PetName(), "cat") {
		pq.catQ.EnQueue(&PetEnqueueEle{pet, time.Now().UnixNano()})
	}
}

func (pq *PetEnqueue) Pull() Pet {
	if pq.dogQ.Size() > 0 && pq.catQ.Size() > 0 {
		if pq.dogQ.Head().(*PetEnqueueEle).Count() < pq.catQ.Head().(*PetEnqueueEle).Count() {
			return pq.dogQ.DeQueueWithValue().(*PetEnqueueEle).pet
		} else {
			return pq.catQ.DeQueueWithValue().(*PetEnqueueEle).pet
		}
	} else if pq.dogQ.Size() > 0 {
		return pq.dogQ.DeQueueWithValue().(*PetEnqueueEle).pet
	} else if pq.catQ.Size() > 0 {
		return pq.catQ.DeQueueWithValue().(*PetEnqueueEle).pet

	} else {
		return nil
	}
}

func main() {
	pq := &PetEnqueue{&common.Queue{}, &common.Queue{}}

	pq.Add(&Cat{"1"})
	pq.Add(&Dog{"1"})
	pq.Add(&Cat{"2"})
	pq.Add(&Dog{"2"})

	c := pq.Pull()
	fmt.Println(c.PetName())

	c = pq.Pull()
	fmt.Println(c.PetName())

	c = pq.Pull()
	fmt.Println(c.PetName())

	c = pq.Pull()
	fmt.Println(c.PetName())
}
