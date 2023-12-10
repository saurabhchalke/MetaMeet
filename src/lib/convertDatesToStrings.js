function convertDatesToStrings (social) {
	Object.values(social.invitees)
		.forEach(invitee => {
			invitee.dates = invitee.dates ? invitee.dates.map(date => date.value) : []
		})
	return social
}

export default convertDatesToStrings
