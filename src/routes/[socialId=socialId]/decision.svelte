<script context="module">
  import Retreat from "$lib/Back.svelte"
</script>

<script>
  import {onMount} from "svelte"

  export let social
  let friendlyDecision
  $: friendlyDecision = new Date(social.decision).toLocaleDateString(undefined, {
    weekday: "long",
    month: "long",
    day: "numeric",
  })

  onMount(() => {
    sessionStorage.setItem("decisionSeen", social.decision)
  })
</script>

<h1>Your social is on {friendlyDecision}</h1>
<p>
  Have fun 🎉 and if you found this app helpful then please share it with others or <a
    href="https://www.producthunt.com/products/schedule-your-social/reviews/new"
    >leave us a review</a
  > 🙏
</p>

<Retreat back="everyone" />
