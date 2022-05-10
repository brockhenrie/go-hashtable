package main

import (
	"fmt"
)

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*Bucket
}


//bucket is a linked list
type Bucket struct {
	head * BucketNode
}

//BucketNode structure
type BucketNode struct {
	key string 
	next *BucketNode
}

//Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string){
	index := hash(key)
	h.array[index].insert(key)
}

// Search will return true if a key is in the hash table array
func (h *HashTable) Search(key string) bool {
	index:= hash(key)
	return h.array[index].search(key)
}

//Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string){
	index := hash(key)
	h.array[index].delete(key)
}

//insert will rtake in a key, create a node with the key and insert the node in the bucket
func ( b *Bucket) insert(k string){
	//if key diesnt exist in bucket
	if !b.search(k){
	newNode := &BucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
	}else {
		fmt.Println(k, " already exists")
	}
}

//search will take in a key and return true if the bucket has 
func (b *Bucket) search(k string) bool{
	currentNode := b.head 
	for currentNode != nil {
		if currentNode.key == k{
			return true
		}
		currentNode = currentNode.next
	}
	return false
}


//delete
func (b *Bucket) delete(k string){

	if b.head.key == k{
		b.head = b.head.next
		return
	}


	previousNode := b.head	
	for previousNode.next != nil {
		if previousNode.next.key == k{
			//delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}


//Hash
func hash(key string) int {
	//get ascii code of each character and sum them to get the code
	sum := 0
	for _, v := range key{
		sum += int(v)
	}
	return sum & ArraySize
}

//Init() will create a bucket in each slot of the hash table
func Init() *HashTable{
//returns address of hash table
	result := &HashTable{}
	for i := range result.array {
	//creates a bucket in each spot of the array
		result.array[i] = &Bucket{}
	}
	return result
}

func main(){
	testHashTable := Init()

	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list{
		testHashTable.Insert(v)
	}

	fmt.Println(testHashTable.Search("KENNY"))
	testHashTable.Delete("KYLE")
	fmt.Println(testHashTable.Search("KYLE"))

	

 

}
