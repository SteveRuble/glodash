package glodash

/**
* This method is like `_.findIndex` except that it iterates over elements
* of `collection` from right to left.
*
* @static
* @memberOf _
* @since 2.0.0
* @category Array
* @param {Array} array The array to inspect.
* @param {Function} [predicate=_.identity] The function invoked per iteration.
* @param {number} [fromIndex=array.length-1] The index to search from.
* @returns {number} Returns the index of the found element, else `-1`.
* @example
*
* var users = [
*   { 'user': 'barney',  'active': true },
*   { 'user': 'fred',    'active': false },
*   { 'user': 'pebbles', 'active': false }
* ];
*
* _.findLastIndex(users, function(o) { return o.user == 'pebbles'; });
* // => 2
*
* // The `_.matches` iteratee shorthand.
* _.findLastIndex(users, { 'user': 'barney', 'active': true });
* // => 0
*
* // The `_.matchesProperty` iteratee shorthand.
* _.findLastIndex(users, ['active', false]);
* // => 2
*
* // The `_.property` iteratee shorthand.
* _.findLastIndex(users, 'active');
* // => 0
*/
func FindLastIndex[T any](array []T, predicate func(T) bool, fromIndex int) int {
	var out int
	return out
}
	