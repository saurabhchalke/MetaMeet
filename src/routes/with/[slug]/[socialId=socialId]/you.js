import faunadb from 'faunadb'
import {toDatabaseId} from '$lib/id'
import convertDatesToStrings from '$lib/convertDatesToStrings'

const q = faunadb.query

const client = new faunadb.Client({
	domain: process.env.FAUNADB_DOMAIN,
	port: process.env.FAUNADB_PORT,
	scheme: process.env.FAUNADB_SCHEME,
	secret: process.env.FAUNADB_SERVER_SECRET,
})

export async function get ({params, locals}) {
	const socialId = params.socialId
	const reference = toDatabaseId(socialId)
	const response = await client.query(
		q.Get(q.Ref(q.Collection('social'), reference))
	)
	const social = convertDatesToStrings(response.data)
	const user = social.invitees[locals.userId]
	
	return {
		status: 200,
		body: {
			social,
			user,
		}
	}
}

export async function put ({locals, params, request}) {
	const socialId = params.socialId
	const reference = toDatabaseId(socialId)

	const {dates} = await request.json()

	await client.query(
		q.Update(
			q.Ref(q.Collection('social'), reference),
			{
				data: {
					invitees: {
						[locals.userId]: {
							dates: dates.map(date => q.Date(new Date(date).toISOString().split('T')[0]))
						}
					}
				}
			}
		)
	)

	return {
		status: 200,
	}
}
