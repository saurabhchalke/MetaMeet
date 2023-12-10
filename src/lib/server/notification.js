import webpush from 'web-push'

export async function send (social, userId, notification) {
	webpush.setVapidDetails(
		'mailto:mrdanielmetcalfe@gmail.com',
		process.env.VAPID_PUBLIC_KEY,
		process.env.VAPID_PRIVATE_KEY,
	)

	const pushSubscriptions = Object.entries(social.invitees)
		.filter(([inviteeId, invitee]) => inviteeId !== userId && invitee.pushSubscriptions)
		.reduce((subscriptions, [_inviteeId, invitee]) => subscriptions.concat(Object.values(invitee.pushSubscriptions)), [])

	await Promise.all(pushSubscriptions.map(pushSubscription => webpush.sendNotification(
		pushSubscription,
		JSON.stringify(notification),
	)))
}

export default {
	send,
}
