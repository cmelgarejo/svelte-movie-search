<script>
  import { _, locale, locales } from "svelte-i18n";
  import "./Nav.css";
  export let segment;
</script>

<nav>
  <ul>
    <li>
      <a class:selected={segment === undefined} href=".">
        {$_('components.Nav.home')}
      </a>
    </li>
    <li>
      <a class:selected={segment === 'about'} href="about">
        {$_('components.Nav.about')}
      </a>
    </li>

    <!-- for the blog link, we're using rel=prefetch so that Sapper prefetches
		     the blog data when we hover over the link or tap it on a touchscreen -->
    <!-- <li>
      <a rel="prefetch" class:selected={segment === 'blog'} href="blog">
        {$_('components.Nav.blog')}
      </a>
    </li> -->
  </ul>
  <ul>
    {#each $locales as item}
      <li>
        <span
          class="a"
          class:selected={$locale.includes(item)}
          href={`#!${item}`}
          on:click={() => ($locale = item)}>
          {item.replace('-', '_')}
        </span>
      </li>
    {/each}
  </ul>
</nav>
