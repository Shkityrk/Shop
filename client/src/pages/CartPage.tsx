import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Minus, Plus, Trash2, ArrowLeft } from 'lucide-react';
import { useCartStore } from '../store/useCartStore';
import { useAuthStore } from '../store/useAuthStore';
import { api } from '../api/client';

interface CartProduct {
  id: number;
  product_id: number;
  quantity: number;
  user_id: number;
  product?: {
    name: string;
    price: number;
    photo: string;
    short_description: string;
  };
}

export function CartPage() {
  const navigate = useNavigate();
  const { updateQuantity, removeItem } = useCartStore();
  const { isAuthenticated } = useAuthStore();
  const [isProcessing, setIsProcessing] = useState(false);
  const [cartProducts, setCartProducts] = useState<CartProduct[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchCartAndProducts = async () => {
      try {
        if (!isAuthenticated) {
          setCartProducts([]);
          setIsLoading(false);
          return;
        }

        // Fetch cart items
        const cartResponse = await api.get('/cart');
        const cartItems = cartResponse.data || [];

        // Fetch product details for each cart item
        const productsResponse = await api.get('/product/list');
        const products = productsResponse.data || [];

        // Combine cart items with product details
        const cartWithProducts = cartItems.map((cartItem: CartProduct) => ({
          ...cartItem,
          product: products.find((p: any) => p.id === cartItem.product_id)
        }));

        setCartProducts(cartWithProducts);
      } catch (error) {
        console.error('Failed to fetch cart products:', error);
        setError('Failed to load cart items. Please try again.');
      } finally {
        setIsLoading(false);
      }
    };

    fetchCartAndProducts();
  }, [isAuthenticated]);

  const total = cartProducts.reduce(
    (sum, item) => sum + ((item.product?.price || 0) * item.quantity),
    0
  );

  const handleQuantityChange = async (productId: number, newQuantity: number) => {
    if (newQuantity < 1) return;
    try {
      await updateQuantity(productId.toString(), newQuantity);
      // Refresh cart data
      const response = await api.get('/cart');
      const cartItems = response.data || [];
      const productsResponse = await api.get('/product/list');
      const products = productsResponse.data || [];
      const cartWithProducts = cartItems.map((cartItem: CartProduct) => ({
        ...cartItem,
        product: products.find((p: any) => p.id === cartItem.product_id)
      }));
      setCartProducts(cartWithProducts);
    } catch (error) {
      console.error('Failed to update quantity:', error);
    }
  };

  const handleRemoveItem = async (productId: number) => {
    try {
      await removeItem(productId.toString());
      setCartProducts(prev => prev.filter(item => item.product_id !== productId));
    } catch (error) {
      console.error('Failed to remove item:', error);
    }
  };

  const handleCheckout = async () => {
    if (!isAuthenticated) {
      navigate('/login');
      return;
    }

    setIsProcessing(true);
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000));
    setIsProcessing(false);
    // Navigate to success page or show confirmation
    alert('Order placed successfully!');
  };

  if (!isAuthenticated) {
    return (
      <div className="flex flex-col items-center justify-center min-h-[400px] space-y-4">
        <p className="text-lg text-gray-600">Please login to view your cart</p>
        <Link
          to="/login"
          className="bg-amber-600 text-white px-6 py-2 rounded-lg hover:bg-amber-700 transition-colors"
        >
          Login
        </Link>
      </div>
    );
  }

  if (isLoading) {
    return (
      <div className="flex justify-center items-center min-h-[400px]">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-amber-600"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex flex-col items-center justify-center min-h-[400px] space-y-4">
        <p className="text-lg text-red-600">{error}</p>
        <button
          onClick={() => window.location.reload()}
          className="text-amber-600 hover:text-amber-700"
        >
          Try again
        </button>
      </div>
    );
  }

  if (cartProducts.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center min-h-[400px] space-y-4">
        <p className="text-lg text-gray-600">Your cart is empty</p>
        <Link
          to="/"
          className="flex items-center text-amber-600 hover:text-amber-700"
        >
          <ArrowLeft className="h-4 w-4 mr-2" />
          Continue Shopping
        </Link>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto">
      <h1 className="text-3xl font-bold text-amber-900 mb-8">Shopping Cart</h1>
      
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2 space-y-4">
          {cartProducts.map((item) => (
            <div
              key={item.id}
              className="flex items-center gap-4 bg-white p-4 rounded-lg shadow-sm"
            >
              <img
                src={item.product?.photo || 'https://images.unsplash.com/photo-1555507036-ab1f4038808a?w=800'}
                alt={item.product?.name}
                className="w-24 h-24 object-cover rounded"
              />
              
              <div className="flex-1">
                <h3 className="font-semibold text-amber-900">{item.product?.name}</h3>
                <p className="text-sm text-gray-500">{item.product?.short_description}</p>
                <p className="font-bold text-amber-600 mt-1">
                  ${(item.product?.price || 0).toFixed(2)}
                </p>
              </div>
              
              <div className="flex items-center gap-2">
                <button
                  onClick={() => handleQuantityChange(item.product_id, item.quantity - 1)}
                  className="p-1 rounded hover:bg-gray-100"
                >
                  <Minus className="h-4 w-4" />
                </button>
                <span className="w-8 text-center">{item.quantity}</span>
                <button
                  onClick={() => handleQuantityChange(item.product_id, item.quantity + 1)}
                  className="p-1 rounded hover:bg-gray-100"
                >
                  <Plus className="h-4 w-4" />
                </button>
              </div>
              
              <button
                onClick={() => handleRemoveItem(item.product_id)}
                className="p-2 text-red-500 hover:text-red-700 transition-colors"
              >
                <Trash2 className="h-5 w-5" />
              </button>
            </div>
          ))}
        </div>
        
        <div className="bg-white p-6 rounded-lg shadow-sm h-fit">
          <h2 className="text-xl font-semibold text-amber-900 mb-4">Order Summary</h2>
          
          <div className="space-y-2 mb-4">
            <div className="flex justify-between">
              <span>Subtotal</span>
              <span>${total.toFixed(2)}</span>
            </div>
            <div className="flex justify-between font-bold text-lg border-t pt-2">
              <span>Total</span>
              <span>${total.toFixed(2)}</span>
            </div>
          </div>
          
          <button
            onClick={handleCheckout}
            disabled={isProcessing}
            className="w-full bg-amber-600 text-white py-3 rounded-lg hover:bg-amber-700 transition-colors disabled:opacity-50"
          >
            {isProcessing ? 'Processing...' : 'Checkout'}
          </button>
          
          <Link
            to="/"
            className="block text-center text-amber-600 hover:text-amber-700 mt-4"
          >
            Continue Shopping
          </Link>
        </div>
      </div>
    </div>
  );
}