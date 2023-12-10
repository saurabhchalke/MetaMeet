import test from 'ava'

import {toDatabaseId, toSocialId} from './id.js'

test('Test converting from database ID to social ID and back again', t => {
	const originalDatabaseId = '320160063924732416'

	const socialId = toSocialId(originalDatabaseId)
	const newDatabaseId = toDatabaseId(socialId)

	t.assert(newDatabaseId === originalDatabaseId, `${newDatabaseId} does not equal ${originalDatabaseId}`)
})
