package tree

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

/*
       3
      / \
     /   \
    9    20
   /	 /  \
  8    15    7
*/
var (
	root              *TreeNode
	ExceptedResOfPre  = []int{3, 9, 8, 20, 15, 7}
	ExceptedResOfIn   = []int{8, 9, 3, 15, 20, 7}
	ExceptedResOfPost = []int{8, 9, 15, 7, 20, 3}
)

func init() {
	root = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9, Left: &TreeNode{Val: 8}},
		Right: &TreeNode{
			Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
}

func ReOutput(f func(root *TreeNode), root *TreeNode) string {
	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(&buf)
	f(root)
	log.SetOutput(os.Stderr)
	return strings.TrimSpace(strings.Join(strings.Split(buf.String(), "\n"), ""))
}

func TestString(t *testing.T) {
	if root.String() != "<<<<nil> 8 <nil>> 9 <nil>> 3 <<<nil> 15 <nil>> 20 <<nil> 7 <nil>>>>" {
		t.Fatal(root.String())
	}
}

func TestBFS(t *testing.T) {
	if r := ReOutput(BFS, root); r != "3 9 20 8 15 7" {
		t.Fatal(r)
	}
}
