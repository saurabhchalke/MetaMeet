<script context="module">
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import Next from "$lib/Next.svelte";
  import { Form, Input, Progress } from "spaper";
  import { theme } from "$lib/theme.store";
  import { writable } from "svelte/store";
  import { spaces } from "$lib/spaces.store";

  export const prerender = true;
  const selectedSpace = writable(null);
</script>

<script>
  let loading;
  $theme.background = "beer";
  let currentSelection = "";

  let spaceArray = [];
  spaces.subscribe((value) => {
    spaceArray = value;
    if (spaceArray.length > 0 && !currentSelection) {
      currentSelection = spaceArray[0].name;
      selectedSpace.set(spaceArray[0]);
    }
  });

  async function handleFormSubmit(e) {
    e.preventDefault();
    loading = true;
    const selectedSpaceName = $selectedSpace ? $selectedSpace.name : "";

    try {
      const formData = new FormData(e.target);
      formData.set("space", selectedSpaceName);
      const name = formData.get("name");

      const response = await fetch("/create.json", {
        method: "POST",
        body: formData,
      });

      const socialId = await response.text();
      await goto(`/${socialId}/you?name=${name}`);
    } finally {
      loading = false;
    }
  }

  function selectSpace(space) {
    selectedSpace.set(space);
    currentSelection = space.name;
  }

  function openSpaceUrl(url) {
    window.open(url, "_blank");
  }

  onMount(() => {
    document.getElementById("name").focus();

    // Add event listeners for hover effect
    const images = document.querySelectorAll(".image-container img");
    images.forEach((img, index) => {
      img.addEventListener("mouseover", () => {
        // Remove glow from all images
        images.forEach((i) => i.classList.remove("glow"));
        // Add glow to the hovered image
        img.classList.add("glow");
        selectSpace(spaceArray[index]);
      });

      img.addEventListener("click", () => {
        selectSpace(spaceArray[index]);
      });

      img.addEventListener("dblclick", () => {
        openSpaceUrl(spaceArray[index].url);
      });
    });

    // Ensure the first image always has a glow
    images.forEach((i) => i.classList.remove("glow"));
    images[0].classList.add("glow");
  });
</script>

<svelte:head>
  <title>The fastest way to meet your friends on the Monaverse</title>
</svelte:head>

<h1>The fastest way to meet your friends on the Monaverse</h1>
<p>Let's start with your name</p>

<Form on:submit={handleFormSubmit}>
  <Input id="name" class="margin-bottom-small" name="name" required />
  <Next disabled={loading} />
  <Progress
    style={`visibility: ${loading ? "visible" : "hidden"}`}
    infinite
    striped
  />
</Form>

{#if currentSelection}
  <p>
    Selected Space: <span style="font-weight: bold; color: #0000EE"
      >{currentSelection}</span
    >
  </p>
{/if}

<div class="image-container">
  {#each spaceArray as { name, image }}
    <div class="image-wrapper">
      <img src={image} alt={name} class="glow" />
      <p class="image-text">{name}</p>
    </div>
  {/each}
</div>

<style>
  form {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
  }

  .image-container {
    display: flex;
    justify-content: center;
    gap: 50px;
    margin-top: 20px;
  }

  .image-container img {
    transition: 0.3s all ease-in-out;
    width: 250px;
    height: auto;
    border-radius: 10px;
    cursor: pointer; /* Indicates the image is interactive */
  }

  .image-container img.glow {
    box-shadow: 0 0 30px rgba(255, 255, 255, 0.8);
    transform: scale(1.05);
  }

  .image-text {
    text-align: center;
    margin-top: 10px;
    color: inherit;
    font-size: 1rem;
    font-weight: bold;
  }
</style>
