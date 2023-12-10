// import { VAPID_PUBLIC_KEY } from '$env/static/public'

// This looks weird for backwards compatibility: https://web.dev/push-notifications-subscribing-a-user/#requesting-permission
// export function askPermission () {
// 	return new Promise(function (resolve, reject) {
// 		const permissionResult = Notification.requestPermission(function (result) {
// 			resolve(result);
// 		});

// 		if (permissionResult) {
// 			permissionResult.then(resolve, reject);
// 		}
// 	}).then(function (permissionResult) {
// 		if (permissionResult !== 'granted') {
// 			throw new Error("We weren't granted permission.");
// 		}
// 	});
// }

function askPermission () {
	return Notification.requestPermission()
}

function getExistingSubscription () {
	return navigator.serviceWorker.ready.then((registration) => {
		return registration.pushManager.getSubscription()
	});
}

function subscribe () {
	return navigator.serviceWorker
		.register('/push-notification-service-worker.js')
		.then((registration) => {
			const subscribeOptions = {
				userVisibleOnly: true,
				applicationServerKey: 'BIJTw4YGppJBHll2m7OTLZ1Sj44nlAmwgZt-nr3qbciS_QsbSNlGRh1fjpMjmWyZ8hpeBa0w2wUIJMb3LNBP0K4',
			};

			return registration.pushManager.subscribe(subscribeOptions);
		})
		.then((pushSubscription) => {
			console.info(
				'Subscribed to push notifications.'
			);
			return pushSubscription;
		});
}

function unsubscribe () {
	return navigator.serviceWorker.ready.then((registration) => {
		registration.pushManager.getSubscription().then((subscription) => {
			subscription
				.unsubscribe()
				.then((successful) => {
					console.info('Unsubscribed from push notifications.')
				})
		});
	});
}

export default {
	askPermission,
	getExistingSubscription,
	subscribe,
	unsubscribe,
}
