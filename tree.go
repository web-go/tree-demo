package main

import "sort"

func tree(list []Category) []Category {
	data := buildData(list)
	result := makeTreeCore(0, data, 0)
	sort.Stable(result)
	return result
	// body, err := json.Marshal(result)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// return string(body)
}

type categorySlice []Category

func (s categorySlice) Len() int           { return len(s) }
func (s categorySlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s categorySlice) Less(i, j int) bool { return s[i].ID > s[j].ID }

func buildData(list categorySlice) map[int]map[int]Category {
	var data map[int]map[int]Category = make(map[int]map[int]Category)
	for _, v := range list {
		id := v.ID
		fid := v.ParentID
		if _, ok := data[fid]; !ok {
			data[fid] = make(map[int]Category)
		}
		data[fid][id] = v
	}
	return data
}

func makeTreeCore(index int, data map[int]map[int]Category, lev int) categorySlice {

	tmp := make(categorySlice, 0)
	for id, item := range data[index] {

		if data[id] != nil {
			item.Children = makeTreeCore(id, data, lev+1)
		}
		item.Level = lev

		tmp = append(tmp, item)
	}
	sort.Stable(tmp)
	return tmp
}
