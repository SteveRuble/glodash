package glodash

/**
* Binds methods of an object to the object itself, overwriting the existing
* method.
*
* **Note:** This method doesn't set the "length" property of bound functions.
*
* @static
* @since 0.1.0
* @memberOf _
* @category Util
* @param {Object} object The object to bind and assign the bound methods to.
* @param {...(string|string[])} methodNames The object method names to bind.
* @returns {Object} Returns `object`.
* @example
*
* var view = {
*   'label': 'docs',
*   'click': function() {
*     console.log('clicked ' + this.label);
*   }
* };
*
* _.bindAll(view, ['click']);
* jQuery(element).on('click', view.click);
* // => Logs 'clicked docs' when clicked.
*/
func BindAll[T any](object T, methodNames any /*...(string|string[])*/) T {
	var out T
	return out
}
	