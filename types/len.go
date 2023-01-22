package types

/*
Array: In the case of an array, it returns the number of elements.
Pointer to array: In the case of the pointer to an array, it returns the number of elements in *v (even if v is nil).
Slice, or map: In the case of a slice or a map, it returns the number of elements in v; if v is nil, len(v) is zero.
String: In the case of a string, it returns the number of bytes in v.
Channel: In the case of a channel, it returns the number of elements queued (unread) in the channel buffer;
*/

func lenPrimitive() {
	len(5)
	len(1.0)
	len("Checking")
	len([]int)
	len(map[string]int)
}