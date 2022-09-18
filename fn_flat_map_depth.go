package glodash

/**
* This method is like `_.flatMap` except that it recursively flattens the
* mapped results up to `depth` times.
*
* @static
* @memberOf _
* @since 4.7.0
* @category Collection
* @param {Array|Object} collection The collection to iterate over.
* @param {Function} [iteratee=_.identity] The function invoked per iteration.
* @param {number} [depth=1] The maximum recursion depth.
* @returns {Array} Returns the new flattened array.
* @example
*
* function duplicate(n) {
*   return [[[n, n]]];
* }
*
* _.flatMapDepth([1, 2], duplicate, 2);
* // => [[1, 1], [2, 2]]
*/
func FlatMapDepth[T any](collection any /*Array|Object*/, iteratee func(T) bool, depth int) []T {
	var out []T
	return out
}
	