const radix = 36

export function toSocialId (databaseId) {
	return BigInt(databaseId).toString(radix)
}

export function toDatabaseId (socialId) {
	return convert(socialId, radix).toString()
}

// Credit goes to Nina Scholz https://stackoverflow.com/a/55646905/3660766
// TODO: Figure out how the hell this works.
function convert(value, radix) {
    return [...value.toString()]
        .reduce((r, v) => r * BigInt(radix) + BigInt(parseInt(v, radix)), BigInt(0));
}
