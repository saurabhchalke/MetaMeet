<script context="module">
  import { goto } from "$app/navigation";
  import Next from "$lib/Next.svelte";
  import { theme } from "$lib/theme.store";
  import { Form, Input, Progress } from "spaper";
  import { onMount } from "svelte";

  export const prerender = true;
</script>

<script>
  let loading;
  // TODO: Find a better way to do this.
  $theme.background = "beer";

  onMount(() => {
    document.getElementById("name").focus();
  });
</script>

<svelte:head>
  <title>The fastest way to find out when your friends are free</title>
</svelte:head>

<h1>The fastest way to find out when your friends are free</h1>
<p>Let's start with your name</p>

<!-- svelte-ignore missing-declaration -->
<Form
  on:submit={async (e) => {
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
  }}
>
  <Input id="name" class="margin-bottom-small" name="name" />
  <Next disabled={loading} />
  <div>
    <Progress
      style={`visibility: ${loading ? "visible" : "hidden"};`}
      infinite
      striped
    />
  </div>
</Form>
