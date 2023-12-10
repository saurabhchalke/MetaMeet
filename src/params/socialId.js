/** @type {import('@sveltejs/kit').ParamMatcher} */
export function match(param) {
	return /^.{12}$/.test(param)
}
