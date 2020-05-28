import resolve from "@rollup/plugin-node-resolve";
import replace from "@rollup/plugin-replace";
import commonjs from "@rollup/plugin-commonjs";
import svelte from "rollup-plugin-svelte";
import babel from "@rollup/plugin-babel";
import { terser } from "rollup-plugin-terser";
import config from "sapper/config/rollup.js";
import pkg from "./package.json";
import json from "@rollup/plugin-json";
import alias from "@rollup/plugin-alias";

const mode = process.env.NODE_ENV;
const dev = mode === "development";
const legacy = !!process.env.SAPPER_LEGACY_BUILD;
const aliases = alias({
  resolve: [".svelte", ".js"],
  entries: [
    { find: "components", replacement: "src/components" },
    { find: "locales", replacement: "src/locales" },
    { find: "util", replacement: "src/util" },
  ],
});

// modified version of onwarn provided by sapper projects
const onwarn = (warning, onwarn) => {
  if (
    (warning.code === "CIRCULAR_DEPENDENCY" &&
      /[/\\]@sapper[/\\]/.test(warning.message)) ||
    warning.code === "THIS_IS_UNDEFINED"
  )
    return;
  onwarn(warning);
};

const dedupe = (importee) =>
  importee === "svelte" || importee.startsWith("svelte/");

export default {
  client: {
    input: config.client.input(),
    output: config.client.output(),
    plugins: [
      replace({
        "process.browser": true,
        "process.env.NODE_ENV": JSON.stringify(mode),
      }),
      svelte({
        dev,
        hydratable: true,
        emitCss: true,
      }),
      resolve({
        browser: true,
        dedupe,
      }),
      commonjs(),
      json({
        compact: true,
      }),
      legacy &&
        babel({
          extensions: [".js", ".mjs", ".html", ".svelte"],
          babelHelpers: "runtime",
          exclude: ["node_modules/@babel/**"],
          presets: [
            [
              "@babel/preset-env",
              {
                targets: "> 0.25%, not dead",
              },
            ],
          ],
          plugins: [
            "@babel/plugin-syntax-dynamic-import",
            [
              "@babel/plugin-transform-runtime",
              {
                useESModules: true,
              },
            ],
          ],
        }),
      !dev &&
        terser({
          module: true,
        }),
      aliases,
    ],
    preserveEntrySignatures: false,
    onwarn,
  },
  server: {
    input: config.server.input(),
    output: config.server.output(),
    plugins: [
      replace({
        "process.browser": false,
        "process.env.NODE_ENV": JSON.stringify(mode),
      }),
      svelte({
        generate: "ssr",
        dev,
      }),
      resolve({
        dedupe,
      }),
      commonjs(),
      json({
        compact: true,
      }),
      aliases,
    ],
    external: Object.keys(pkg.dependencies).concat(
      require("module").builtinModules ||
        Object.keys(process.binding("natives"))
    ),
    preserveEntrySignatures: false,
    onwarn,
  },

  serviceworker: {
    input: config.serviceworker.input(),
    output: config.serviceworker.output(),
    plugins: [
      resolve(),
      replace({
        "process.browser": true,
        "process.env.NODE_ENV": JSON.stringify(mode),
      }),
      commonjs(),
      json({
        compact: true,
      }),
      !dev && terser(),
      aliases,
    ],
    preserveEntrySignatures: false,
    onwarn,
  },
};
