package glodash

/**
* Creates an array of elements, sorted in ascending order by the results of
* running each element in a collection thru each iteratee. This method
* performs a stable sort, that is, it preserves the original sort order of
* equal elements. The iteratees are invoked with one argument: (value).
*
* @static
* @memberOf _
* @since 0.1.0
* @category Collection
* @param {Array|Object} collection The collection to iterate over.
* @param {...(Function|Function[])} [iteratees=[_.identity]]
*  The iteratees to sort by.
* @returns {Array} Returns the new sorted array.
* @example
*
* var users = [
*   { 'user': 'fred',   'age': 48 },
*   { 'user': 'barney', 'age': 36 },
*   { 'user': 'fred',   'age': 40 },
*   { 'user': 'barney', 'age': 34 }
* ];
*
* _.sortBy(users, [function(o) { return o.user; }]);
* // => objects for [['barney', 36], ['barney', 34], ['fred', 48], ['fred', 40]]
*
* _.sortBy(users, ['user', 'age']);
* // => objects for [['barney', 34], ['barney', 36], ['fred', 40], ['fred', 48]]
*/
func SortBy[T any](collection any /*Array|Object*/, iteratees any /*...(Function|Function[])*/) []T {
	var out []T
	return out
}
	