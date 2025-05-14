import { create } from 'zustand'
import { api } from '../api/client'
import { CartItem } from '../types'
import { useAuthStore } from './useAuthStore'

interface CartState {
  items: CartItem[]
  isLoading: boolean
  addItem: (productId: string) => Promise<void>
  removeItem: (itemId: number) => Promise<void>
  updateQuantity: (itemId: number, productId: number, quantity: number) => Promise<void>
  clearCart: () => Promise<void>
  fetchCart: () => Promise<void>
}

export const useCartStore = create<CartState>((set) => ({
  items: [],
  isLoading: false,

  fetchCart: async () => {
    if (!useAuthStore.getState().isAuthenticated) {
      set({ items: [], isLoading: false })
      return
    }

    set({ isLoading: true })
    try {
      const cartResponse = await api.get('/cart')
      const cartItems: CartItem[] = cartResponse.data || []

      const productsResponse = await api.get('/product/list')
      const products = productsResponse.data || []

      const cartWithProducts = cartItems.map((ci) => ({
        ...ci,
        product: products.find((p: any) => p.id === ci.product_id) || null
      }))

      set({ items: cartWithProducts, isLoading: false })
    } catch (err: any) {
      // при 401 просто выходим без повторных попыток
      if (err.response?.status === 401) {
        set({ items: [], isLoading: false })
        return
      }
      console.error('Failed to fetch cart:', err)
      set({ items: [], isLoading: false })
    }
  },

  addItem: async (productId: string) => {
    if (!useAuthStore.getState().isAuthenticated) {
      throw new Error('Please login to add items to cart')
    }
    try {
      await api.post('/cart/add', {
        product_id: parseInt(productId, 10),
        quantity: 1
      })
      await useCartStore.getState().fetchCart()
    } catch (err: any) {
      if (err.response?.status === 401) return
      console.error('Failed to add item:', err)
      throw err
    }
  },

  removeItem: async (itemId: number) => {
    try {
      await api.delete(`/cart/delete/${itemId}`)
      await useCartStore.getState().fetchCart()
    } catch (err: any) {
      if (err.response?.status === 401) return
      console.error('Failed to remove item:', err)
      throw err
    }
  },

  updateQuantity: async (itemId: number, productId: number, quantity: number) => {
    try {
      await api.put(`/cart/update/${itemId}`, {
        product_id: productId,
        quantity
      })
      await useCartStore.getState().fetchCart()
    } catch (err: any) {
      if (err.response?.status === 401) return
      console.error('Failed to update quantity:', err)
      throw err
    }
  },

  clearCart: async () => {
    try {
      const resp = await api.get('/cart')
      const cartItems: CartItem[] = resp.data || []
      for (const item of cartItems) {
        await api.delete(`/cart/delete/${item.id}`)
      }
      set({ items: [] })
    } catch (err: any) {
      if (err.response?.status === 401) {
        set({ items: [] })
        return
      }
      console.error('Failed to clear cart:', err)
      throw err
    }
  }
}))
