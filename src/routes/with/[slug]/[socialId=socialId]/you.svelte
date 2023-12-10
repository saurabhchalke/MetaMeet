<script>
	import Datepicker from "$lib/Datepicker.svelte"
	import {goto} from '$app/navigation'
	import Next from '$lib/Next.svelte'
	import {page} from '$app/stores'

	export let user
	let name = $page.url.searchParams.get('name') || user.name
	let dates = user.dates.map(date => new Date(date).getTime()) || []
</script>

<svelte:head>
	<title>You</title>
</svelte:head>


<!-- TODO: Should I use stores or session here instead of the query param? -->
<!-- Should use load function to support rendering existing selection here. -->
<!-- Need to add ability to edit name here too because pre-rendered page won't show data if for some reason the user comes back here. -->
<h1>Hello {name}, when are you free?</h1>
<!-- svelte-ignore missing-declaration -->
<form
	id="you"
	on:submit|preventDefault={async _ => {
		await fetch(
			'you',
			{
				method: 'PUT',
				body: JSON.stringify({
					dates,
				}),
			}
		)
		goto('everyone')
	}}
>
	<Datepicker bind:selected={dates}/>
</form>
<Next form="you"/>
