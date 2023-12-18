<script context="module">
  import { goto } from "$app/navigation";
  import Onward from "$lib/Next.svelte";
  import { Input } from "spaper";
  import { onMount } from "svelte";
  import { spaces } from "$lib/spaces.store";
  import { writable } from "svelte/store";
</script>

<script>
  const selectedSpace = writable({ name: "", url: "", image: "" });

  onMount(async () => {
    document.getElementById("name").focus();

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

<svelte:head>
  <title>Invite</title>
</svelte:head>

<h1>You've been invited to a Space</h1>

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

<p>Let's start with your name</p>

<form
  on:submit|preventDefault={async (e) => {
    const formData = new FormData(e.target);
    const name = formData.get("name");

    await fetch("./invite", {
      method: "POST",
      body: formData,
    });

    await goto(`you?name=${name}`);
  }}
>
  <Input id="name" name="name" required />
  <Onward />
</form>

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
