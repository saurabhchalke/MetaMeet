<script context="module">
	import addAvatars from '$lib/addAvatars'
</script>

<script>
	export let invitees
	export let inviteesCount
	export let selected
	
	let bestDates
	let inviteesByDate
	$: {
		inviteesCount = Object.values(invitees).length
		addAvatars(invitees)
		inviteesByDate = {}
		for (const [_, invitee] of Object.entries(invitees)) {
			for (const date of invitee.dates || []) {
				const hasDateKey = date in inviteesByDate
				if (!hasDateKey) inviteesByDate[date] = new Set()
				inviteesByDate[date].add(invitee)
			}
		}
		bestDates = Object.entries(inviteesByDate)
			.sort(([dateOne, inviteesOne], [dateTwo, inviteesTwo]) => {
				if (inviteesOne.size < inviteesTwo.size) return 1
				if (inviteesOne.size > inviteesTwo.size) return -1
				if (dateOne > dateTwo) return 1
				if (dateOne < dateTwo) return -1
				return 0
			})
			.slice(0, 10)
	}
</script>

<table
	id="best-dates"
	on:click={event => {
		const tableRow = event.composedPath().find(node => node.tagName === 'TR')
		const radioInput = tableRow.getElementsByTagName('input')[0]
		selected = radioInput.value
	}}>
	<thead>
		<tr>
			<th></th>
			<th>Rank</th>
			<th>Date</th>
			<th>Available Invitees</th>
			<th>Invitees Total</th>
		</tr>
	</thead>
	<tbody>
		{#each bestDates as [date, inviteesForDate], i}
			<tr>
				<td>
					<input bind:group={selected} name="best-dates" type="radio" value={date}/>
				</td>
				<td>{i + 1}</td>
				<td>
					{new Date(date).toLocaleDateString(undefined, {
						weekday: 'long',
						month: 'long',
						day: 'numeric',
					})}
				</td>
				<td>
					{#each Array.from(inviteesForDate) as invitee}
						{invitee.avatar}
					{/each}
				</td>
				<td>{inviteesForDate.size}/{inviteesCount}</td>
			</tr>
		{/each}
	</tbody>
</table>

<style>
	#best-dates {
		font-size: x-large;
		color: white;
		text-align: left;
		font-family: 'Port Lligat Slab', serif;
	}

	#best-dates td {
		font-family: 'Port Lligat Slab', serif;
		font-size: large;
	}

	th, td {
		padding: 0.61rem;
	}

	tbody tr {
		position: relative;
		cursor: pointer;
	}

	@media (min-width: 900px) {
		table thead tr::before {
			/* Hack required otherwise the before pseudo element nudges the whole table over to the right as though its added a new td element. */
			content: ""; 
		}

		tbody tr::before {
			content: "👉";
			position: absolute;
			left: -2rem;
			visibility: hidden;

		}

		tbody tr:hover::before {
			content: "👉";
			visibility: visible;
		}
	}
</style>