import { defineConfig } from 'vite'


export default defineConfig({
  build: {
    server: {
      proxy: {
        '/static': 'http://127.0.0.1:8000'
      }
    },
    // generate manifest.json in outDir
    manifest: true,

    rollupOptions: {
      input: 'assets/js/main.ts',
      output: {
        file: 'main.compiled.js',
        dir: 'dist'
      }
    }
  }
})
