<script context="module">
  import Bubble from "./Bubble.svelte"
</script>

<div class="bubbles">
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
  <Bubble />
</div>

<style>
  .bubbles {
    /* The background shouldn't have to go here logically but if it goes in a container with z-index > -1 then it covers the bubbles. */
    background: linear-gradient(
      to right,
      rgba(192, 86, 0, 1) 0%,
      rgba(220, 112, 0, 1) 21%,
      rgba(232, 124, 0, 1) 51%,
      rgba(232, 124, 0, 1) 81%,
      rgba(220, 112, 0, 1) 100%
    );
    position: absolute;
    overflow: hidden;
    height: 100%;
    width: 100%;
    z-index: -1;
  }
</style>
