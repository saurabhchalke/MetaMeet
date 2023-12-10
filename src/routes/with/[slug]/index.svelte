<script context="module">
	import {goto} from '$app/navigation'
	import Onward from '$lib/Next.svelte'

	export const prerender = true
</script>

<svelte:head>
	<title>The fastest way to find out when your friends are free</title>
</svelte:head>

<h2>Schedule your next social with</h2>
<h2 id="sponsor">Odin's Ale House 🍻</h2>
<p>Let's start with your name</p>
<!-- svelte-ignore missing-declaration -->
<form
	on:submit|preventDefault={async e => {
		const formData = new FormData(e.target)
		const name = formData.get('name')

		const response = await fetch(
			'/create.json',
			{
				method: 'POST',
				body: formData
			}
		)

		const socialId = await response.text()
		await goto(`/with/odinsalehouse/${socialId}/you?name=${name}`)
	}}
>
	<!-- TODO: Hardcode a few different placeholder names to select at random. -->
	<input autofocus id="name" name="name"> 
	<Onward/>
</form>
<!-- TODO: Can I do a sweet transition from this page to the next? -->

<style>
	#sponsor {
		font-family: 'Port Lligat Slab', serif;
	}
</style>