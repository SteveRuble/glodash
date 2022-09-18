package glodash

/**
* Removes the property at `path` of `object`.
*
* **Note:** This method mutates `object`.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Object
* @param {Object} object The object to modify.
* @param {Array|string} path The path of the property to unset.
* @returns {boolean} Returns `true` if the property is deleted, else `false`.
* @example
*
* var object = { 'a': [{ 'b': { 'c': 7 } }] };
* _.unset(object, 'a[0].b.c');
* // => true
*
* console.log(object);
* // => { 'a': [{ 'b': {} }] };
*
* _.unset(object, ['a', '0', 'b', 'c']);
* // => true
*
* console.log(object);
* // => { 'a': [{ 'b': {} }] };
*/
func Unset[T any](object T, path any /*Array|string*/) any /*boolean*/ {
	var out any /*boolean*/
	return out
}
	