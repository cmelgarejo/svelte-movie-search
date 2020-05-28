import {
  register,
  init,
  getLocaleFromNavigator,
  waitLocale,
} from "svelte-i18n";

register("en", () => import("locales/en.json"));
register("es", () => import("locales/es.json"));
init({
  fallbackLocale: "en",
  initialLocale: getLocaleFromNavigator(),
});
