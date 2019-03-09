package solutions

import (
        "encoding/json"
	"fmt"
)

type Tree struct {
   Value interface{}
   Left *Tree
   Right *Tree
 }

func StrToMap(str  []byte)map[string]interface{}{
        f := map[string]interface{}{}
        json.Unmarshal(str, &f)

        return f
}

func CreateTree(f map[string]interface{}) *Tree{
        if f == nil{
                return nil
        }
        t := Tree{}
        t.Value = f["value"]
        if f["left"] == nil{
                t.Left = nil
        } else {
                t.Left = CreateTree(f["left"].(map[string]interface{}))
        }
        if f["right"] == nil{
                t.Right = nil
        } else {
                t.Right = CreateTree(f["right"].(map[string]interface{}))
        }
        //fmt.Printf("%#v\n", t)
        return &t

}
func isHasPathWithGivenSum(t *Tree, v int, s int) bool{

        if t == nil{
                return v == s
        }
        newVal := t.Value.(int) + v

        if t.Left == nil && t.Right == nil{
                return newVal == s
        }
        if (t.Left != nil && isHasPathWithGivenSum(t.Left, newVal, s)){
                return true
        }
        if (t.Right != nil && isHasPathWithGivenSum(t.Right, newVal, s)){
                return true
        }

        return false
}


func HasPathWithGivenSum(t *Tree, s int) bool {
        if t == nil{
                return s == 0
        }
        s -= int(t.Value.(float64))

        if t.Left == nil && t.Right == nil{
                return s == 0
        }

        return (t.Right != nil && HasPathWithGivenSum(t.Right,  s)) || (t.Left != nil && HasPathWithGivenSum(t.Left,  s))
}

