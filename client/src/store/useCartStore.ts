import { create } from 'zustand';
import { api } from '../api/client';
import { CartItem } from '../types';
import { useAuthStore } from './useAuthStore';

interface CartState {
  items: CartItem[];
  itemCount: number;
  addItem: (productId: string) => Promise<void>;
  removeItem: (itemId: number) => Promise<void>;
  updateQuantity: (itemId: number, productId: number, quantity: number) => Promise<void>;
  clearCart: () => Promise<void>;
  fetchCart: () => Promise<void>;
}

export const useCartStore = create<CartState>()((set) => {
  // Initialize cart on store creation
  const initializeCart = async () => {
    try {
      if (!useAuthStore.getState().isAuthenticated) {
        set({ items: [], itemCount: 0 });
        return;
      }
      const response = await api.get('/cart');
      const items = response.data || [];
      const itemCount = items.reduce((acc: number, item: any) => acc + (item.quantity || 0), 0);
      set({ items, itemCount });
    } catch (error) {
      console.error('Failed to initialize cart:', error);
      set({ items: [], itemCount: 0 });
    }
  };

  // Subscribe to auth changes
  useAuthStore.subscribe((state: any) => {
    if (state.isAuthenticated) {
      initializeCart();
    } else {
      set({ items: [], itemCount: 0 });
    }
  });

  // Initial cart fetch
  initializeCart();

  return {
    items: [],
    itemCount: 0,
    fetchCart: async () => {
      try {
        if (!useAuthStore.getState().isAuthenticated) {
          set({ items: [], itemCount: 0 });
          return;
        }
        const response = await api.get('/cart');
        const items = response.data || [];
        const itemCount = items.reduce((acc: number, item: any) => acc + (item.quantity || 0), 0);
        set({ items, itemCount });
      } catch (error) {
        console.error('Failed to fetch cart:', error);
        set({ items: [], itemCount: 0 });
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
        const items = response.data || [];
        const itemCount = items.reduce((acc: number, item: any) => acc + (item.quantity || 0), 0);
        set({ items, itemCount });
      } catch (error) {
        console.error('Failed to add item:', error);
        throw error;
      }
    },
    removeItem: async (itemId: number) => {
      try {
        await api.delete(`/cart/delete/${itemId}`);
        const response = await api.get('/cart');
        const items = response.data || [];
        const itemCount = items.reduce((acc: number, item: any) => acc + (item.quantity || 0), 0);
        set({ items, itemCount });
      } catch (error) {
        console.error('Failed to remove item:', error);
        throw error;
      }
    },
    updateQuantity: async (itemId: number, productId: number, quantity: number) => {
      try {
        await api.put(`/cart/update/${itemId}`, {
          product_id: productId,
          quantity
        });
        const response = await api.get('/cart');
        const items = response.data || [];
        const itemCount = items.reduce((acc: number, item: any) => acc + (item.quantity || 0), 0);
        set({ items, itemCount });
      } catch (error) {
        console.error('Failed to update quantity:', error);
        throw error;
      }
    },
    clearCart: async () => {
      try {
        const response = await api.get('/cart');
        if (response.data) {
          for (const item of response.data) {
            await api.delete(`/cart/delete/${item.id}`);
          }
        }
        set({ items: [], itemCount: 0 });
      } catch (error) {
        console.error('Failed to clear cart:', error);
        throw error;
      }
    }
  };
});