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
        target: 'http://localhost',
        changeOrigin: true,
        secure: false,
      },
      '/product': {
        target: 'http://localhost',
        changeOrigin: true,
        secure: false,
      },
      '/cart': {
        target: 'http://localhost',
        changeOrigin: true,
        secure: false,
      },
    },
  },
});