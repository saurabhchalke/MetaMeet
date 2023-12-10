<script context="module">
  import { Hero } from "spaper";
  import Logo from "./Logo.svelte";
</script>

<Hero type="primary">
  <div slot="title">
    <Logo width="120px" />
    <h1>The fastest way to find out when your friends are free</h1>
  </div>
  <svelte:fragment slot="textLead">
    Schedule Your Social is the social planning app that has no signup, no ads,
    and is completely free. We help you find a date that works for everyone in
    three easy steps.
  </svelte:fragment>
  <ol>
    <li>Enter your name and availability 📅</li>
    <li>Share the link with your friends 🔗</li>
    <li>Choose the best date ✅</li>
  </ol>

  <div>
    <a href="/">Get started!</a>
  </div>
</Hero>

<style>
  h1,
  :global(p),
  li {
    color: black;
    margin: 1rem 0;
    text-align: left;
    max-width: unset;
  }

  :global(p),
  li {
    font-family: "Port Lligat Slab", serif;
  }

  div {
    text-align: center;
  }
  div a {
    font-size: xx-large;
  }

  div[slot="title"] {
    display: flex;
    align-items: center;
  }

  @media (max-width: 500px) {
    /* TODO: Fix this */
    div[slot="title"] {
      display: flex;
      flex-direction: column;
    }
  }
</style>
