<script context="module">
  import "papercss/dist/paper.min.css";
  import { dev } from "$app/env";
  import Bubbles from "$lib/Bubbles/index.svelte";
  import Footer from "$lib/Footer.svelte";
  import { theme } from "$lib/theme.store";
  import * as Sentry from "@sentry/browser";
  import { BrowserTracing } from "@sentry/tracing";
  import Plausible from "plausible-tracker";
  import { onMount, setContext } from "svelte";

  export const prerender = true;
</script>

<script>
  let plausible;
  setContext("analytics", {
    getAnalytics: () => plausible,
  });

  let background;
  $: background = $theme.background;

  onMount(() => {
    Sentry.init({
      dsn: "https://5b7a4ba6b3ff446ea520d3b5d7b854f2@o1232542.ingest.sentry.io/6380542",
      integrations: [new BrowserTracing()],
      tracesSampleRate: 1.0,
      enabled: !dev,
    });

    // TODO: Raise a PR to define exports properly so this works with Vite.
    plausible = Plausible();
    plausible.enableAutoPageviews();
    plausible.enableAutoOutboundTracking();
  });
</script>

{#if background === "beer"}
  <Bubbles />
{/if}

<div class="main-container">
  <main>
    <slot />
  </main>
  <Footer />
</div>

<style>
  :root {
    --beer: #e87c00;
    --light-beer: rgba(255, 156, 26, 0.836);
    --good: #adff2f;
    --bad: red;
    --white: #fff;
    --purple: rgb(85, 26, 139);
    --pink: rgba(218, 88, 218, 0.747);

    /* Paper CSS overrides */
    --primary: var(--purple);
    --primary-light: var(--light-beer);
    /* --secondary: var(--pink); */
    --success: #608d1c;
    --success-light: var(--good);
    /* --main-background: var(--beer); */
  }

  :global(h1, h2, h3, h4, h5, h6) {
    font-family: "Amatic SC", cursive;
  }

  :global(html, body, button, input, div, span, applet, object, iframe, p, blockquote, pre, a, abbr, acronym, address, big, cite, code, del, dfn, em, img, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var, b, u, i, center, dl, dt, dd, ol, ul, li, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td, article, aside, canvas, details, embed, figure, figcaption, footer, header, hgroup, menu, nav, output, ruby, section, summary, time, mark, audio, video) {
    font-family: "Amatic SC", cursive;
  }

  :global(code) {
    font-family: monospace;
  }

  /* :global(html) {
    font-size: calc(1em + 0.5vw);
  } */

  :global(h1, h2, h3, p, button, li) {
    color: white;
    /* font-size: x-large; */
    text-align: center;
  }

  :global(h1) {
    font-size: 4.17em;
    font-weight: 700;
  }

  @media (max-width: 500px) {
    :global(h1) {
      font-size: 3em;
      font-weight: 700;
    }
  }

  :global(h2) {
    font-size: 2.59em;
  }

  :global(h3) {
    font-size: 1.61em;
  }

  :global(p) {
    font-size: 1em;
    max-width: 900px;
    margin: 1em;
  }

  :global(span) {
    font-size: 1em;
  }

  :global(li) {
    text-align: left;
    margin: 1em;
  }

  :global(button) {
    color: black;
    cursor: pointer;
  }

  :global(button:disabled) {
    filter: saturate(30%);
  }

  :global(input, button) {
    /* border-radius: 0.2em; */
    background-color: white;
    color: black;
    display: inline-block;
    font-size: 1em;
    /* padding: 0.24em; */
  }

  :global(body) {
    position: relative;
  }

  :global(.main-container) {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    padding: 1rem;
    max-width: 100vw;
  }

  main {
    align-items: center;
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    justify-content: center;
    margin: auto;
  }

  /* Paper CSS & overrides */
  /* @import "papercss/dist/paper.min.css"; */

  :global(a) {
    background-image: none;
    text-decoration-color: rgb(0, 0, 238);
    text-decoration-line: underline;
    text-decoration-style: solid;
    text-decoration-thickness: auto;
  }

  :global(.paper-btn, [type="button"], button) {
    padding: 0.5rem;
  }

  :global(h1) {
    margin: 0;
  }

  :global(h2) {
    margin: 0;
  }
</style>
