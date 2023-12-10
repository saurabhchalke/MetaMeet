/**@type {ServiceWorkerGlobalScope} sw */
const sw = self

sw.addEventListener('install', event => {
	sw.skipWaiting();
	console.info('Installing…');
});

sw.addEventListener('activate', event => {
	console.info('Service worker activated!');
});

sw.addEventListener('push', event => {
	const payload = event.data.json()
	const promiseChain = sw.registration.showNotification(payload.title, {
		badge: '/favicon.png',
		body: payload.body,
		icon: '/favicon.png',
		data: {
			socialUrl: payload.socialUrl
		}
	});
	event.waitUntil(promiseChain);
})

self.addEventListener('notificationclick', (event) => {
	const socialUrl = event.notification.data.socialUrl
	event.notification.close(); // Android needs explicit close.

	event.waitUntil(
		clients.matchAll({ type: 'window' }).then(windowClients => {
			// Check if there is already a window/tab open with the target URL
			for (var i = 0; i < windowClients.length; i++) {
				var client = windowClients[i];
				// If so, just focus it.
				if (client.url === socialUrl && 'focus' in client) {
					return client.focus();
				}
			}
			// If not, then open the target URL in a new window/tab.
			if (clients.openWindow) {
				return clients.openWindow(socialUrl);
			}
		})
	);
});
