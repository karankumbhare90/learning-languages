package flatten

func Flatten(nested interface{}) []interface{} {
    result := make([]interface{}, 0) // ensures non-nil slice

    var helper func(v interface{})
    helper = func(v interface{}) {
        if v == nil {
            return
        }

        switch x := v.(type) {
        case []interface{}:
            for _, elem := range x {
                helper(elem)
            }

        default:
            result = append(result, x)
        }
    }

    helper(nested)
    return result
}