func isSymmetric(t1 *Tree, t2 *Tree) bool {
	if t1 == nil && t2 == nil{
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Value == t2.Value{
		return isSymmetric(t1.Left, t2.Right) && isSymmetric(t1.Right, t2.Left)
	}
	return false
}
func isTreeSymmetric(t *Tree) bool {
	if t == nil{
		return true
	}
	return isSymmetric(t.Left, t.Right)

}


func getChild(isLeft bool, parent string) string{
	e, d := "Engineer", "Doctor"
	if parent == e{
		if isLeft {
			return e
		}
		return d
	}
	if isLeft {
		return d
	}
	return e
}

func pow2(n int) int{
	res := 1
	for i := 1; i <= n; i++{
		res *= 2
	}
	return res
}

func FindProfession(level int, pos int) string {
	parent := "Engineer"
	for i := level -1 ; i > 0; i--{
		mid := pow2(i) / 2
		if mid < pos{
			parent = getChild(false, parent)
			pos -= mid
		}else{
			parent = getChild(true, parent)
		}
		fmt.Println(parent)

	}

	return parent
}

func KthSmallestInBST(t *Tree, k int) int {
	n := 0
	for (t != nil){
		if t.Left == nil{
			n += 1
			t = t.Right
		}else{
			child := t.Left
			for(child.Right != nil && child.Right != t){

				child = child.Right
			}

			if (child.Right == nil){
				child.Right = t;
				t = t.Left;

			}else {
				n += 1
				child.Right = nil;
				if n == k{
					return int(t.Value.(float64))
				}
				t = t.Right;

			}

		}
		if n == k{
			return int(t.Value.(float64))
		}
	}
	return 0
}

func IsEquals(t1 *Tree, t2 *Tree) bool {
	if t1 == nil && t2 == nil{
		return true
	}
	if t1 != nil && t2 != nil{
		return t1.Value == t2.Value && IsEquals(t1.Left, t2.Left) && IsEquals(t1.Right, t2.Right)
	}
	return false
}
func IsSubtree(t1 *Tree, t2 *Tree) bool {
	if (t1 == nil && t2 != nil){
		return false
	}
	if IsEquals(t1, t2) {
		return true
	}
	return IsSubtree(t1.Right, t2) || IsSubtree(t1.Left, t2)

}

func getRootIndex(rootEl int, inorder []int) int {
	for i, val := range inorder {
		if val == rootEl {
			return i
		}
	}
	return -1
}

func RestoreBinaryTree(inorder []int, preorder []int) *Tree {
	if len(inorder) == 0 {
		return nil
	}
	t := &Tree{preorder[0], nil, nil}
	if len(inorder) > 1 {
		rootIn := getRootIndex(preorder[0], inorder)
		fmt.Printf("L root in : %d\t in:%v \t pre: %v\n\n", rootIn, inorder[ : rootIn], preorder[1 : rootIn + 1])
		t.Left = RestoreBinaryTree(inorder[ : rootIn], preorder[1 : rootIn + 1])
		fmt.Printf("R root in : %d\t in:%v \t pre: %v\n\n", rootIn, inorder[rootIn + 1 : ], preorder[rootIn + 1 : ])
		t.Right = RestoreBinaryTree(inorder[rootIn + 1 : ], preorder[rootIn + 1 : ])
	}

	return t
}

type TreeL struct {
	Children map[rune]*TreeL
	Last     bool
}

func addToTree(t *TreeL, s string) {
	current := t
	for _, l := range s {
		if _, ok := current.Children[l]; !ok {
			current.Children[l] = &TreeL{map[rune]*TreeL{}, false}
		}

		current = current.Children[l]

	}
	current.Last = true

}

func printTree(t *TreeL) {
	if len(t.Children) == 0 {
		fmt.Println(t.Last)
		return
	}
	for k, val := range t.Children {
		fmt.Println(string(k))
		printTree(val)
	}
}

func findSubInWord(t *TreeL, s string) string {
	begin, end := 0, 0

	for i, _ := range s {
		current := t
		j := i
		subLen := 0
		for ; j < len(s); j++ {

			if _, ok := current.Children[rune(s[j])]; ok {
				current = current.Children[rune(s[j])]
				if current.Last {
					subLen = j - i + 1

				}

			} else {
				break
			}

		}
		if subLen > 0 && end - begin < subLen {
			begin = i
			end = i + subLen
		}

	}

	if begin == end {
		return s
	}
	return s[:begin] + "[" + s[begin:end] + "]" + s[end:]
}

func FindSubstrings(words []string, parts []string) []string {
	root := TreeL{map[rune]*TreeL{}, false}

	for _, s := range parts {
		addToTree(&root, s)
	}
	fmt.Println(root.Children[rune('b')].Last)
	res := []string{}
	for _, w := range words {
		res = append(res, findSubInWord(&root, w))
	}
	return res
}

//******


func DeleteFromBST(t *Tree, queries []int) *Tree {
	for _, el := range queries {
		t = deleteEl(el, t)
		if t == nil{
			fmt.Println(el)
			return nil
		}
	}
	return t
}

func PrintTree(t *Tree, s string) {
	if t != nil {
		fmt.Println(s, int(t.Value.(float64)))
		PrintTree(t.Left, "Left")
		PrintTree(t.Right, "Right")
	}
}


//
// Definition for binary tree:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }
//
func getLastRight(root *Tree) interface{} {
	if root.Right != nil {
		return getLastRight(root.Right)
	}

	return root.Value
}

func deleteRightEl(el int, t *Tree) *Tree{
	// val := int(t.Value.(float64))
	val := t.Value.(int)
	if val == el{
		if t.Left != nil{
			return t.Left
		}
		return nil
	}
	t.Right = deleteRightEl(el, t.Right)
	return t
}


func deleteEl(el int, t *Tree) *Tree{
	if t == nil {
		return nil
	}
	val := t.Value.(int)
	if val < el {
		t.Right = deleteEl(el, t.Right)
	} else if val > el {
		t.Left = deleteEl(el, t.Left)
	} else {
		if t.Left == nil {
			return t.Right
		} else {
			t.Value = getLastRight(t.Left)
			t.Left = deleteRightEl(t.Value.(int), t.Left)
		}
	}
	return t
}


func TraverseTree(t *Tree) []int {
	res := []int{}
	if t != nil{
		nodes := []*Tree{t}
		i := 0
		for len(nodes) != 0 {
			tmp := []*Tree{}
			for _, n := range nodes {
				//res = append(res, n.Value.(int))
				res = append(res, int(n.Value.(float64)))
				if n.Left != nil {
					tmp = append(tmp, n.Left)
				}
				if n.Right != nil {
					tmp = append(tmp, n.Right)
				}
			}
			fmt.Println(len(tmp),len(nodes), res)
			nodes = tmp
			i+= 1
			if i == 10{
				break
			}
			//break
		}
	}
	return res
}


func SumInRange(nums []int, queries [][]int) int {
	sums := []int{0, nums[0]}
	m := 1000000007
	for i:= 2; i < len(nums); i++{
		fmt.Println(sums[i-1])
		sums = append(sums, ((sums[i-1] + nums[i-1]) % m))
	}
	sum := 0


	for _, sl := range queries{
		fmt.Println(sl)
		fmt.Println(sl[1]+1)
		fmt.Println(sl[0])
		fmt.Println("len", len(sums))
		//sum += (sums[sl[1]+1] - sums[sl[0]]) % m

	}

	if sum < 0{
		return 1000000007 + sum
	}
	return sum
}


func getAllBranches(t *Tree, l *[]int64, val int64){
	if t == nil{
		return
	}
	val = val * 10 + int64(t.Value.(float64))
	fmt.Println("-", val)
	fmt.Println("*", int64(t.Value.(float64)))
	fmt.Println("#", t.Value.(float64))
	if t.Left == nil && t.Right == nil{
		*l = append(*l, val)
		return
	}
	getAllBranches(t.Right, l, val)
	getAllBranches(t.Left, l, val)
}

func DigitTreeSum(t *Tree) int64 {
	numbers := []int64{}
	getAllBranches(t, &numbers, 0)
	fmt.Println(numbers)
	var res int64 = 0
	for _, n := range numbers{
		res += n
	}
	return res

}