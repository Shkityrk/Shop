import { create } from 'zustand';
import { api } from '../api/client';

interface User {
  username: string;
  email: string;
  first_name: string;
  last_name: string;
}

interface AuthState {
  user: User | null;
  isAuthenticated: boolean;
  login: (username: string, password: string) => Promise<void>;
  register: (data: {
    username: string;
    email: string;
    password: string;
    first_name: string;
    last_name: string;
  }) => Promise<void>;
  logout: () => Promise<void>;
}

export const useAuthStore = create<AuthState>((set) => ({
  user: null,
  isAuthenticated: Boolean(document.cookie.includes('access_token=')),
  login: async (username: string, password: string) => {
    try {
      const response = await api.post('/auth/login', { username, password });
      set({ user: response.data.user, isAuthenticated: true });
    } catch (error) {
      console.error('Login failed:', error);
      throw error;
    }
  },
  register: async (data) => {
    try {
      const response = await api.post('/auth/register', data);
      set({ user: response.data.user, isAuthenticated: true });
    } catch (error) {
      console.error('Registration failed:', error);
      throw error;
    }
  },
  logout: async () => {
    try {
      await api.post('/auth/logout');
      document.cookie = 'access_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
      set({ user: null, isAuthenticated: false });
    } catch (error) {
      console.error('Logout failed:', error);
      throw error;
    }
  },
}));