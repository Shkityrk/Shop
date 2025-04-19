import { create } from 'zustand';
import { api } from '../api/client';
import { CartItem } from '../types';
import { useAuthStore } from './useAuthStore';

interface CartState {
  items: CartItem[];
  addItem: (productId: string) => Promise<void>;
  removeItem: (productId: string) => Promise<void>;
  updateQuantity: (productId: string, quantity: number) => Promise<void>;
  clearCart: () => Promise<void>;
  fetchCart: () => Promise<void>;
}

export const useCartStore = create<CartState>((set) => ({
  items: [],
  fetchCart: async () => {
    try {
      const response = await api.get('/cart');
      set({ items: response.data.items || [] });
    } catch (error) {
      console.error('Failed to fetch cart:', error);
      set({ items: [] });
    }
  },
  addItem: async (productId: string) => {
    try {
      if (!useAuthStore.getState().isAuthenticated) {
        throw new Error('Please login to add items to cart');
      }
      await api.post('/cart/add', {
        product_id: parseInt(productId),
        quantity: 1
      });
      const response = await api.get('/cart');
      set({ items: response.data.items || [] });
    } catch (error) {
      console.error('Failed to add item:', error);
      throw error;
    }
  },
  removeItem: async (productId: string) => {
    try {
      await api.delete(`/cart/delete/${productId}`);
      const response = await api.get('/cart');
      set({ items: response.data.items || [] });
    } catch (error) {
      console.error('Failed to remove item:', error);
      throw error;
    }
  },
  updateQuantity: async (productId: string, quantity: number) => {
    try {
      await api.put(`/cart/update/${productId}`, {
        product_id: parseInt(productId),
        quantity
      });
      const response = await api.get('/cart');
      set({ items: response.data.items || [] });
    } catch (error) {
      console.error('Failed to update quantity:', error);
      throw error;
    }
  },
  clearCart: async () => {
    try {
      const response = await api.get('/cart');
      if (response.data.items) {
        for (const item of response.data.items) {
          await api.delete(`/cart/delete/${item.product_id}`);
        }
      }
      set({ items: [] });
    } catch (error) {
      console.error('Failed to clear cart:', error);
      throw error;
    }
  },
}));