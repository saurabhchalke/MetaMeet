<script context="module">
  import { theme } from "$lib/theme.store";

  export const prerender = true;
</script>

<script>
  $theme.background = "beer";
</script>

<svelte:head>
  <title>Help</title>
</svelte:head>

<p>
  If you need help with anything or just want to provide feedback then please
  <a
    href="https://github.com/moderatemisbehaviour/scheduleyoursocial/issues/new"
    target="_blank">raise an issue on GitHub.</a
  >
</p>
