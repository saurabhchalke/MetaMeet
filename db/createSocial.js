import faunadb from 'faunadb'
import { toSocialId } from '../src/lib/id.js'

const q = faunadb.query

const client = new faunadb.Client({
	domain: process.env.FAUNADB_DOMAIN,
	port: process.env.FAUNADB_PORT,
	scheme: process.env.FAUNADB_SCHEME,
	secret: process.env.FAUNADB_SERVER_SECRET,
})

function todayPlus(days) {
	return q.TimeAdd(q.ToDate(q.Now()), days, "days")
}

const response = await client.query(
	q.Create(
		q.Collection('social'),
		{
			data: {
				invitees: {
					"cbc9de47-185b-4115-a2a5-e4e8dbe6e249": {
						name: "Dan",
						dates: [
							todayPlus(1),
							todayPlus(5),
							todayPlus(6),
						]
					},
					"d742dc09-c309-4f2c-9b33-c817b51c1537": {
						name: "Max",
						dates: [
							todayPlus(2),
							todayPlus(6),
						]
					},
					"5c15e5ec-3a5c-4518-9224-591635a38bb8": {
						name: "Tom",
						dates: [
							todayPlus(8),
							todayPlus(6),
							todayPlus(10),
						]
					},
					"9d5c15e5-ec3a-4c35-98d2-24591635a38b": {
						name: "King Daddy Sexy Robo III",
						dates: [
							todayPlus(1),
							todayPlus(6),
						]
					},
					"6bb56257-f312-484f-ac7b-458bcaa76b5f": {
						name: "Jim",
						dates: [
							todayPlus(5),
							todayPlus(6),
						]
					},
					"fe9d5c15-e5ec-4a5c-b518-d224591635a3": {
						name: "Mike",
						dates: [
							todayPlus(15),
							todayPlus(3),
							todayPlus(6),
						]
					}
				}
			}
		}
	)
)
const socialId = toSocialId(response.ref.id)
console.info(`http://localhost:6602/${socialId}/everyone`)
