// TODO: Add multiple avatars per letter so that people's have unique ones
const avatars = [
	'🤴',
	'🧙‍♂️',
	'👽',
	'🤖',
	'🤡',
	'👹',
	'👾',
	'👶',
	'💩',
]

function addAvatars (invitees) {
	Object.values(invitees).forEach((invitee, index) => {
		invitee.avatar = avatars[index] || '🤮'
	})
}

export default addAvatars
