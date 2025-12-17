import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { api } from '../api/client';
import { useAuthStore } from '../store/useAuthStore';

interface StorageRule {
  id: number;
  name: string;
  description?: string;
  is_hazardous: boolean;
  is_oversized: boolean;
  temp_min?: number | null;
  temp_max?: number | null;
}

interface StorageRuleForm {
  name: string;
  description: string;
  is_hazardous: boolean;
  is_oversized: boolean;
  temp_min: string;
  temp_max: string;
}

export function AdminStorageRulesPage() {
  const navigate = useNavigate();
  const { user, fetchUser, isLoading: authLoading } = useAuthStore();

  const [rules, setRules] = useState<StorageRule[]>([]);
  const [loading, setLoading] = useState(false);
  const [form, setForm] = useState<StorageRuleForm>({
    name: '',
    description: '',
    is_hazardous: false,
    is_oversized: false,
    temp_min: '',
    temp_max: '',
  });
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!user) {
      fetchUser();
    }
  }, [user, fetchUser]);

  const loadRules = async () => {
    setLoading(true);
    try {
      const resp = await api.get('/warehouse/storage-rules');
      setRules(resp.data || []);
    } catch (e) {
      console.error('Failed to load storage rules:', e);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadRules();
  }, []);

  if (authLoading) {
    return (
      <div className="flex items-center justify-center min-h-[400px]">
        <div className="text-lg text-gray-600">Загрузка...</div>
      </div>
    );
  }

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
    const { name, value, type } = e.target;
    if (type === 'checkbox') {
      const checked = (e.target as HTMLInputElement).checked;
      setForm((prev) => ({ ...prev, [name]: checked }));
    } else {
      setForm((prev) => ({ ...prev, [name]: value }));
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      const payload = {
        name: form.name,
        description: form.description || null,
        is_hazardous: form.is_hazardous,
        is_oversized: form.is_oversized,
        temp_min: form.temp_min ? parseFloat(form.temp_min) : null,
        temp_max: form.temp_max ? parseFloat(form.temp_max) : null,
      };

      await api.post('/warehouse/storage-rules', payload);
      setForm({
        name: '',
        description: '',
        is_hazardous: false,
        is_oversized: false,
        temp_min: '',
        temp_max: '',
      });
      await loadRules();
    } catch (e: any) {
      setError(e.response?.data?.detail || e.response?.data?.error || 'Ошибка создания');
    }
  };

  const handleDelete = async (id: number) => {
    if (!confirm('Удалить это условие хранения?')) return;
    try {
      await api.delete(`/warehouse/storage-rules/${id}`);
      await loadRules();
    } catch (e) {
      console.error('Failed to delete:', e);
    }
  };

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold mb-4">Условия хранения</h1>

      {/* Форма создания */}
      <form onSubmit={handleSubmit} className="bg-white p-4 rounded shadow space-y-4">
        <h2 className="font-semibold text-lg">Создать новое условие</h2>

        <div className="grid md:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Название *
            </label>
            <input
              name="name"
              value={form.name}
              onChange={handleChange}
              className="border p-2 rounded w-full"
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Описание
            </label>
            <input
              name="description"
              value={form.description}
              onChange={handleChange}
              className="border p-2 rounded w-full"
            />
          </div>
        </div>

        <div className="grid md:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Мин. температура (°C)
            </label>
            <input
              name="temp_min"
              type="number"
              step="0.1"
              value={form.temp_min}
              onChange={handleChange}
              className="border p-2 rounded w-full"
              placeholder="например: -18"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Макс. температура (°C)
            </label>
            <input
              name="temp_max"
              type="number"
              step="0.1"
              value={form.temp_max}
              onChange={handleChange}
              className="border p-2 rounded w-full"
              placeholder="например: 4"
            />
          </div>
        </div>

        <div className="flex gap-6">
          <label className="flex items-center gap-2">
            <input
              type="checkbox"
              name="is_hazardous"
              checked={form.is_hazardous}
              onChange={handleChange}
              className="rounded"
            />
            <span className="text-sm">Опасный груз</span>
          </label>

          <label className="flex items-center gap-2">
            <input
              type="checkbox"
              name="is_oversized"
              checked={form.is_oversized}
              onChange={handleChange}
              className="rounded"
            />
            <span className="text-sm">Крупногабаритный</span>
          </label>
        </div>

        {error && (
          <p className="text-red-600 text-sm">{error}</p>
        )}

        <button
          type="submit"
          className="bg-amber-600 text-white py-2 px-4 rounded hover:bg-amber-700"
        >
          Создать условие
        </button>
      </form>

      {/* Список */}
      <div className="bg-white rounded shadow p-4">
        <h2 className="font-semibold mb-4">Список условий хранения</h2>

        {loading ? (
          <p>Загрузка...</p>
        ) : rules.length === 0 ? (
          <p className="text-gray-500">Условия хранения не созданы</p>
        ) : (
          <div className="overflow-x-auto">
            <table className="w-full text-sm">
              <thead>
                <tr className="border-b bg-gray-50">
                  <th className="text-left p-3">ID</th>
                  <th className="text-left p-3">Название</th>
                  <th className="text-left p-3">Описание</th>
                  <th className="text-left p-3">Температура</th>
                  <th className="text-left p-3">Флаги</th>
                  <th className="text-left p-3">Действия</th>
                </tr>
              </thead>
              <tbody>
                {rules.map((rule) => (
                  <tr key={rule.id} className="border-b hover:bg-gray-50">
                    <td className="p-3">{rule.id}</td>
                    <td className="p-3 font-medium">{rule.name}</td>
                    <td className="p-3 text-gray-600">{rule.description || '—'}</td>
                    <td className="p-3">
                      {rule.temp_min !== null || rule.temp_max !== null ? (
                        <span className="text-blue-600">
                          {rule.temp_min ?? '—'}°C — {rule.temp_max ?? '—'}°C
                        </span>
                      ) : (
                        <span className="text-gray-400">Не указано</span>
                      )}
                    </td>
                    <td className="p-3">
                      <div className="flex gap-2">
                        {rule.is_hazardous && (
                          <span className="bg-red-100 text-red-700 px-2 py-1 rounded text-xs">
                            ⚠️ Опасный
                          </span>
                        )}
                        {rule.is_oversized && (
                          <span className="bg-yellow-100 text-yellow-700 px-2 py-1 rounded text-xs">
                            📦 Крупный
                          </span>
                        )}
                        {!rule.is_hazardous && !rule.is_oversized && (
                          <span className="text-gray-400">—</span>
                        )}
                      </div>
                    </td>
                    <td className="p-3">
                      <button
                        onClick={() => handleDelete(rule.id)}
                        className="text-red-600 hover:text-red-800"
                      >
                        Удалить
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}

