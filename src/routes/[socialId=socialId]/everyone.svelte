<script>
  import { Form, Progress, Switch } from "spaper";
  import { goto } from "$app/navigation";
  import { page, session } from "$app/stores";
  // import Datepicker from '$lib/Datepicker.svelte'
  import convertDatesToStrings from "$lib/convertDatesToStrings";
  import BestDates from "$lib/BestDates.svelte";
  import { toDatabaseId } from "$lib/id";
  import Next from "$lib/Next.svelte";
  import Retreat from "$lib/Back.svelte";
  import Invitees from "$lib/Invitees.svelte";
  import faunadb from "faunadb";
  import { getContext, onMount } from "svelte";
  import StreamingStatus from "$lib/StreamingStatus.svelte";
  import Inviter from "$lib/Inviter.svelte";
  import push from "$lib/push";

  const { getAnalytics } = getContext("analytics");

  export let social;
  export let user;

  let invitees;
  let inviteesCount;
  let decision = social.decision || null;
  let status = "Not started";
  let loading;

  $: invitees = social.invitees;
  $: inviteesCount = Object.values(invitees).length;
  $: inviteesWithDates = Object.values(invitees).filter(
    (invitee) => invitee.dates && invitee.dates.length
  );

  let pushBrowserSupport;
  let pushPermission;
  let pushSubscription;
  let pushSwitch = true;
  $: pushSwitch = pushBrowserSupport && pushPermission !== "denied";

  onMount(async () => {
    const q = faunadb.query;
    const client = new faunadb.Client({ ...$session.faunadb });

    const databaseId = toDatabaseId($page.params.socialId);
    const docRef = q.Ref(q.Collection("social"), databaseId);

    let stream;

    console.info("Streaming social", databaseId);
    const startStream = () => {
      stream = client.stream
        .document(docRef)
        .on("snapshot", (snapshot) => {
          social = convertDatesToStrings(snapshot.data);
        })
        .on("version", (version) => {
          social = convertDatesToStrings(version.document.data);

          const newDecision =
            social.decision &&
            social.decision.value !== sessionStorage.getItem("decisionSeen");
          if (newDecision) goto("decision");

          status = "📡 Updated: someone joined the party!";
        })
        .on("error", (error) => {
          console.error(error);
          stream.close();
          setTimeout(startStream, 1000);
        })
        .start();
    };

    startStream();
    status = "📡 Live-streaming updates";

    pushBrowserSupport =
      "serviceWorker" in navigator &&
      "PushManager" in window &&
      "Notification" in window;
    if (!pushBrowserSupport) return;

    pushPermission = Notification.permission;
    pushSubscription = await push.getExistingSubscription();

    if (
      pushSubscription &&
      (!user.pushSubscriptions ||
        !Object.values(user.pushSubscriptions).find(
          (savedPushSubscription) =>
            pushSubscription.endpoint === savedPushSubscription
        ))
    ) {
      await fetch("push.json", {
        method: "POST",
        body: JSON.stringify({
          push: pushSubscription,
        }),
      });
    }
  });
</script>

<svelte:head>
  <title>Everyone</title>
</svelte:head>

<StreamingStatus {status} />

<h1>
  {inviteesCount > 1 ? "Here are the others!" : "You are the only one here..."}
</h1>
<Invitees {invitees} />
<Inviter />

<p style="margin-top: 1rem;">
  Enable push notifications and we'll update you when people join the social or
  a decision is made.
</p>
<Form
  title={!pushBrowserSupport
    ? "Your browser does not support push notifications"
    : pushPermission === "denied"
    ? "You've denied notification permissions. Update your browser settings for this website and click this button again if you'd like to enable them."
    : !pushSubscription
    ? "Push permissions are granted but you need to enable them."
    : "Push notifications are enabled 📡"}
>
  <Switch
    name="push-notifications"
    inline
    checked={pushSwitch}
    on:change={async (event) => {
      const switchOn = event.detail;

      if (!switchOn) {
        getAnalytics().trackEvent("Disable push notifications");
        await push.unsubscribe();
        await fetch("push.json", {
          method: "DELETE",
          body: JSON.stringify({
            push: pushSubscription,
          }),
        });
        pushSubscription = null;
        return;
      }

      getAnalytics().trackEvent("Enable push notifications");
      pushPermission = await push.askPermission();
      const permissionGranted = pushPermission === "granted";
      if (permissionGranted) {
        pushSubscription = await push.subscribe();
        await fetch("push.json", {
          method: "POST",
          body: JSON.stringify({
            push: pushSubscription,
          }),
        });
      }
    }}
    style="background-color: cornflowerblue;"
    >{pushPermission === "granted" && pushSubscription
      ? "DISABLE PUSH NOTIFICATIONS"
      : "ENABLE PUSH NOTIFICATIONS"}
  </Switch>
</Form>

<!-- <h2>Here's everyone's availability</h2>
<Datepicker disabledTo={10000} selected={dates}/> -->

{#if inviteesWithDates.length > 1}
  <h1>Choose a date</h1>
  <!-- svelte-ignore missing-declaration -->
  <Form
    id="everyone"
    on:submit={async (e) => {
      loading = true;

      try {
        const formData = new FormData(e.target);
        const decision = formData.get("best-dates");

        await fetch("everyone", {
          method: "PATCH",
          body: JSON.stringify({
            decision,
          }),
        });
        goto("decision");
      } finally {
        loading = false;
      }
    }}
  >
    <BestDates {invitees} bind:selected={decision} />
  </Form>
{:else}
  <h1>Next steps</h1>
  <ol>
    <li>Wait for others to join 👆</li>
    <li>See which dates everyone is available 📅</li>
    <li>Choose a date and you're done! ✅</li>
  </ol>
{/if}

<div style="margin: 1rem;">
  <Retreat back="you" />
  <Next disabled={inviteesWithDates.length < 2 || !decision} form="everyone" />
  <div style="margin: 1rem;">
    <Progress
      style={`visibility: ${loading ? "visible" : "hidden"};`}
      infinite
      striped
    />
  </div>
</div>

<style>
  h1 {
    align-self: center;
  }
</style>
