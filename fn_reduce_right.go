package glodash

/**
* This method is like `_.reduce` except that it iterates over elements of
* `collection` from right to left.
*
* @static
* @memberOf _
* @since 0.1.0
* @category Collection
* @param {Array|Object} collection The collection to iterate over.
* @param {Function} [iteratee=_.identity] The function invoked per iteration.
* @param {*} [accumulator] The initial value.
* @returns {*} Returns the accumulated value.
* @see _.reduce
* @example
*
* var array = [[0, 1], [2, 3], [4, 5]];
*
* _.reduceRight(array, function(flattened, other) {
*   return flattened.concat(other);
* }, []);
* // => [4, 5, 2, 3, 0, 1]
*/
func ReduceRight[T any](collection any /*Array|Object*/, iteratee func(T) bool, accumulator any /***/) any /***/ {
	var out any /***/
	return out
}
	