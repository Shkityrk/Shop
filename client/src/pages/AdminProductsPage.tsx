import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { api } from '../api/client';
import { useAuthStore } from '../store/useAuthStore';
import type { Product } from '../types';

export function AdminProductsPage() {
  const navigate = useNavigate();
  const { user, fetchUser, isLoading: authLoading } = useAuthStore();

  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(false);
  const [form, setForm] = useState<Omit<Product, 'id'>>({
    name: '',
    short_description: '',
    full_description: '',
    composition: '',
    weight: 0,
    price: 0,
    photo: '',
  });

  // Загружаем данные пользователя при монтировании
  useEffect(() => {
    if (!user) {
      fetchUser();
    }
  }, [user, fetchUser]);

  const load = async () => {
    setLoading(true);
    try {
      const resp = await api.get('/product/list');
      setProducts(resp.data || []);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    load();
  }, []);

  // Показываем загрузку пока получаем данные пользователя
  if (authLoading) {
    return (
      <div className="flex items-center justify-center min-h-[400px]">
        <div className="text-lg text-gray-600">Загрузка...</div>
      </div>
    );
  }

  // Проверяем доступ
  if (!user || user.user_role === 'client') {
    return (
      <div className="flex flex-col items-center justify-center min-h-[400px] space-y-4">
        <p className="text-lg text-red-600">Доступ запрещён</p>
        <button
          onClick={() => navigate('/')}
          className="text-amber-600 hover:text-amber-700"
        >
          На главную
        </button>
      </div>
    );
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: name === 'weight' || name === 'price' ? Number(value) : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await api.post('/product/add', form);
    setForm({
      name: '',
      short_description: '',
      full_description: '',
      composition: '',
      weight: 0,
      price: 0,
      photo: '',
    });
    await load();
  };

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold mb-4">Управление продуктами</h1>

      <form onSubmit={handleSubmit} className="grid gap-4 md:grid-cols-2 bg-white p-4 rounded shadow">
        <input
          name="name"
          value={form.name}
          onChange={handleChange}
          placeholder="Название"
          className="border p-2 rounded"
          required
        />
        <input
          name="price"
          type="number"
          value={form.price}
          onChange={handleChange}
          placeholder="Цена"
          className="border p-2 rounded"
          required
        />
        <input
          name="weight"
          type="number"
          value={form.weight}
          onChange={handleChange}
          placeholder="Вес"
          className="border p-2 rounded"
          required
        />
        <input
          name="photo"
          value={form.photo}
          onChange={handleChange}
          placeholder="URL фото"
          className="border p-2 rounded"
        />
        <textarea
          name="short_description"
          value={form.short_description}
          onChange={handleChange}
          placeholder="Краткое описание"
          className="border p-2 rounded md:col-span-2"
        />
        <textarea
          name="full_description"
          value={form.full_description}
          onChange={handleChange}
          placeholder="Полное описание"
          className="border p-2 rounded md:col-span-2"
        />
        <textarea
          name="composition"
          value={form.composition}
          onChange={handleChange}
          placeholder="Состав"
          className="border p-2 rounded md:col-span-2"
        />
        <button
          type="submit"
          className="bg-amber-600 text-white py-2 px-4 rounded hover:bg-amber-700 md:col-span-2"
        >
          Добавить продукт
        </button>
      </form>

      <div className="bg-white rounded shadow p-4">
        <h2 className="font-semibold mb-2">Список продуктов</h2>
        {loading ? (
          <p>Загрузка...</p>
        ) : (
          <table className="w-full text-sm">
            <thead>
            <tr className="border-b">
              <th className="text-left p-2">ID</th>
              <th className="text-left p-2">Название</th>
              <th className="text-left p-2">Цена</th>
              <th className="text-left p-2">Вес</th>
            </tr>
            </thead>
            <tbody>
            {products.map((p) => (
              <tr key={p.id} className="border-b">
                <td className="p-2">{p.id}</td>
                <td className="p-2">{p.name}</td>
                <td className="p-2">{p.price}</td>
                <td className="p-2">{p.weight}</td>
              </tr>
            ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
  );
}

