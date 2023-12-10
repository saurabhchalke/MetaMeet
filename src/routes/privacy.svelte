<script context="module">
  import { theme } from "$lib/theme.store";

  export const prerender = true;
</script>

<script>
  $theme.background = "beer";
</script>

<svelte:head>
  <title>Privacy</title>
</svelte:head>

<h2>Data we keep for 30 days</h2>
<p>We delete this data automatically 30 days after it was last updated.</p>
<p>Examples: Name, available dates, date of social.</p>

<h2>Data we keep indefinitely</h2>
<p>
  We collect anonymous statistics about our visitors to understand how the
  product is being used. No personally identifiable data is collected.
</p>
<p>
  Examples: How many people have visited a particular page, how many people
  clicked a particular button.
</p>
