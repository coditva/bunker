package lib

type Args map[string]string

func (args *Args) Value(label string) string {
    if val, ok := (*args)[label]; ok {
        return val
    }
    return ""
}
