<!-- Credit to Rob M Andrews for his work in this Codepen: https://codepen.io/rob-m-andrews/pen/pvVjNB -->
<script context="module">
	const defaults = {
		MAX_BUBBLE_SIZE: 100,
		MAX_BLUR: 10,
		MAX_ANIMATION_DURATION: 10000,
	};

	const BUBBLE_COUNT = 20
</script>

<script>
	const xPos = Math.floor(Math.random() * 100);
	const yPos = Math.floor(Math.random() * 100);
	const bubbleSize = Math.floor(Math.random() * defaults.MAX_BUBBLE_SIZE + 1);
	const bubbleDelay = Math.floor(Math.random() * 100 * BUBBLE_COUNT + 1);
	const bubbleScale = bubbleSize / defaults.MAX_BUBBLE_SIZE;
	const bubbleBlur = bubbleScale * defaults.MAX_BLUR;
	const bubbleSpeed = (1 - bubbleScale * bubbleScale) * defaults.MAX_ANIMATION_DURATION;
	const bubbleOpacity = 0.25 * (1 - bubbleScale);

	const style = `
		width: ${bubbleSize}px;
		height: ${bubbleSize}px;
		left: ${xPos}vw;
		bottom: ${yPos - 100}vh;
		animation-delay: ${bubbleDelay}ms;
		animation-duration: ${bubbleSpeed}ms;
		filter: blur(${bubbleBlur}px);
		opacity: ${bubbleOpacity};
	`
</script>

<svg class="bubble" viewBox="0 0 100 100" {style} xmlns="http://www.w3.org/2000/svg">
  <circle cx="50" cy="50" r="50" />
</svg>

<style>
	@keyframes bubble {
		from {
			transform: translateY(200vh);
		}
		to {
			transform: translateY(-200vh);
		}
	}

	.bubble {
		animation-duration: 10s;
		animation-name: bubble;
		animation-iteration-count: infinite;
		animation-timing-function: linear;

		opacity: 0.5;
		fill: white;
		filter: blur(2px);

		position: absolute;
		height: 20px;
		width: 20px;
	}
</style>
