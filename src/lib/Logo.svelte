<script context="module">
  import logo from "$lib/logo - 240.png";
</script>

<script>
  export let width;
</script>

<img
  class="logo"
  src={logo}
  alt="Schedule Your Social logo"
  {width}
  height={width}
/>

<style>
  img.logo {
    display: inline-block;
    margin: 0 1rem 0 0;
    border: none;
    border-radius: 8px;
    vertical-align: middle;
  }
</style>
