<script context="module">
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import Next from "$lib/Next.svelte";
  import { Form, Input, Progress } from "spaper";
  import { theme } from "$lib/theme.store";

  export const prerender = true;
</script>

<script>
  let loading;
  $theme.background = "beer";

  async function handleFormSubmit(e) {
    loading = true;

    try {
      const formData = new FormData(e.target);
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
      });
    });

    // Ensure the first image always has a glow
    images[0].classList.add("glow");
  });
</script>

<svelte:head>
  <title>The fastest way to meet your friends on the Monaverse</title>
</svelte:head>

<h1>The fastest way to meet your friends on the Monaverse</h1>
<p>Let's start with your name</p>

<Form on:submit={handleFormSubmit}>
  <Input id="name" class="margin-bottom-small" name="name" />
  <Next disabled={loading} />
  <Progress
    style={`visibility: ${loading ? "visible" : "hidden"}`}
    infinite
    striped
  />
</Form>

<p>Select the Monaverse space</p>

<div class="image-container">
  <div class="image-wrapper">
    <a
      href="https://monaverse.com/spaces/neon-city-streets/details"
      target="_blank"
    >
      <img src="/neon-city.png" alt="Space 1" class="glow" />
    </a>
    <p class="image-text">Neon City</p>
  </div>
  <div class="image-wrapper">
    <a
      href="https://monaverse.com/spaces/caelestia:-the-lost-fields/details"
      target="_blank"
    >
      <img src="/temple-garden.png" alt="Space 1" class="glow" />
    </a>
    <p class="image-text">Temple Garden</p>
  </div>
  <div class="image-wrapper">
    <a
      href="https://monaverse.com/spaces/caelestia:-the-lost-fields/details"
      target="_blank"
    >
      <img src="/caelestia.png" alt="Space 1" class="glow" />
    </a>
    <p class="image-text">Caelestia</p>
  </div>
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
