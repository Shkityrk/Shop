import { create } from 'zustand';
import { persist } from 'zustand/middleware';
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
  clearAuth: () => void;
}


export const useAuthStore = create<AuthState>()(
    persist(
        (set) => ({
            user: null,
            isAuthenticated: Boolean(document.cookie.includes('access_token=')),
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
            // ← вот здесь, внутри функции:
            clearAuth: () => set({ user: null, isAuthenticated: false }),
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
