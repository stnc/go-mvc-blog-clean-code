package stnccollection

//https://golangcode.com/check-if-row-exists-in-slice/
//FindSlice elemnt
func FindSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
