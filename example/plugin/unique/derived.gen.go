// Code generated by goderive DO NOT EDIT.

package unique

// deriveUnique returns a list containing only the unique items from the input list.
// It does this by reusing the input list.
func deriveUnique(list []*Visitor) []*Visitor {
	if len(list) == 0 {
		return nil
	}
	u := 1
	for i := 1; i < len(list); i++ {
		if !deriveContains(list[:u], list[i]) {
			if i != u {
				list[u] = list[i]
			}
			u++
		}
	}
	return list[:u]
}

// deriveContains returns whether the item is contained in the list.
func deriveContains(list []*Visitor, item *Visitor) bool {
	for _, v := range list {
		if deriveEqual(v, item) {
			return true
		}
	}
	return false
}

// deriveEqual returns whether this and that are equal.
func deriveEqual(this, that *Visitor) bool {
	return (this == nil && that == nil) ||
		this != nil && that != nil &&
			((this.UserName == nil && that.UserName == nil) || (this.UserName != nil && that.UserName != nil && *(this.UserName) == *(that.UserName))) &&
			this.RemoteAddr == that.RemoteAddr
}
