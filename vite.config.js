import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    hmr: true, // Enable Hot Module Replacement
  },
  root: 'assets',
  build: {
    manifest: 'manifest.json',
    rollupOptions: {
      input: {
        main: '/assets/js/main.js', // TODO: to change to main.ts
      },
    },
    // server: {
    //   origin: 'http://127.0.0.1:8000',
    // },
  }
})
