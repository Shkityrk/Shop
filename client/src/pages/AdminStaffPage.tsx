import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { api } from '../api/client';
import { useAuthStore } from '../store/useAuthStore';

interface StaffFormData {
  username: string;
  email: string;
  password: string;
  first_name: string;
  last_name: string;
  user_role: 'admin' | 'staff' | 'warehouse' | 'courier';
}

const roleLabels: Record<string, string> = {
  admin: 'Администратор',
  staff: 'Работник',
  warehouse: 'Сотрудник склада',
  courier: 'Курьер',
};

export function AdminStaffPage() {
  const navigate = useNavigate();
  const { user, fetchUser, isLoading: authLoading } = useAuthStore();

  // Загружаем данные пользователя при монтировании
  useEffect(() => {
    if (!user) {
      fetchUser();
    }
  }, [user, fetchUser]);

  const [formData, setFormData] = useState<StaffFormData>({
    username: '',
    email: '',
    password: '',
    first_name: '',
    last_name: '',
    user_role: 'staff',
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

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

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);
    setSuccess(null);

    try {
      await api.post('/auth/register', formData);
      setSuccess(`Сотрудник "${formData.username}" (${roleLabels[formData.user_role]}) успешно создан!`);
      setFormData({
        username: '',
        email: '',
        password: '',
        first_name: '',
        last_name: '',
        user_role: 'staff',
      });
    } catch (e: any) {
      setError(e.response?.data?.detail || e.response?.data?.error || 'Ошибка при создании сотрудника');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      <div className="text-center mb-8">
        <h1 className="text-3xl font-bold text-gray-800">Регистрация сотрудника</h1>
        <p className="text-gray-600 mt-2">Создание нового аккаунта для работника</p>
      </div>

      <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg shadow-md space-y-4">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Имя
            </label>
            <input
              type="text"
              name="first_name"
              value={formData.first_name}
              onChange={handleChange}
              className="w-full border p-2 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500"
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Фамилия
            </label>
            <input
              type="text"
              name="last_name"
              value={formData.last_name}
              onChange={handleChange}
              className="w-full border p-2 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500"
              required
            />
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Логин (username)
          </label>
          <input
            type="text"
            name="username"
            value={formData.username}
            onChange={handleChange}
            className="w-full border p-2 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500"
            required
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Email
          </label>
          <input
            type="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            className="w-full border p-2 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500"
            required
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Пароль
          </label>
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            className="w-full border p-2 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500"
            required
            minLength={6}
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Роль
          </label>
          <select
            name="user_role"
            value={formData.user_role}
            onChange={handleChange}
            className="w-full border p-2 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500"
            required
          >
            <option value="admin">👑 Администратор (admin)</option>
            <option value="staff">👔 Работник (staff)</option>
            <option value="warehouse">🏭 Сотрудник склада (warehouse)</option>
            <option value="courier">🚚 Курьер (courier)</option>
          </select>
          <p className="text-xs text-gray-500 mt-1">
            {formData.user_role === 'admin' && 'Полный доступ ко всему'}
            {formData.user_role === 'staff' && 'Полный доступ к админ-панели'}
            {formData.user_role === 'warehouse' && 'Доступ к управлению складами и инвентарём'}
            {formData.user_role === 'courier' && 'Доступ к просмотру и управлению доставками'}
          </p>
        </div>

        {error && (
          <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg">
            {error}
          </div>
        )}

        {success && (
          <div className="bg-green-50 border border-green-200 text-green-700 px-4 py-3 rounded-lg">
            {success}
          </div>
        )}

        <button
          type="submit"
          disabled={isLoading}
          className="w-full bg-amber-600 text-white py-3 rounded-lg hover:bg-amber-700 transition-colors disabled:opacity-50"
        >
          {isLoading ? 'Создание...' : 'Создать сотрудника'}
        </button>
      </form>

      <div className="bg-amber-50 p-4 rounded-lg">
        <h3 className="font-semibold text-amber-800 mb-2">Описание ролей:</h3>
        <ul className="text-sm text-amber-700 space-y-1">
          <li><strong>Работник (staff)</strong> — полный доступ ко всем разделам админ-панели</li>
          <li><strong>Сотрудник склада (warehouse)</strong> — управление складами, полками и инвентарём</li>
          <li><strong>Курьер (courier)</strong> — просмотр и обновление статусов доставок</li>
        </ul>
      </div>
    </div>
  );
}
