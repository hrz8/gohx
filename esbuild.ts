import fs from 'node:fs';

import * as esbuild from 'esbuild';
import {copy} from 'esbuild-plugin-copy';

const config: esbuild.BuildOptions = {
  entryPoints: [
    {
      in: 'src/js/main.js',
      out: 'js/bundled.min'
    },
  ],
  bundle: true,
  minify: true,
  logLevel: 'debug',
  metafile: true,
  sourcemap: true,
  legalComments: 'linked',
  allowOverwrite: true,
  outbase: 'src',
  outdir: './assets/public/',
  target: ['es2020', 'chrome58', 'edge16', 'firefox57', 'safari11'],
  loader: {
    '.png': 'dataurl',
    '.woff': 'dataurl',
    '.woff2': 'dataurl',
    '.eot': 'dataurl',
    '.ttf': 'dataurl',
    '.svg': 'dataurl',
  },
  plugins: [
    copy({
      resolveFrom: 'cwd',
      assets: {
        from: ['./node_modules/htmx.org/dist/htmx.min.js'],
        to: ['./assets/public/vendor'],
      },
      watch: true,
    }),
  ],
};

async function build() {
  try {  
    const result = await esbuild.build(config);
    fs.writeFileSync('meta.json', JSON.stringify(result.metafile));
  } catch (error) {
    console.error('esbuild error', error);
    process.exit(1);
  }
}

build();
