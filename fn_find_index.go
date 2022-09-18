package glodash

/**
* This method is like `_.find` except that it returns the index of the first
* element `predicate` returns truthy for instead of the element itself.
*
* @static
* @memberOf _
* @since 1.1.0
* @category Array
* @param {Array} array The array to inspect.
* @param {Function} [predicate=_.identity] The function invoked per iteration.
* @param {number} [fromIndex=0] The index to search from.
* @returns {number} Returns the index of the found element, else `-1`.
* @example
*
* var users = [
*   { 'user': 'barney',  'active': false },
*   { 'user': 'fred',    'active': false },
*   { 'user': 'pebbles', 'active': true }
* ];
*
* _.findIndex(users, function(o) { return o.user == 'barney'; });
* // => 0
*
* // The `_.matches` iteratee shorthand.
* _.findIndex(users, { 'user': 'fred', 'active': false });
* // => 1
*
* // The `_.matchesProperty` iteratee shorthand.
* _.findIndex(users, ['active', false]);
* // => 0
*
* // The `_.property` iteratee shorthand.
* _.findIndex(users, 'active');
* // => 2
*/
func FindIndex[T any](array []T, predicate func(T) bool, fromIndex int) int {
	var out int
	return out
}
	