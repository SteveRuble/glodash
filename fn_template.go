package glodash

/**
* Creates a compiled template function that can interpolate data properties
* in "interpolate" delimiters, HTML-escape interpolated data properties in
* "escape" delimiters, and execute JavaScript in "evaluate" delimiters. Data
* properties may be accessed as free variables in the template. If a setting
* object is given, it takes precedence over `_.templateSettings` values.
*
* **Note:** In the development build `_.template` utilizes
* [sourceURLs](http://www.html5rocks.com/en/tutorials/developertools/sourcemaps/#toc-sourceurl)
* for easier debugging.
*
* For more information on precompiling templates see
* [lodash's custom builds documentation](https://lodash.com/custom-builds).
*
* For more information on Chrome extension sandboxes see
* [Chrome's extensions documentation](https://developer.chrome.com/extensions/sandboxingEval).
*
* @static
* @since 0.1.0
* @memberOf _
* @category String
* @param {string} [string=''] The template string.
* @param {Object} [options={}] The options object.
* @param {RegExp} [options.escape=_.templateSettings.escape]
*  The HTML "escape" delimiter.
* @param {RegExp} [options.evaluate=_.templateSettings.evaluate]
*  The "evaluate" delimiter.
* @param {Object} [options.imports=_.templateSettings.imports]
*  An object to import into the template as free variables.
* @param {RegExp} [options.interpolate=_.templateSettings.interpolate]
*  The "interpolate" delimiter.
* @param {string} [options.sourceURL='lodash.templateSources[n]']
*  The sourceURL of the compiled template.
* @param {string} [options.variable='obj']
*  The data object variable name.
* @param- {Object} [guard] Enables use as an iteratee for methods like `_.map`.
* @returns {Function} Returns the compiled template function.
* @example
*
* // Use the "interpolate" delimiter to create a compiled template.
* var compiled = _.template('hello <%= user %>!');
* compiled({ 'user': 'fred' });
* // => 'hello fred!'
*
* // Use the HTML "escape" delimiter to escape data property values.
* var compiled = _.template('<b><%- value %></b>');
* compiled({ 'value': '<script>' });
* // => '<b>&lt;script&gt;</b>'
*
* // Use the "evaluate" delimiter to execute JavaScript and generate HTML.
* var compiled = _.template('<% _.forEach(users, function(user) { %><li><%- user %></li><% }); %>');
* compiled({ 'users': ['fred', 'barney'] });
* // => '<li>fred</li><li>barney</li>'
*
* // Use the internal `print` function in "evaluate" delimiters.
* var compiled = _.template('<% print("hello " + user); %>!');
* compiled({ 'user': 'barney' });
* // => 'hello barney!'
*
* // Use the ES template literal delimiter as an "interpolate" delimiter.
* // Disable support by replacing the "interpolate" delimiter.
* var compiled = _.template('hello ${ user }!');
* compiled({ 'user': 'pebbles' });
* // => 'hello pebbles!'
*
* // Use backslashes to treat delimiters as plain text.
* var compiled = _.template('<%= "\\<%- value %\\>" %>');
* compiled({ 'value': 'ignored' });
* // => '<%- value %>'
*
* // Use the `imports` option to import `jQuery` as `jq`.
* var text = '<% jq.each(users, function(user) { %><li><%- user %></li><% }); %>';
* var compiled = _.template(text, { 'imports': { 'jq': jQuery } });
* compiled({ 'users': ['fred', 'barney'] });
* // => '<li>fred</li><li>barney</li>'
*
* // Use the `sourceURL` option to specify a custom sourceURL for the template.
* var compiled = _.template('hello <%= user %>!', { 'sourceURL': '/basic/greeting.jst' });
* compiled(data);
* // => Find the source of "greeting.jst" under the Sources tab or Resources panel of the web inspector.
*
* // Use the `variable` option to ensure a with-statement isn't used in the compiled template.
* var compiled = _.template('hi <%= data.user %>!', { 'variable': 'data' });
* compiled.source;
* // => function(data) {
* //   var __t, __p = '';
* //   __p += 'hi ' + ((__t = ( data.user )) == null ? '' : __t) + '!';
* //   return __p;
* // }
*
* // Use custom template delimiters.
* _.templateSettings.interpolate = /{{([\s\S]+?)}}/g;
* var compiled = _.template('hello {{ user }}!');
* compiled({ 'user': 'mustache' });
* // => 'hello mustache!'
*
* // Use the `source` property to inline compiled templates for meaningful
* // line numbers in error messages and stack traces.
* fs.writeFileSync(path.join(process.cwd(), 'jst.js'), '\
*   var JST = {\
*     "main": ' + _.template(mainText).source + '\
*   };\
* ');
 */
func Template[T any](str string) (func(T) (string, error), error) {
	return func(t T) (string, error) {
		return str, nil
	}, nil
}
