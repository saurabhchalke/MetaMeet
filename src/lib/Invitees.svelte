<script context="module">
	import addAvatars from '$lib/addAvatars'
</script>

<script>
	export let invitees
	export let selected

	let inviteeArray
	$: {
		addAvatars(invitees)
		inviteeArray = Object.entries(invitees)
	}
</script>

<div class="invitees">
	{#each inviteeArray as [inviteeId, invitee], i}
		<span class="invitee"
			on:mouseenter={() => {
				selected = inviteeId
			}}
			on:mouseleave={() => {
				selected = null
			}}
		>
			<div class="name">{invitee.name}</div>
			<div class="avatar">{invitee.avatar}</div>
		</span>
	{/each}
</div>

<style>
	.invitees {
		text-align: center;
	}

	.invitee {
		display: inline-block;
		margin: 1rem;
		text-align: center;
		width: 3rem;
	}

	.avatar {
		text-align: center;
	}
</style>