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
	for _, value := range table.values[keyIndex]{
		res = append(res, value)
	}

	return table.values[keyIndex]
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
