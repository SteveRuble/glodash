package glodash

/**
* This method is like `_.flatMap` except that it recursively flattens the
* mapped results.
*
* @static
* @memberOf _
* @since 4.7.0
* @category Collection
* @param {Array|Object} collection The collection to iterate over.
* @param {Function} [iteratee=_.identity] The function invoked per iteration.
* @returns {Array} Returns the new flattened array.
* @example
*
* function duplicate(n) {
*   return [[[n, n]]];
* }
*
* _.flatMapDeep([1, 2], duplicate);
* // => [1, 1, 2, 2]
*/
func FlatMapDeep[T any](collection any /*Array|Object*/, iteratee func(T) bool) []T {
	var out []T
	return out
}
	