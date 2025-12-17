import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import { api } from '../api/client';
import type { User } from '../types';

interface AuthState {
  user: User | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (username: string, password: string) => Promise<void>;
  register: (data: {
    username: string;
    email: string;
    password: string;
    first_name: string;
    last_name: string;
  }) => Promise<void>;
  logout: () => Promise<void>;
  clearAuth: () => void;
  fetchUser: () => Promise<void>;
}


export const useAuthStore = create<AuthState>()(
    persist(
        (set) => ({
            user: null,
            isAuthenticated: Boolean(document.cookie.includes('access_token=')),
            isLoading: false,
            login: async (username, password) => {
                const resp = await api.post('/auth/login', { username, password });
                set({ user: resp.data.user, isAuthenticated: true });
            },
            register: async (data) => {
                const resp = await api.post('/auth/register', data);
                set({ user: resp.data.user, isAuthenticated: true });
            },
            logout: async () => {
                await api.post('/auth/logout');
                document.cookie = 'access_token=; expires=0; path=/;';
                set({ user: null, isAuthenticated: false });
            },
            clearAuth: () => set({ user: null, isAuthenticated: false }),
            fetchUser: async () => {
                set({ isLoading: true });
                try {
                    const resp = await api.get('/auth/info');
                    set({ user: resp.data, isAuthenticated: true, isLoading: false });
                } catch {
                    set({ user: null, isAuthenticated: false, isLoading: false });
                }
            },
        }),
        {
            name: 'auth-storage',
            partialize: (state) => ({
                user: state.user,
                isAuthenticated: state.isAuthenticated,
            }),
        }
    )
);
