package glodash

/**
* Assigns own enumerable string keyed properties of source objects to the
* destination object. Source objects are applied from left to right.
* Subsequent sources overwrite property assignments of previous sources.
*
* **Note:** This method mutates `object` and is loosely based on
* [`Object.assign`](https://mdn.io/Object/assign).
*
* @static
* @memberOf _
* @since 0.10.0
* @category Object
* @param {Object} object The destination object.
* @param {...Object} [sources] The source objects.
* @returns {Object} Returns `object`.
* @see _.assignIn
* @example
*
* function Foo() {
*   this.a = 1;
* }
*
* function Bar() {
*   this.c = 3;
* }
*
* Foo.prototype.b = 2;
* Bar.prototype.d = 4;
*
* _.assign({ 'a': 0 }, new Foo, new Bar);
* // => { 'a': 1, 'c': 3 }
*/
func Assign[T any](object T, sources any /*...Object*/) T {
	var out T
	return out
}
	