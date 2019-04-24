package bunker

// Args is a map of arguments that are passed to the structs that implement
// Command interface.
type Args map[string]string

// Returns the value of a key or empty string if it does not exist
func (args *Args) Value(label string) string {
    if val, ok := (*args)[label]; ok {
        return val
    }
    return ""
}
