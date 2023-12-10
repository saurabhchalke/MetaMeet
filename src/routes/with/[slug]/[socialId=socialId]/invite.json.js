import faunadb from 'faunadb'
import { toDatabaseId } from '$lib/id'

const q = faunadb.query

const client = new faunadb.Client({
	domain: process.env.FAUNADB_DOMAIN,
	port: process.env.FAUNADB_PORT,
	scheme: process.env.FAUNADB_SCHEME,
	secret: process.env.FAUNADB_SERVER_SECRET,
})

export async function post ({ locals, params, request }) {
	const socialId = params.socialId
	const reference = toDatabaseId(socialId)

	const data = await request.formData()
	const name = data.get('name')


	await client.query(
		q.Update(
			q.Ref(q.Collection('social'), reference),
			{
				data: {
					invitees: {
						[locals.userId]: {
							name
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
