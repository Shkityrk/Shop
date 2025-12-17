import { ShoppingCart, User, Croissant } from 'lucide-react';
import { Link } from 'react-router-dom';
import { useAuthStore } from '../store/useAuthStore';
import { useCartStore } from '../store/useCartStore';

export function Header() {
  const { isAuthenticated, user } = useAuthStore();
  const { items } = useCartStore();
  const itemCount = items.reduce((acc, item) => acc + item.quantity, 0);

  return (
    <header className="bg-amber-100 shadow-md">
      <div className="container mx-auto px-4 py-4">
        <div className="flex items-center justify-between">
          <Link to="/" className="flex items-center space-x-2">
            {/*<Croissant className="h-8 w-8 text-amber-600" />*/}
            <span className="text-2xl font-bold text-amber-900">Магазин</span>
          </Link>
          
          <nav className="flex items-center space-x-6">
            {user && user.user_role && user.user_role !== 'client' && (
              <Link to="/admin" className="text-amber-900 hover:text-amber-700 font-medium">
                Админ-панель
              </Link>
            )}
            <Link
              to="/cart"
              className="relative flex items-center text-amber-900 hover:text-amber-700"
            >
              <ShoppingCart className="h-6 w-6" />
              {itemCount > 0 && (
                <span className="absolute -top-2 -right-2 flex h-5 w-5 items-center justify-center rounded-full bg-amber-600 text-xs text-white">
                  {itemCount}
                </span>
              )}
            </Link>
            
            <Link
              to={isAuthenticated ? "/profile" : "/login"}
              className="flex items-center text-amber-900 hover:text-amber-700"
            >
              <User className="h-6 w-6" />
            </Link>
          </nav>
        </div>
      </div>
    </header>
  );
}