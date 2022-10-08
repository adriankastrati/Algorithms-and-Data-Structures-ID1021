package breadth

import (
	"algo/binary-tree/binaryTree"
	"algo/queue/queue"
)

var testTreeIt treeIterator 
var testTree binaryTree.BinaryTree

func setUp(){
	treeIt.queue = &queue.QueueTwo{}
	treeIt
}
func TestIteratorEmpty(t *testing.T) {

