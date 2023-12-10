<script context="module">
  import { theme } from "$lib/theme.store";
  import Promo from "$lib/Promo.svelte";

  export const prerender = true;
</script>

<script>
  $theme.background = "plain";
</script>

<div class="content">
  <slot />
</div>

<Promo />

<style>
  .content {
    max-width: 900px;
    /* font-size: 1rem; */
  }

  .content :global(*) {
    color: black;
  }
  .content :global(h1),
  .content :global(h2),
  .content :global(h3) {
    font-family: "Amatic SC", cursive;
    text-align: left;
  }

  .content :global(h1) {
    font-size: 2.59em;
    /* width: 50%; */
  }
  /* 
  .content :global(.hero) {
    
  } */

  .content :global(h2),
  .content :global(h3) {
    font-family: "Amatic SC", cursive;
    margin: 1rem 0;
    max-width: 65ch;
  }

  .content :global(p) {
    font-family: "Port Lligat Slab";
    margin: 1rem 0;
    max-width: 65ch;
    text-align: left;
    /* letter-spacing: 0.1rem; */
  }

  .content :global(a) {
    font-family: "Port Lligat Slab", serif;
  }

  .content :global(li) {
    font-family: "Port Lligat Slab", serif;
  }
</style>
