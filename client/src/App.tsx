import { useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Header } from './components/Header';
import { HomePage } from './pages/HomePage';
import { ProductPage } from './pages/ProductPage';
import { LoginPage } from './pages/LoginPage';
import { RegisterPage } from './pages/RegisterPage';
import { CartPage } from './pages/CartPage';
import { ProfilePage } from './pages/ProfilePage';
import { useAuthStore } from './store/useAuthStore';
import { useCartStore } from './store/useCartStore';

function App() {
  const { isAuthenticated } = useAuthStore();
  const { fetchCart } = useCartStore();

  useEffect(() => {
    let mounted = true;

    const loadCart = async () => {
      if (isAuthenticated && mounted) {
        await fetchCart();
      }
    };

    loadCart();

    return () => {
      mounted = false;
    };
  }, [isAuthenticated]); // Remove fetchCart from dependencies

  return (
      <Router>
        <div className="min-h-screen bg-amber-50">
          <Header />
          <main className="container mx-auto px-4 py-8">
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/product/:id" element={<ProductPage />} />
              <Route path="/cart" element={<CartPage />} />
              <Route path="/profile" element={<ProfilePage />} />
              <Route path="/login" element={<LoginPage />} />
              <Route path="/register" element={<RegisterPage />} />
            </Routes>
          </main>
        </div>
      </Router>
  );
}
export default App;