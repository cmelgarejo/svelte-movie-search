<script>
  import { _ } from "svelte-i18n";
  import Search from "components/Search/Search.svelte";
  import Results from "components/Results/Results.svelte";

  let error;
  let loader = { loading: false, message: $_("route.layout.loading.message") };
  let movies = [];

  const onError = e => (error = e.detail);
  const onFound = e => (movies = e.detail);
  const onSearching = e => {
    loader = e.detail;
    if (loader.loading) movies = [];
  };
</script>

<style>
  :global(html) {
    font-size: 10px;
  }

  :global(*) {
    box-sizing: border-box;
  }

  :global(body) {
    margin: 0;
    padding: 0;
    color: #333;
  }

  :global(p) {
    font-size: 1.6rem;
  }

  :global(small) {
    font-size: 1.2rem;
  }

  :global(.error) {
    text-align: center;
    color: #ff0000;
  }
</style>

<svelte:head>
  <title>{$_('route.index.title')}</title>
</svelte:head>

<Search on:error={onError} on:found={onFound} on:searching={onSearching} />
{#if error}
  <p class="error">{error}</p>
{/if}
<!-- {#if loader.loading}
  <p class="error">{loader.message}</p>
{/if} -->
<Results on:error={onError} {movies} searching={loader.loading} />
