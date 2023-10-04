import { defineConfig } from 'vite'


export default defineConfig({
  build: {
    // generate manifest.json in outDir
    manifest: true,

    rollupOptions: {
      input: 'src/main.ts',
      output: {
        file: 'main.compiled.js',
        dir: 'dist'
      }
    }
  }
})
