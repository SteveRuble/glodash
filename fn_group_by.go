package glodash

/**
* Creates an object composed of keys generated from the results of running
* each element of `collection` thru `iteratee`. The order of grouped values
* is determined by the order they occur in `collection`. The corresponding
* value of each key is an array of elements responsible for generating the
* key. The iteratee is invoked with one argument: (value).
*
* @static
* @memberOf _
* @since 0.1.0
* @category Collection
* @param {Array|Object} collection The collection to iterate over.
* @param {Function} [iteratee=_.identity] The iteratee to transform keys.
* @returns {Object} Returns the composed aggregate object.
* @example
*
* _.groupBy([6.1, 4.2, 6.3], Math.floor);
* // => { '4': [4.2], '6': [6.1, 6.3] }
*
* // The `_.property` iteratee shorthand.
* _.groupBy(['one', 'two', 'three'], 'length');
* // => { '3': ['one', 'two'], '5': ['three'] }
*/
func GroupBy[T any](collection any /*Array|Object*/, iteratee func(T) bool) T {
	var out T
	return out
}
	