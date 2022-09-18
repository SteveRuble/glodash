package glodash

/**
* Creates a slice of `array` from `start` up to, but not including, `end`.
*
* **Note:** This method is used instead of
* [`Array#slice`](https://mdn.io/Array/slice) to ensure dense arrays are
* returned.
*
* @static
* @memberOf _
* @since 3.0.0
* @category Array
* @param {Array} array The array to slice.
* @param {number} [start=0] The start position.
* @param {number} [end=array.length] The end position.
* @returns {Array} Returns the slice of `array`.
*/
func Slice[T any](array []T, start int, end int) []T {
	var out []T
	return out
}
	