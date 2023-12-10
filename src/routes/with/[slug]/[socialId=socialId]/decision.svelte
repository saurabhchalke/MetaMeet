<script context="module">
	import Retreat from "$lib/Back.svelte"
</script>

<script>
	export let social
	let decision
	$: decision = new Date(social.decision).toLocaleDateString(undefined, {
		weekday: 'long',
		month: 'long',
		day: 'numeric',
	})
</script>

<h2>You're social is on {decision}</h2>
<h2 id="sponsor">Odin's Ale House 🍻</h2>
<p>Would love to host you! Click <a href="http://example.com/">here to book a table.</a></p>

<Retreat back="everyone"/>

<style>
	#sponsor {
		font-family: 'Port Lligat Slab', serif;
	}
</style>