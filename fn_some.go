package glodash

/**
* Checks if `predicate` returns truthy for **any** element of `collection`.
* Iteration is stopped once `predicate` returns truthy. The predicate is
* invoked with three arguments: (value, index|key, collection).
*
* @static
* @memberOf _
* @since 0.1.0
* @category Collection
* @param {Array|Object} collection The collection to iterate over.
* @param {Function} [predicate=_.identity] The function invoked per iteration.
* @param- {Object} [guard] Enables use as an iteratee for methods like `_.map`.
* @returns {boolean} Returns `true` if any element passes the predicate check,
*  else `false`.
* @example
*
* _.some([null, 0, 'yes', false], Boolean);
* // => true
*
* var users = [
*   { 'user': 'barney', 'active': true },
*   { 'user': 'fred',   'active': false }
* ];
*
* // The `_.matches` iteratee shorthand.
* _.some(users, { 'user': 'barney', 'active': false });
* // => false
*
* // The `_.matchesProperty` iteratee shorthand.
* _.some(users, ['active', false]);
* // => true
*
* // The `_.property` iteratee shorthand.
* _.some(users, 'active');
* // => true
*/
func Some[T any](collection any /*Array|Object*/, predicate func(T) bool) any /*boolean*/ {
	var out any /*boolean*/
	return out
}
	