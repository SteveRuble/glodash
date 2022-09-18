package glodash

/**
* Creates a function that invokes `func` with the `this` binding of the
* create function and an array of arguments much like
* [`Function#apply`](http://www.ecma-international.org/ecma-262/7.0/#sec-function.prototype.apply).
*
* **Note:** This method is based on the
* [spread operator](https://mdn.io/spread_operator).
*
* @static
* @memberOf _
* @since 3.2.0
* @category Function
* @param {Function} func The function to spread arguments over.
* @param {number} [start=0] The start position of the spread.
* @returns {Function} Returns the new function.
* @example
*
* var say = _.spread(function(who, what) {
*   return who + ' says ' + what;
* });
*
* say(['fred', 'hello']);
* // => 'fred says hello'
*
* var numbers = Promise.all([
*   Promise.resolve(40),
*   Promise.resolve(36)
* ]);
*
* numbers.then(_.spread(function(x, y) {
*   return x + y;
* }));
* // => a Promise of 76
*/
func Spread[T any](fn func(T) bool, start int) func(T) bool {
	var out func(T) bool
	return out
}
	