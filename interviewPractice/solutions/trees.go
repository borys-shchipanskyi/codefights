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


func isEquals(t1 *Tree, t2 *Tree) bool {
	if t1 == nil && t2 == nil{
		return true
	}
	if t1 != nil && t2 != nil{
		return t1.Value == t2.Value && isEquals(t1.Left, t2.Left) && isEquals(t1.Right, t2.Right)
	}
	return false
}
func IsSubtree(t1 *Tree, t2 *Tree) bool {
	if (t1 == nil && t2 != nil){
		return false
	}
	if isEquals(t1, t2){
		return true
	}
	return IsSubtree(t1.Right, t2) || IsSubtree(t1.Left, t2)

}
