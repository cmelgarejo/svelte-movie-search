<script>
  import { _, locale } from "svelte-i18n";
  import "./Search.css";
  import { createEventDispatcher } from "svelte";
  const dispatcher = createEventDispatcher();
  async function search(e) {
    let loadingObj = {
      loading: true,
      message: $_("components.Search.loading.message")
    };
    dispatcher("searching", loadingObj);
    try {
      let query = e.target.query.value;
      if (query === "") {
        dispatcher("error", $_("components.Search.query.errors.empty"));
      } else {
        const res = await fetch(
          `https://api.themoviedb.org/3/search/movie?api_key=32f5a11a2577818213010e38bacc5a55&language=${$locale}&query=${query}&page=1&include_adult=true`
        );
        const json = await res.json();
        if (json.results && json.results.length)
          dispatcher("found", json.results);
        else
          dispatcher(
            "error",
            $_("components.Search.query.errors.notFound", { values: { query } })
          );
      }
    } catch (error) {
      dispatcher(
        "error",
        $_("components.Search.query.errors.generic", {
          values: { query, error }
        })
      );
    }
    loadingObj.loading = false;
    dispatcher("searching", loadingObj);
  }
</script>

<div class="container">
  <h1 class="title">{$_('components.Search.title')}</h1>
  <form class="form" on:submit|preventDefault={search}>
    <label class="label" htmlFor="query">
      {$_('components.Search.query.label')}
    </label>
    <input
      class="input"
      type="text"
      name="query"
      placeholder={$_('components.Search.query.placeholder')} />
    <button class="button" type="submit">
      {$_('components.Search.query.button')}
    </button>
  </form>
</div>
