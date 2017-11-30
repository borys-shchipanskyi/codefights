package solutions

import "fmt"

type HashTable struct {
	keys []string
	values [][]string
}

func (table *HashTable) getKeyIndex(key string )int{
	for i, value := range table.keys{
		if value == key{
			return i
		}
	}
	return -1
}

func  (table *HashTable) addKey(key string) int{
	table.keys = append(table.keys, key)
	table.values = append(table.values, []string{})
	return len(table.keys) -1
}

func (table *HashTable) Add(key string, value string){
	keyIndex := table.getKeyIndex(key)
	if keyIndex == -1{
		keyIndex = table.addKey(key)
	}
	table.values[keyIndex] = append(table.values[keyIndex], value)
}

func (table *HashTable) GetKeys()[]string{
	res := []string{}
	for _, key := range table.keys{
		res = append(res, key)
	}

	return res
}

func (table *HashTable) GetValues(key string)[]string{
	keyIndex := table.getKeyIndex(key)
	res := []string{}
	if keyIndex == -1{
		return res
	}
	for _, value := range table.values[keyIndex]{
		res = append(res, value)
	}

	return res
}

type HashTableInt struct {
	keys []int
	values [][]int
}

func (table *HashTableInt) getKeyIndex(key int )int{
	for i, value := range table.keys{
		if value == key{
			return i
		}
	}
	return -1
}

func  (table *HashTableInt) addKey(key int) int{
	table.keys = append(table.keys, key)
	table.values = append(table.values, []int{})
	return len(table.keys) -1
}

func (table *HashTableInt) Add(key int, value int){
	keyIndex := table.getKeyIndex(key)
	if keyIndex == -1{
		keyIndex = table.addKey(key)
	}
	table.values[keyIndex] = append(table.values[keyIndex], value)
}

func (table *HashTableInt) GetKeys()[]int{
	res := []int{}
	for _, key := range table.keys{
		res = append(res, key)
	}

	return res
}

func (table *HashTableInt) GetValues(key int)[]int{
	keyIndex := table.getKeyIndex(key)
	res := []int{}
	if keyIndex == -1{
		return res
	}
	for _, value := range table.values[keyIndex]{
		res = append(res, value)
	}

	return res
}



func IsHashTablesEquals(a [][]string, b[][]string)bool{
	if len(a) != len(b){
		return false
	}
	for i:= 0; i < len(a); i++{
		for j := 0; j< len(a[i]); j++{
			if len(a[i]) != len(b[i]){
				return false
			}
			if a[i][j] != b[i][j]{
				return false
			}
		}
	}
	return true
}

func PrintHashTable(a [][]string){
	for i:= 0; i < len(a); i++{
		for j := 0; j< len(a[i]); j++{
			fmt.Printf("%s,", a[i][j])
		}
		fmt.Println()
	}
}

func fillHashTable(a [][]string) *HashTable{
	hashTable := HashTable{}
	for i:= 0; i < len(a); i++{
		dish := a[i][0]
		for j :=1; j< len(a[i]); j++{
			ingredient := a[i][j]
			hashTable.Add(ingredient, dish)
		}
	}
	return &hashTable
}

func sortList(source []string) []string{
	for i:= 0; i < len(source); i++{
		for j:= i+1; j< len(source); j++{
			if source[i] > source[j]{
				tmp := source[i]
				source[i] = source[j]
				source[j] = tmp
			}
		}
	}
	return source
}


func GroupingDishes(dishes [][]string) [][]string {
	hashTable := fillHashTable(dishes)
	res := [][]string{}
	keys := hashTable.GetKeys()

	for _, key := range sortList(keys){
		values :=  hashTable.GetValues(key)
		if len(values) > 1{
			row := []string{key}
			for _, dish := range sortList(values){
				row = append(row, dish)

			}
			res = append(res, row)
		}
	}
	return res
}

