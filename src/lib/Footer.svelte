<footer>
	<!-- TODO: Use something like `class:active={$page.path === '/about'}` to indicate which page we're currently on. -->
	<span><a sveltekit:prefetch href="/">Schedule another social</a></span>
	<span><a sveltekit:prefetch href="/help">Help</a></span>
	<span><a sveltekit:prefetch href="/about">About</a></span>
	<span><a sveltekit:prefetch href="/privacy">Privacy</a></span>
</footer>

<style>
	footer {
		display: flex;
		justify-content: center;
	}

	footer a {
		font-weight: bold;
	}

	/* Why dis? */
	@media (min-width: 480px) {
		footer {
			padding: 40px 0;
		}
	}

	span {
		margin: 0.5rem;
	}
</style>
