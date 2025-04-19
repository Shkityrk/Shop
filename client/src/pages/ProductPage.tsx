import { useParams, useNavigate } from 'react-router-dom';
import { ArrowLeft, ShoppingCart } from 'lucide-react';
import { mockProducts } from '../data/mockProducts';
import { useCartStore } from '../store/useCartStore';

export function ProductPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const product = mockProducts.find(p => p.id === id);
  const addItem = useCartStore((state) => state.addItem);

  if (!product) {
    return (
      <div className="flex flex-col items-center justify-center min-h-[400px] space-y-4">
        <div className="text-red-600">Product not found</div>
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
            src={product.imageUrl}
            alt={product.name}
            className="h-full w-full object-cover"
          />
        </div>
        
        <div className="space-y-6">
          <div>
            <h1 className="text-3xl font-bold text-amber-900">{product.name}</h1>
            <p className="mt-2 text-sm text-gray-500">{product.category}</p>
          </div>
          
          <p className="text-gray-600">{product.description}</p>
          
          <div className="text-2xl font-bold text-amber-600">
            ${product.price.toFixed(2)}
          </div>
          
          <button
            onClick={() => {
              addItem(product.id);
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