func AreFollowingPatterns(strings []string, patterns []string) bool {
	hashTable := HashTable{}
	for i:= 0; i < len(strings); i++{
		key := patterns[i]
		value := strings[i]
		values := hashTable.GetValues(key)
		if len(values) == 0{
			hashTable.Add(key, value)
		} else {
			if values[0] != value{
				return false
			}
		}

	}
	hashTable = HashTable{}
	for i:= 0; i < len(strings); i++{
		value := patterns[i]
		key := strings[i]
		values := hashTable.GetValues(key)
		if len(values) == 0{
			hashTable.Add(key, value)
		} else {
			if values[0] != value{
				return false
			}
		}

	}

	return true
}

func ContainsCloseNums(nums []int, k int) bool {
	hashTable := HashTableInt{}

	for i, n := range nums{
		hashTable.Add(n, i)
	}
	fmt.Printf("%#v\n", hashTable)
	for _, key := range hashTable.GetKeys(){
		values := hashTable.GetValues(key)
		if len(values) <2 {
			continue
		}
		for j := 1; j< len(values); j++{
			if (values[j] - values[j-1]) <= k{
				return true
			}

		}
	}
	return false
}

func PossibleSums(coins []int, quantity []int) int {
	out := map[int]int{}
	res := []int {}
	pos := 0
	for i, val := range coins{
		for j := 1; j <= quantity[i]; j++{
			res = append(res, j*val)
			out[j*val] = -1
			if i > 0 {
				for k := 0; k < pos; k++{
					res = append(res, j*val + res[k])
					out[j*val + res[k]] = -1
				}
			}
		}
		pos = len(res)
	}
	return len(out)
}


// SwapLexOrder

func get_pair_for_index(pairs *[][]int, index int)[]int{
	res := []int{}
	for _, pair := range *pairs{
		if pair[0] == index{
			res = append(res, pair[1])
		}
		if pair[1] == index{
			res = append(res, pair[0])
		}
	}
	return res
}
func isContains( target *[]int, val int) bool{
	for _, el := range *target{
		if el == val{
			return true
		}
	}
	return false
}

func marge_two_slices(slice1 *[]int, slice2 *[]int){
	for _, el := range *slice2{
		if !isContains(slice1, el){
			*slice1 = append(*slice1, el)
		}
	}
}

func find_pairs_for_index(pairs *[][]int, index int)[]int{
	res := []int{index}
	res = append(res, get_pair_for_index(pairs, index)...)
	for i:=1; i<len(res); i++{
		tmp := get_pair_for_index(pairs, res[i])
		marge_two_slices(&res, &tmp)
	}
	return res
}

func generate_indexes_map(pairs *[][]int)map[int][]int{
	res := map[int][]int{}

	for _, pair := range *pairs{
		if _, ok := res[pair[0]]; !ok{
			res[pair[0]] = find_pairs_for_index(pairs, pair[0])
		}
		if _, ok := res[pair[1]]; !ok{
			res[pair[1]] = find_pairs_for_index(pairs, pair[1])
		}
	}
	return res
}

func get_right_letter(usedIndexes *[]int, indexes *[]int, str string)string{
	free_indexes := []int{}
	for _, in := range *indexes{
		if (!isContains(usedIndexes, in)){
			free_indexes = append(free_indexes, in)
		}
	}
	max := str[free_indexes[0]-1]
	usedIndex := free_indexes[0]
	for i:=1; i < len(free_indexes); i++{
		if str[free_indexes[i] - 1] > max{
			max = str[free_indexes[i]-1]
			usedIndex = free_indexes[i]
		}
	}
	*usedIndexes = append(*usedIndexes, usedIndex)
	return string(max)
}

func SwapLexOrder(str string, pairs [][]int) string {
	if len(pairs) < 2{
		return str
	}
	res := ""
	usedIndexes := []int{}
	indexes_map := generate_indexes_map(&pairs)

	for i:= 1; i <= len(str); i++{
		if val, ok := indexes_map[i]; ok{
			if len(val) == 1{
				res += string(str[val[0] -1])
			} else {
				res += get_right_letter(&usedIndexes, &val, str)
			}

		} else {
			res += string(str[i -1])

		}
	}
	return res
}
