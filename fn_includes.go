package glodash

/**
* Checks if `value` is in `collection`. If `collection` is a string, it's
* checked for a substring of `value`, otherwise
* [`SameValueZero`](http://ecma-international.org/ecma-262/7.0/#sec-samevaluezero)
* is used for equality comparisons. If `fromIndex` is negative, it's used as
* the offset from the end of `collection`.
*
* @static
* @memberOf _
* @since 0.1.0
* @category Collection
* @param {Array|Object|string} collection The collection to inspect.
* @param {*} value The value to search for.
* @param {number} [fromIndex=0] The index to search from.
* @param- {Object} [guard] Enables use as an iteratee for methods like `_.reduce`.
* @returns {boolean} Returns `true` if `value` is found, else `false`.
* @example
*
* _.includes([1, 2, 3], 1);
* // => true
*
* _.includes([1, 2, 3], 1, 2);
* // => false
*
* _.includes({ 'a': 1, 'b': 2 }, 1);
* // => true
*
* _.includes('abcd', 'bc');
* // => true
*/
func Includes[T any](collection any /*Array|Object|string*/, value any /***/, fromIndex int) any /*boolean*/ {
	var out any /*boolean*/
	return out
}
	