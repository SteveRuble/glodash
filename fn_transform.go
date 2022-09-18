package glodash

/**
* An alternative to `_.reduce`; this method transforms `object` to a new
* `accumulator` object which is the result of running each of its own
* enumerable string keyed properties thru `iteratee`, with each invocation
* potentially mutating the `accumulator` object. If `accumulator` is not
* provided, a new object with the same `[[Prototype]]` will be used. The
* iteratee is invoked with four arguments: (accumulator, value, key, object).
* Iteratee functions may exit iteration early by explicitly returning `false`.
*
* @static
* @memberOf _
* @since 1.3.0
* @category Object
* @param {Object} object The object to iterate over.
* @param {Function} [iteratee=_.identity] The function invoked per iteration.
* @param {*} [accumulator] The custom accumulator value.
* @returns {*} Returns the accumulated value.
* @example
*
* _.transform([2, 3, 4], function(result, n) {
*   result.push(n *= n);
*   return n % 2 == 0;
* }, []);
* // => [4, 9]
*
* _.transform({ 'a': 1, 'b': 2, 'c': 1 }, function(result, value, key) {
*   (result[value] || (result[value] = [])).push(key);
* }, {});
* // => { '1': ['a', 'c'], '2': ['b'] }
*/
func Transform[T any](object T, iteratee func(T) bool, accumulator any /***/) any /***/ {
	var out any /***/
	return out
}
	