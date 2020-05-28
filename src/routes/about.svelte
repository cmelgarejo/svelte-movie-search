<script>
  import { _ } from "svelte-i18n";
  const xp = new Date().getFullYear() - 2005;
  const links = [
    "https://cmelgarejo.net",
    "https://www.linkedin.com/in/cmelgarejo",
    "https://github.com/cmelgarejo",
    "https://profile.codersrank.io/user/cmelgarejo/"
  ];
</script>

<style>
  a {
    color: teal;
  }
</style>

<svelte:head>
  <title>{$_('route.about.title')}</title>
</svelte:head>

<h1>{$_('route.about.title')}</h1>
<p>{$_('route.about.dev', { values: { xp } })}</p>
<ul>
  {#each links as link}
    <a href={link}>{link}</a>
  {/each}
</ul>
