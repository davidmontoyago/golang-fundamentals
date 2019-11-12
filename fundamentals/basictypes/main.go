package main

import "fmt"

func main() {
	// Slices
	fmt.Println("Slices don't hold data. Are pointers to arrays like:")
	elems := make([]string, 2)
	elems[0] = "Hello"
	elems[1] = "World!"
	printSlice("elems", elems)

	fmt.Println("Slices is how you create dinamically sized arrays:")
	moreElems := make([]string, 1)
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated.
	// The returned slice will point to the newly allocated array.
	moreElems = append(moreElems, "Fire")
	moreElems = append(moreElems, "Water")
	moreElems = append(moreElems, "Wind")
	moreElems = append(moreElems, "Earth")
	printSlice("moreElems", moreElems)

	// Maps
	fmt.Println("The make function returns a map of the given type, initialized and ready for use.")
	kv := make(map[string]string)
	kv["first"] = "Hello"
	kv["second"] = "World!"
	fmt.Printf("%v\n", kv)

	kv["third"] = "Is there anybody in there?"
	third := kv["third"]
	delete(kv, "third")
	fmt.Println(third)
}

func printSlice(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
