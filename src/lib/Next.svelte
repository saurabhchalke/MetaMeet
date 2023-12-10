<script context="module">
  import {Button} from "spaper"
</script>

<script>
  export let disabled
  export let form
</script>

<span><Button {disabled} {form} type="success" nativeType="submit">NEXT</Button></span>

<style>
  span {
    margin-left: 0.5rem;
  }
</style>
