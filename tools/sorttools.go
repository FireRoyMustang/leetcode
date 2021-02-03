package tools

//  从小到大排序
type byteSlice []byte

func (bytes byteSlice) Len() int {
	return len(bytes)
}

func (bytes byteSlice) Swap(i, j int) {
	bytes[i], bytes[j] = bytes[j], bytes[i]
}

func (bytes byteSlice) Less(i, j int) bool {
	return bytes[i] < bytes[j]
}
