<script>
  import {goto} from "$app/navigation"
  import {page, session} from "$app/stores"
  // import Datepicker from '$lib/Datepicker.svelte'
  import convertDatesToStrings from "$lib/convertDatesToStrings"
  import BestDates from "$lib/BestDates.svelte"
  import {toDatabaseId} from "$lib/id"
  import Next from "$lib/Next.svelte"
  import Retreat from "$lib/Back.svelte"
  import Invitees from "$lib/Invitees.svelte"
  import faunadb from "faunadb"
  import {onMount} from "svelte"
  import StreamingStatus from "$lib/StreamingStatus.svelte"
  import Inviter from "$lib/Inviter.svelte"

  export let social

  let invitees
  let inviteesCount
  // let dates
  let decision = social.decision ? social.decision : null
  let status = "Not started"

  $: invitees = social.invitees
  $: inviteesCount = Object.values(invitees).length
  $: inviteesWithDates = Object.values(invitees).filter(
    invitee => invitee.dates && invitee.dates.length
  )
  // $: dates = Object.entries(invitees)
  // 	.filter(([inviteeId, _]) => selected ? inviteeId === selected : true)
  // 	.map(([_, invitee]) => invitee.dates)
  // 	.reduce((allDates, inviteeDates) => allDates.concat(inviteeDates), [])
  // 	.map(date => new Date(date).getTime())

  onMount(() => {
    const q = faunadb.query
    const client = new faunadb.Client({...$session.faunadb})

    const databaseId = toDatabaseId($page.params.socialId)
    const docRef = q.Ref(q.Collection("social"), databaseId)

    let stream

    console.info("Streaming social", databaseId)
    const startStream = () => {
      stream = client.stream
        .document(docRef)
        .on("snapshot", snapshot => {
          console.debug("snapshot")
          social = convertDatesToStrings(snapshot.data)
        })
        .on("version", version => {
          console.debug("version")
          social = convertDatesToStrings(version.document.data)

          const decisionUpdated = social.decision && social.decision !== decision
          if (decisionUpdated) goto("decision")

          status = "Updated: someone joined the party!"
        })
        .on("error", error => {
          console.debug("Error:", error)
          stream.close()
          setTimeout(startStream, 1000)
        })
        .start()
    }

    startStream()
    status = "Live-streaming updates 📡"
  })
</script>

<svelte:head>
  <title>Everyone</title>
</svelte:head>

<StreamingStatus {status} />

<h2>
  {inviteesCount > 1 ? "Here are the other invitees!" : "You are the only invitee here..."}
</h2>
<Invitees {invitees} />
<Inviter />

<!-- <h2>Here's everyone's availability</h2>
<Datepicker disabledTo={10000} selected={dates}/> -->

{#if inviteesWithDates.length > 1}
  <h2>Choose a date</h2>
  <!-- svelte-ignore missing-declaration -->
  <form
    id="everyone"
    on:submit|preventDefault={async e => {
      const formData = new FormData(e.target)
      const decision = formData.get("best-dates")

      await fetch("everyone", {
        method: "PATCH",
        body: JSON.stringify({
          decision,
        }),
      })
      goto("decision")
    }}
  >
    <BestDates {invitees} bind:selected={decision} />
  </form>
{:else}
  <h2>Next steps</h2>
  <ol>
    <li>Invite more invitees 👆</li>
    <li>Choose a date 📅</li>
    <li>You're done! ✅</li>
  </ol>
{/if}

<div>
  <Retreat back="you" />
  <Next disabled={inviteesWithDates.length < 2 || !decision} form="everyone" />
</div>

<style>
  h2 {
    align-self: flex-start;
  }
</style>
