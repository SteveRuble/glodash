package glodash

/**
* This method is like `_.update` except that it accepts `customizer` which is
* invoked to produce the objects of `path`.  If `customizer` returns `undefined`
* path creation is handled by the method instead. The `customizer` is invoked
* with three arguments: (nsValue, key, nsObject).
*
* **Note:** This method mutates `object`.
*
* @static
* @memberOf _
* @since 4.6.0
* @category Object
* @param {Object} object The object to modify.
* @param {Array|string} path The path of the property to set.
* @param {Function} updater The function to produce the updated value.
* @param {Function} [customizer] The function to customize assigned values.
* @returns {Object} Returns `object`.
* @example
*
* var object = {};
*
* _.updateWith(object, '[0][1]', _.constant('a'), Object);
* // => { '0': { '1': 'a' } }
*/
func UpdateWith[T any](object T, path any /*Array|string*/, updater func(T) bool, customizer func(T) bool) T {
	var out T
	return out
}
	