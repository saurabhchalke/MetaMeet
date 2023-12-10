<script context="module">
  import {Form, Input, Progress} from "spaper"
  import Datepicker from "$lib/Datepicker.svelte"
  import {goto} from "$app/navigation"
  import Next from "$lib/Next.svelte"
  import {page} from "$app/stores"
</script>

<script>
  export let user
  export let social

  let organizer = social.invitees[social.organizer]
  let name = $page.url.searchParams.get("name") || user.name
  let loading
  let dates = user.dates.map(date => new Date(date).getTime()) || []
</script>

<svelte:head>
  <title>You</title>
</svelte:head>

<!-- TODO: Should I use stores or session here instead of the query param? -->
<!-- Should use load function to support rendering existing selection here. -->
<!-- Need to add ability to edit name here too because pre-rendered page won't show data if for some reason the user comes back here. -->
<h1>Hello {name}, when are you free?</h1>
<!-- svelte-ignore missing-declaration -->
<Form
  id="you"
  on:submit={async _ => {
    loading = true
    try {
      await fetch("you", {
        method: "PUT",
        body: JSON.stringify({
          dates,
        }),
      })
      goto("everyone")
    } finally {
      loading = false
    }
  }}
>
  <Datepicker bind:selected={dates} marked={organizer.dates} />
  <p>👑 = days the organizer has chosen</p>
  <div style="text-align: center;">
    <Next />
  </div>
  <div style="margin: 1rem;">
    <Progress style={`visibility: ${loading ? "visible" : "hidden"};`} infinite striped />
  </div>
</Form>
