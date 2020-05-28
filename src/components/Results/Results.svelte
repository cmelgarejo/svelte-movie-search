<script>
  import "./Results.css";
  export let movies = [];
  export let searching = false;
</script>

{#if searching}
  <div class="loader">
    <div class="lds-ellipsis">
      <div />
      <div />
      <div />
      <div />
    </div>
  </div>
{/if}
<div class="container">
  <div class="row">
    {#each movies as m}
      <div class="column">
        <div class="card">
          <img
            class="card-image"
            src={`https://image.tmdb.org/t/p/w200/${m.poster_path}`}
            alt={m.title} />
          <div class="card-info">
            <div class="card-title">{m.title}</div>
            <div class="card-content">{m.overview}</div>
          </div>
          <!-- <div class="card-action">
            <div>{m.release_date}</div>
            <div>{m.vote_average}</div>
            <div>{m.vote_count}</div>
            <div>{m.overview}</div>
            <a href="#">{m.link}</a>
          </div> -->
        </div>
      </div>
    {/each}
  </div>
</div>
