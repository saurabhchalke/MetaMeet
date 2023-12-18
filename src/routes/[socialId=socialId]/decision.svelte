<script context="module">
  import Retreat from "$lib/Back.svelte";
  import { spaces } from "$lib/spaces.store";
  import { writable } from "svelte/store";
</script>

<script>
  import { onMount } from "svelte";

  export let social;
  let friendlyDecision;
  $: friendlyDecision = new Date(social.decision).toLocaleDateString(
    undefined,
    {
      weekday: "long",
      month: "long",
      day: "numeric",
    }
  );

  const selectedSpace = writable({ name: "", url: "", image: "" });

  onMount(async () => {
    sessionStorage.setItem("decisionSeen", social.decision);

    // Get the Monaverse space name
    const response = await fetch("./space.json", {
      method: "GET",
    });
    const jsonResponse = await response.json();
    const spaceName = jsonResponse.social.space;

    spaces.subscribe((spaceArray) => {
      const space = spaceArray.find((s) => s.name === spaceName);
      if (space) {
        selectedSpace.set(space);
      }
    });
  });

  function openSpaceUrl(url) {
    window.open(url, "_blank");
  }
</script>

<h1>Your social is on {friendlyDecision}</h1>
{#if $selectedSpace.name}
  <div class="space-details">
    <img
      src={$selectedSpace.image}
      alt={$selectedSpace.name}
      class="space-image"
      on:dblclick={() => openSpaceUrl($selectedSpace.url)}
    />
    <p class="image-text">{$selectedSpace.name}</p>
  </div>
{/if}
<p>
  Have fun 🎉 and if you found this app helpful then please share it with others
  🙏
</p>

<Retreat back="everyone" />

<style>
  .space-details {
    text-align: center;
    margin-top: 20px;
  }

  .space-image {
    width: 250px;
    height: auto;
    border-radius: 10px;
    cursor: pointer;
  }

  .image-text {
    text-align: center;
    margin-top: 10px;
    color: inherit;
    font-size: 1rem;
    font-weight: bold;
  }
</style>
