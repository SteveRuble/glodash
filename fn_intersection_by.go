package glodash

/**
* This method is like `_.intersection` except that it accepts `iteratee`
* which is invoked for each element of each `arrays` to generate the criterion
* by which they're compared. The order and references of result values are
* determined by the first array. The iteratee is invoked with one argument:
* (value).
*
* @static
* @memberOf _
* @since 4.0.0
* @category Array
* @param {...Array} [arrays] The arrays to inspect.
* @param {Function} [iteratee=_.identity] The iteratee invoked per element.
* @returns {Array} Returns the new array of intersecting values.
* @example
*
* _.intersectionBy([2.1, 1.2], [2.3, 3.4], Math.floor);
* // => [2.1]
*
* // The `_.property` iteratee shorthand.
* _.intersectionBy([{ 'x': 1 }], [{ 'x': 2 }, { 'x': 1 }], 'x');
* // => [{ 'x': 1 }]
*/
func IntersectionBy[T any](arrays any /*...Array*/, iteratee func(T) bool) []T {
	var out []T
	return out
}
	