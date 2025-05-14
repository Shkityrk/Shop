import { create } from 'zustand';
import { api } from '../api/client';
import { CartItem } from '../types';
import { useAuthStore } from './useAuthStore';

interface CartState {
  items: CartItem[];
  isLoading: boolean;
  addItem: (productId: string) => Promise<void>;
  removeItem: (itemId: number) => Promise<void>;
  updateQuantity: (itemId: number, productId: number, quantity: number) => Promise<void>;
  clearCart: () => Promise<void>;
  fetchCart: () => Promise<void>;
}

export const useCartStore = create<CartState>((set) => ({
  items: [],
  isLoading: false,
  fetchCart: async () => {
    try {
      if (!useAuthStore.getState().isAuthenticated) {
        set({ items: [] });
        return;
      }

      set({ isLoading: true });
      const cartResponse = await api.get('/cart');
      const cartItems = cartResponse.data || [];

      // Fetch product details for each cart item
      const productsResponse = await api.get('/product/list');
      const products = productsResponse.data || [];

      // Combine cart items with product details
      const cartWithProducts = cartItems.map((cartItem: CartItem) => ({
        ...cartItem,
        product: products.find((p: any) => p.id === cartItem.product_id)
      }));

      set({ items: cartWithProducts, isLoading: false });
    } catch (error) {
      console.error('Failed to fetch cart:', error);
      set({ items: [], isLoading: false });
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
      const cartResponse = await api.get('/cart');
      const cartItems = cartResponse.data || [];
      const productsResponse = await api.get('/product/list');
      const products = productsResponse.data || [];
      const cartWithProducts = cartItems.map((cartItem: CartItem) => ({
        ...cartItem,
        product: products.find((p: any) => p.id === cartItem.product_id)
      }));
      set({ items: cartWithProducts });
    } catch (error) {
      console.error('Failed to add item:', error);
      throw error;
    }
  },
  removeItem: async (itemId: number) => {
    try {
      await api.delete(`/cart/delete/${itemId}`);
      const response = await api.get('/cart');
      set({ items: response.data || [] });
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
      const cartResponse = await api.get('/cart');
      const cartItems = cartResponse.data || [];
      const productsResponse = await api.get('/product/list');
      const products = productsResponse.data || [];
      const cartWithProducts = cartItems.map((cartItem: CartItem) => ({
        ...cartItem,
        product: products.find((p: any) => p.id === cartItem.product_id)
      }));
      set({ items: cartWithProducts });
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
      set({ items: [] });
    } catch (error) {
      console.error('Failed to clear cart:', error);
      throw error;
    }
  },
}));