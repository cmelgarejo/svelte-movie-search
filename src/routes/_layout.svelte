<script context="module">
  import { _, waitLocale, isLoading } from "svelte-i18n";
  export async function preload() {
    return waitLocale();
  }
</script>

<script>
  import Nav from "components/Nav/Nav.svelte";
  export let segment;
</script>

<style>
  main {
    position: relative;
    max-width: 56em;
    background-color: white;
    padding: 2em;
    margin: 0 auto;
    box-sizing: border-box;
  }
  .loading {
    text-align: center;
  }
</style>

{#if $isLoading}
  <div class="loading">{$_('route.layout.loading.message')}</div>
{:else}
  <Nav {segment} />

  <main>
    <slot />
  </main>
{/if}
