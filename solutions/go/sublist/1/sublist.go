package sublist

const (
    Equal           Relation = "equal"
    SubListRelation Relation = "sublist"
    Superlist       Relation = "superlist"
    Unequal         Relation = "unequal"
)

func Sublist(l1, l2 []int) Relation {
    if equal(l1, l2) {
        return Equal
    }

    if isSublist(l2, l1) {
        return Superlist
    }

    if isSublist(l1, l2) {
        return SubListRelation
    }

    return Unequal
}

func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func isSublist(a, b []int) bool {
    if len(a) == 0 {
        return true
    }
    if len(a) > len(b) {
        return false
    }

    for i := 0; i <= len(b)-len(a); i++ {
        if equal(a, b[i:i+len(a)]) {
            return true
        }
    }

    return false
}
