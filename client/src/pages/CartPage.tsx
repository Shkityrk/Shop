import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Minus, Plus, Trash2, ArrowLeft, CheckCircle } from 'lucide-react';
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
  const { updateQuantity, removeItem, clearCart } = useCartStore();
  const { isAuthenticated, user } = useAuthStore();
  const [isProcessing, setIsProcessing] = useState(false);
  const [cartProducts, setCartProducts] = useState<CartProduct[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [wmsMessage, setWmsMessage] = useState<string | null>(null);
  const [orderSuccess, setOrderSuccess] = useState<{ trackingCode: string } | null>(null);
  const [address, setAddress] = useState('');

  useEffect(() => {
    const fetchCartAndProducts = async () => {
      try {
        if (!isAuthenticated) {
          setCartProducts([]);
          setIsLoading(false);
          return;
        }

        const cartResponse = await api.get('/cart');
        const cartItems = cartResponse.data || [];

        const productsResponse = await api.get('/product/list');
        const products = productsResponse.data || [];

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

  const buildWmsPayload = () => ({
    items: cartProducts.map((item) => ({
      product_id: item.product_id,
      quantity: item.quantity,
    })),
  });

  const handleWmsCheck = async () => {
    setWmsMessage(null);
    try {
      const resp = await api.post('/warehouse/wms/check', buildWmsPayload());
      if (resp.data.ok) {
        setWmsMessage('✅ Все товары есть на складах. Можно оформить заказ!');
      } else {
        setWmsMessage('❌ Недостаточно товара: ' + JSON.stringify(resp.data.shortages));
      }
    } catch (e: any) {
      setWmsMessage('❌ ' + (e.response?.data?.detail || 'Ошибка проверки на складах'));
    }
  };

  const handleCheckout = async () => {
    if (!isAuthenticated) {
      navigate('/login');
      return;
    }

    if (!address.trim()) {
      setWmsMessage('❌ Пожалуйста, укажите адрес доставки');
      return;
    }

    setIsProcessing(true);
    setWmsMessage(null);

    try {
      // 1. Проверяем наличие товаров на складах
      const checkResp = await api.post('/warehouse/wms/check', buildWmsPayload());
      if (!checkResp.data.ok) {
        setWmsMessage('❌ Недостаточно товара на складе: ' + JSON.stringify(checkResp.data.shortages));
        setIsProcessing(false);
        return;
      }

      // 2. Создаём доставку (shipping)
      // Получаем user_id из корзины (там он точно есть) или из user
      const userId = cartProducts[0]?.user_id || user?.id;
      if (!userId) {
        setWmsMessage('❌ Не удалось определить пользователя');
        setIsProcessing(false);
        return;
      }

      const shippingPayload = {
        order_id: Date.now(),
        user_id: userId,
        address: address.trim(),
        items: cartProducts.map((item) => ({
          product_id: item.product_id,
          quantity: item.quantity,
        })),
      };

      const shippingResp = await api.post('/shipping', shippingPayload);
      const trackingCode = shippingResp.data.tracking_code;

      // 3. Списываем товары со склада
      const commitResp = await api.post('/warehouse/wms/commit', buildWmsPayload());
      if (!commitResp.data.ok) {
        setWmsMessage('⚠️ Заказ создан, но возникла проблема со списанием: ' + JSON.stringify(commitResp.data.shortages));
        setIsProcessing(false);
        return;
      }

      // 4. Очищаем корзину
      await clearCart();
      setCartProducts([]);

      // 5. Показываем успех
      setOrderSuccess({ trackingCode });
      setWmsMessage(null);

    } catch (e: any) {
      setWmsMessage('❌ ' + (e.response?.data?.detail || e.response?.data?.error || 'Ошибка оформления заказа'));
    } finally {
      setIsProcessing(false);
    }
  };

  const handleQuantityChange = async (itemId: number, productId: number, newQuantity: number) => {
    if (newQuantity < 1) return;
    try {
      await updateQuantity(itemId, productId, newQuantity);
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

  const handleRemoveItem = async (itemId: number) => {
    try {
      await removeItem(itemId);
      setCartProducts(prev => prev.filter(item => item.id !== itemId));
    } catch (error) {
      console.error('Failed to remove item:', error);
    }
  };

  // Экран успешного заказа
  if (orderSuccess) {
    return (
      <div className="flex flex-col items-center justify-center min-h-[400px] space-y-6">
        <CheckCircle className="h-20 w-20 text-green-500" />
        <h1 className="text-3xl font-bold text-green-700">Заказ оформлен!</h1>
        <div className="bg-white p-6 rounded-lg shadow-md text-center">
          <p className="text-gray-600 mb-2">Ваш трек-номер для отслеживания:</p>
          <p className="text-2xl font-mono font-bold text-amber-600 bg-amber-50 px-4 py-2 rounded">
            {orderSuccess.trackingCode}
          </p>
        </div>
        <p className="text-gray-500">Сохраните этот номер для отслеживания доставки</p>
        <Link
          to="/"
          className="bg-amber-600 text-white px-6 py-3 rounded-lg hover:bg-amber-700 transition-colors"
        >
          Вернуться на главную
        </Link>
      </div>
    );
  }

  if (!isAuthenticated) {
    return (
      <div className="flex flex-col items-center justify-center min-h-[400px] space-y-4">
        <p className="text-lg text-gray-600">Войдите, чтобы видеть корзину</p>
        <Link
          to="/login"
          className="bg-amber-600 text-white px-6 py-2 rounded-lg hover:bg-amber-700 transition-colors"
        >
          Войти
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
        <p className="text-lg text-gray-600">Ваша корзина пустая</p>
        <Link
          to="/"
          className="flex items-center text-amber-600 hover:text-amber-700"
        >
          <ArrowLeft className="h-4 w-4 mr-2" />
          Продолжить покупки
        </Link>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto">
      <h1 className="text-3xl font-bold text-amber-900 mb-8">Корзина</h1>

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
                  {(item.product?.price || 0).toFixed(2)} ₽
                </p>
              </div>

              <div className="flex items-center gap-2">
                <button
                  onClick={() => handleQuantityChange(item.id, item.product_id, item.quantity - 1)}
                  className="p-1 rounded hover:bg-gray-100"
                >
                  <Minus className="h-4 w-4" />
                </button>
                <span className="w-8 text-center">{item.quantity}</span>
                <button
                  onClick={() => handleQuantityChange(item.id, item.product_id, item.quantity + 1)}
                  className="p-1 rounded hover:bg-gray-100"
                >
                  <Plus className="h-4 w-4" />
                </button>
              </div>

              <button
                onClick={() => handleRemoveItem(item.id)}
                className="p-2 text-red-500 hover:text-red-700 transition-colors"
              >
                <Trash2 className="h-5 w-5" />
              </button>
            </div>
          ))}
        </div>

        <div className="bg-white p-6 rounded-lg shadow-sm h-fit">
          <h2 className="text-xl font-semibold text-amber-900 mb-4">Оформление заказа</h2>

          <div className="space-y-2 mb-4">
            <div className="flex justify-between">
              <span>Предварительная цена</span>
              <span>{total.toFixed(2)} ₽</span>
            </div>
            <div className="flex justify-between font-bold text-lg border-t pt-2">
              <span>Всего</span>
              <span>{total.toFixed(2)} ₽</span>
            </div>
          </div>

          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Адрес доставки
            </label>
            <input
              type="text"
              value={address}
              onChange={(e) => setAddress(e.target.value)}
              placeholder="Введите адрес доставки"
              className="w-full border p-2 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500"
            />
          </div>

          <button
            onClick={handleWmsCheck}
            disabled={isProcessing}
            className="w-full mb-2 bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50"
          >
            Проверить наличие
          </button>
          <button
            onClick={handleCheckout}
            disabled={isProcessing || !address.trim()}
            className="w-full bg-amber-600 text-white py-3 rounded-lg hover:bg-amber-700 transition-colors disabled:opacity-50"
          >
            {isProcessing ? 'Обработка...' : 'Оформить заказ'}
          </button>
          {wmsMessage && (
            <p className="mt-3 text-sm text-gray-800 whitespace-pre-wrap">{wmsMessage}</p>
          )}

          <Link
            to="/"
            className="block text-center text-amber-600 hover:text-amber-700 mt-4"
          >
            Продолжить покупки
          </Link>
        </div>
      </div>
    </div>
  );
}

