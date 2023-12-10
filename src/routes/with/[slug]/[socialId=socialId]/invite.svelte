<script context="module">
	import {goto} from '$app/navigation'
	import Onward from '$lib/Next.svelte'

	export const prerender = true
</script>

<svelte:head>
	<title>Invite</title>
</svelte:head>

<h1>You've been invited to a social</h1>
<p>Let's start with your name</p>
<!-- svelte-ignore missing-declaration -->
<form
	on:submit|preventDefault={async e => {
		const formData = new FormData(e.target)
		const name = formData.get('name')

		await fetch(
			'./invite.json',
			{
				method: 'POST',
				body: formData
			}
		)

		await goto(`you?name=${name}`)
	}}
>
	<!-- TODO: Hardcode a few different placeholder names to select at random. -->
	<input autofocus id="name" name="name"> 
	<Onward/>
</form>
