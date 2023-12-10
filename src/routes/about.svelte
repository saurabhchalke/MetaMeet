<script context="module">
  import { theme } from "$lib/theme.store";
  import { browser, dev } from "$app/env";

  // we don't need any JS on this page, though we'll load
  // it in dev so that we get hot module replacement...
  export const hydrate = dev;

  // ...but if the client-side router is already loaded
  // (i.e. we came here from elsewhere in the app), use it
  export const router = browser;

  // since there's no dynamic data here, we can prerender
  // it so that it gets served as a static asset in prod
  export const prerender = true;
</script>

<script>
  $theme.background = "beer";
</script>

<svelte:head>
  <title>About</title>
</svelte:head>

<h1>The fastest way to find out when your friends are free</h1>
<p>
  Trying to find a day when friends are free to meet can be a painful process,
  especially over text. We save you time by collecting everyone's availability
  and presenting you with the best dates with a minimum of fuss.
</p>

<!-- TODO: Add product hunt widgets -->
