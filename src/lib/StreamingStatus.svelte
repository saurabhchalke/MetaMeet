<script>
	export let status
</script>

<div class="streaming-status">{status}</div>

<style>
	.streaming-status {
		position: fixed;
		top: 1rem;
		right: 1rem;
	}

	@media (max-width: 900px) {
		.streaming-status {
			position: unset;
		}
	}
</style>
