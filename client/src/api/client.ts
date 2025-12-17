import axios from 'axios';
import { useAuthStore } from '../store/useAuthStore'

const API_URL = import.meta.env.VITE_API_URL || 'https://localhost';

export const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
  withCredentials: true,
});

// Add request interceptor to handle authentication
api.interceptors.request.use((config: any) => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('access_token='))
    ?.split('=')[1];

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  // Ensure headers are properly set for all requests
  if (!config.headers) config.headers = {};
  // Use type assertion to avoid AxiosHeaders/Record mismatch
  (config.headers as Record<string, string>)['Content-Type'] = 'application/json';
  (config.headers as Record<string, string>)['Accept'] = 'application/json';
  return config;
});

// Add response interceptor to handle errors
api.interceptors.response.use(
    response => response,
    error => {
      const pathname = window.location.pathname;
      // Guard against cases where the network failed and error.response is undefined
      if (error.response?.status === 401) {
        // 1) Сброс состояния
        useAuthStore.getState().clearAuth();
        // 2) Если мы уже не на /login — редиректим
        if (pathname !== '/login') {
          // replace чтобы не захламлять историю
          window.location.replace('/login');
        }

      }
      return Promise.reject(error)
    }
);

export const getAuthToken = () => {
  const cookies = document.cookie.split(';');
  const tokenCookie = cookies.find(cookie => cookie.trim().startsWith('access_token='));
  return tokenCookie ? tokenCookie.split('=')[1] : null;
};