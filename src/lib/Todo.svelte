<script context="module">
  import { CopyButton, Hero } from "spaper";
</script>

<script>
  export let items;
  const copyableItems = items.map((item, i) => `${i + 1}. ${item}`).join("\n");

  async function share() {
    try {
      await navigator.share({
        title: "Board game night todo list",
        text: copyableItems,
      });
      // trackShareInviteLink("native");
    } catch (error) {
      console.warn("Unable to use native sharing", error);
      if (error.name === "AbortError") return; // If we try to fallback to 'copy' after this it results in a NotAllowedErorr. I'm guesing this is because when the user aborts share it resets Safari's concept of if is a valid and safe clipboard operation triggered by a user interaction (see https://webkit.org/blog/10247/new-webkit-features-in-safari-13-1/).

      try {
        await copy(copyableItems);
        // trackShareInviteLink("auto copy with dialog");
      } catch (error) {
        // One example of when this would happen is in some Android Webviews that don't provide copy permissions.
        console.warn("Unable to copy to clipboard", error);
        // trackShareInviteLink("manual copy with dialog");
      }
    }
  }

  async function copy(text) {
    if (!navigator.clipboard)
      throw new Error("Navigator clipboard not available");
    await navigator.clipboard.writeText(text);
  }
</script>

<Hero
  title="Host todo list"
  type="secondary"
  textLead="This copyable todo list summarises the steps outlined throughout the article in more detail ."
>
  <ol>
    {#each items as item, i}
      <li>{item}</li>
    {/each}
  </ol>
  <CopyButton text={copyableItems} on:click={share} block={false}
    >Copy</CopyButton
  >
</Hero>

<style>
  li {
    font-family: "Port Lligat Slab", serif;
  }
</style>
