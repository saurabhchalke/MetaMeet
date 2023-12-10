import cookie from 'cookie'
import { v4 as uuid } from '@lukeed/uuid'

export async function handle ({ event, resolve }) {
	const cookieHeader = event.request.headers.get('cookie')
	const cookies = cookie.parse(cookieHeader || '')
	event.locals.userId = cookies.userId || uuid()

	// TODO https://github.com/sveltejs/kit/issues/1046
	if (event.url.searchParams.has('_method')) {
		event.method = event.url.searchParams.get('_method').toUpperCase()
	}

	const response = await resolve(event)
	if (!cookies.userId) {
		// if this is the first time the user has visited this app,
		// set a cookie so that we recognise them when they return
		const userIdCookie = cookie.serialize('userId', event.locals.userId, {
			maxAge: 2592000,
			path: '/',
		})
		response.headers.set('set-cookie', userIdCookie)
	}

	return response;
};

export function getSession (event) {
	return {
		userId: event.locals.userId,
		faunadb: {
			domain: process.env.FAUNADB_DOMAIN,
			port: process.env.FAUNADB_PORT,
			scheme: process.env.FAUNADB_SCHEME,
			secret: process.env.FAUNADB_INVITEE_SECRET,
		},
	}
}
