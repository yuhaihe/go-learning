/*
可以使用内建函数 make 或使用 map 关键字来定义 Map:
使用 make 函数
map_variable := make(map[KeyType]ValueType, initialCapacity)
*/
package main

import "fmt"

func main() {
	return
	// 创建一个初始容量为 10 的 Map
	// m := make(map[string]int, 10)

	// 使用字面量创建 Map
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(m)

	// 获取键值对
	v1, o1 := m["apple"]
	v2, ok := m["pear"] // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
	fmt.Println(v1, o1, v2, ok)

	m["apple"] = 5
	fmt.Println(m)

	// 获取 Map 的长度
	len := len(m)
	fmt.Println(len)

	// 遍历 Map
	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	// 删除键值对
	delete(m, "banana")
	fmt.Println(m)
}

func init() {
	var siteMap map[string]string /*创建集合 */
	siteMap = make(map[string]string)

	/* map 插入 key - value 对,各个国家对应的首都 */
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"

	/*使用键输出地图值 */
	for site := range siteMap {
		fmt.Println(site, "首都是", siteMap[site])
	}

	/*查看元素在集合中是否存在 */
	name, ok := siteMap["Facebook"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("Facebook 的 站点是", name)
	} else {
		fmt.Println("Facebook 站点不存在")
	}
	fmt.Println(siteMap)
}
