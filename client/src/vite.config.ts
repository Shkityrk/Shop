import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  optimizeDeps: {
    exclude: ['lucide-react'],
  },
  server: {
    proxy: {
      '/auth': {
        target: 'http://79.137.197.216',
        changeOrigin: false,
        secure: false,
      },
      '/product': {
        target: 'http://79.137.197.216',
        changeOrigin: false,
        secure: false,
      },
      '/cart': {
        target: 'http://79.137.197.216',
        changeOrigin: false,
        secure: false,
      },
    },
  },
});