import cookie from 'cookie'

export default getUserId

function getUserId () {
	const {userId} = cookie.parse(document.cookie)
	if (!userId) throw new Error("No user ID cookie set.")
	return userId
}
