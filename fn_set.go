package glodash

/**
* Sets the value at `path` of `object`. If a portion of `path` doesn't exist,
* it's created. Arrays are created for missing index properties while objects
* are created for all other missing properties. Use `_.setWith` to customize
* `path` creation.
*
* **Note:** This method mutates `object`.
*
* @static
* @memberOf _
* @since 3.7.0
* @category Object
* @param {Object} object The object to modify.
* @param {Array|string} path The path of the property to set.
* @param {*} value The value to set.
* @returns {Object} Returns `object`.
* @example
*
* var object = { 'a': [{ 'b': { 'c': 3 } }] };
*
* _.set(object, 'a[0].b.c', 4);
* console.log(object.a[0].b.c);
* // => 4
*
* _.set(object, ['x', '0', 'y', 'z'], 5);
* console.log(object.x[0].y.z);
* // => 5
*/
func Set[T any](object T, path any /*Array|string*/, value any /***/) T {
	var out T
	return out
}
	