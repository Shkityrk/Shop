import { Link, useNavigate } from 'react-router-dom';
import { ShoppingCart, Plus, Minus } from 'lucide-react';
import { Product } from '../types';
import { useCartStore } from '../store/useCartStore';
import { useAuthStore } from '../store/useAuthStore';

interface ProductCardProps {
  product: Product;
}

export function ProductCard({ product }: ProductCardProps) {
  const navigate = useNavigate();
  const { addItem, items, updateQuantity } = useCartStore();
  const { isAuthenticated } = useAuthStore();

  const cartItem = items.find(item => item.product_id === product.id);

  const handleAddToCart = async () => {
    if (!isAuthenticated) {
      navigate('/login');
      return;
    }

    try {
      await addItem(product.id.toString());
    } catch (error) {
      console.error('Failed to add item to cart:', error);
      alert('Failed to add item to cart. Please try again.');
    }
  };

  const handleUpdateQuantity = async (newQuantity: number) => {
    if (!cartItem) return;
    try {
      await updateQuantity(cartItem.id, product.id, newQuantity);
    } catch (error) {
      console.error('Failed to update quantity:', error);
    }
  };

  return (
      <div className="group block overflow-hidden rounded-lg bg-white shadow-md">
        <Link
            to={`/product/${product.id}`}
            className="block overflow-hidden"
        >
          <div className="aspect-square overflow-hidden">
            <img
                src={product.photo || 'https://images.unsplash.com/photo-1555507036-ab1f4038808a?w=800'}
                alt={product.name}
                className="h-full w-full object-cover transition-transform group-hover:scale-105"
            />
          </div>
          <div className="p-4">
            <h3 className="text-xl font-semibold text-amber-900">{product.name}</h3>
            <p className="mt-2 text-sm text-gray-600 line-clamp-2">{product.short_description}</p>
            <div className="mt-4 flex items-center justify-between">
              <span className="text-lg font-bold text-amber-600">{product.price.toFixed(2)} ₽</span>
              <span className="text-sm text-gray-500">{product.weight}g</span>
            </div>
          </div>
        </Link>
        <div className="p-4 pt-0">
          {isAuthenticated && cartItem ? (
              <div className="flex items-center justify-between gap-2">
                <button
                    onClick={() => handleUpdateQuantity(cartItem.quantity - 1)}
                    className="p-2 text-amber-600 hover:bg-amber-50 rounded-md transition-colors"
                    disabled={cartItem.quantity <= 1}
                >
                  <Minus className="h-4 w-4" />
                </button>
                <span className="font-medium text-amber-900">{cartItem.quantity}</span>
                <button
                    onClick={() => handleUpdateQuantity(cartItem.quantity + 1)}
                    className="p-2 text-amber-600 hover:bg-amber-50 rounded-md transition-colors"
                >
                  <Plus className="h-4 w-4" />
                </button>
              </div>
          ) : (
              <button
                  onClick={handleAddToCart}
                  className="w-full flex items-center justify-center bg-amber-600 text-white py-2 px-4 rounded-md hover:bg-amber-700 transition-colors"
              >
                <ShoppingCart className="h-4 w-4 mr-2" />
                Add to Cart
              </button>
          )}
        </div>
      </div>
  );
}