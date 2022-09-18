package glodash

/**
* Creates an object composed of the `object` properties `predicate` returns
* truthy for. The predicate is invoked with two arguments: (value, key).
*
* @static
* @memberOf _
* @since 4.0.0
* @category Object
* @param {Object} object The source object.
* @param {Function} [predicate=_.identity] The function invoked per property.
* @returns {Object} Returns the new object.
* @example
*
* var object = { 'a': 1, 'b': '2', 'c': 3 };
*
* _.pickBy(object, _.isNumber);
* // => { 'a': 1, 'c': 3 }
*/
func PickBy[T any](object T, predicate func(T) bool) T {
	var out T
	return out
}
	