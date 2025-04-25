import { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { ArrowLeft, ShoppingCart } from 'lucide-react';
import { useCartStore } from '../store/useCartStore';
import { api } from '../api/client';
import { Product } from '../types';

export function ProductPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [product, setProduct] = useState<Product | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const addItem = useCartStore((state) => state.addItem);

  useEffect(() => {
    const fetchProduct = async () => {
      try {
        const response = await api.get(`/product/info/${id}`);
        setProduct(response.data.product);
      } catch (err) {
        console.error('Error fetching product:', err);
        setError('Failed to load product information');
      } finally {
        setIsLoading(false);
      }
    };

    fetchProduct();
  }, [id]);

  if (isLoading) {
    return (
      <div className="flex justify-center items-center min-h-[400px]">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-amber-600"></div>
      </div>
    );
  }

  if (error || !product) {
    return (
      <div className="flex flex-col items-center justify-center min-h-[400px] space-y-4">
        <div className="text-red-600">{error || 'Product not found'}</div>
        <button
          onClick={() => navigate('/')}
          className="flex items-center text-amber-600 hover:text-amber-700"
        >
          <ArrowLeft className="h-4 w-4 mr-2" />
          Back to products
        </button>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto">
      <button
        onClick={() => navigate('/')}
        className="flex items-center text-amber-600 hover:text-amber-700 mb-6"
      >
        <ArrowLeft className="h-4 w-4 mr-2" />
        Back to products
      </button>
      
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div className="aspect-square overflow-hidden rounded-lg">
          <img
            src={product.photo}
            alt={product.name}
            className="h-full w-full object-cover"
          />
        </div>
        
        <div className="space-y-6">
          <div>
            <h1 className="text-3xl font-bold text-amber-900">{product.name}</h1>
            <p className="mt-2 text-sm text-gray-500">{product.short_description}</p>
          </div>
          
          <p className="text-gray-600">{product.full_description}</p>
          
          <div className="text-2xl font-bold text-amber-600">
            ${product.price.toFixed(2)}
          </div>

          <div className="text-sm text-gray-500">
            Weight: {product.weight}g
          </div>

          <div className="text-sm text-gray-600">
            <h3 className="font-semibold mb-1">Composition:</h3>
            <p>{product.composition}</p>
          </div>
          
          <button
            onClick={() => {
              addItem(product.id.toString());
              navigate('/cart');
            }}
            className="flex items-center justify-center w-full bg-amber-600 text-white py-3 px-6 rounded-lg hover:bg-amber-700 transition-colors"
          >
            <ShoppingCart className="h-5 w-5 mr-2" />
            Add to Cart
          </button>
        </div>
      </div>
    </div>
  );
}