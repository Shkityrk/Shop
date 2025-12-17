import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { api } from '../api/client';
import { useAuthStore } from '../store/useAuthStore';

interface ShipmentItem {
  id: number;
  product_id: number;
  quantity: number;
}

interface Shipment {
  id: number;
  order_id: number;
  user_id: number;
  address: string;
  tracking_code: string;
  status: string;
  courier_id?: number | null;
  created_at: string;
  updated_at: string;
  items: ShipmentItem[];
}

interface Staff {
  id: number;
  first_name: string;
  last_name: string;
  user_role: string;
}

export function AdminShipmentsPage() {
  const navigate = useNavigate();
  const { user, fetchUser, isLoading: authLoading } = useAuthStore();

  const [shipments, setShipments] = useState<Shipment[]>([]);
  const [couriers, setCouriers] = useState<Staff[]>([]);
  const [filterStatus, setFilterStatus] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [editingId, setEditingId] = useState<number | null>(null);
  const [editStatus, setEditStatus] = useState('');
  const [editCourierId, setEditCourierId] = useState('');

  // Загружаем данные пользователя при монтировании
  useEffect(() => {
    if (!user) {
      fetchUser();
    }
  }, [user, fetchUser]);

  // Загружаем список курьеров
  useEffect(() => {
    const loadCouriers = async () => {
      try {
        const resp = await api.get('/auth/staff');
        // Фильтруем только курьеров
        const allStaff: Staff[] = resp.data || [];
        setCouriers(allStaff.filter((s) => s.user_role === 'courier'));
      } catch (e) {
        console.error('Failed to load couriers:', e);
      }
    };
    loadCouriers();
  }, []);

  const fetchShipments = async () => {
    setLoading(true);
    setError(null);
    try {
      const params = new URLSearchParams();
      if (filterStatus) params.append('status', filterStatus);
      const resp = await api.get(`/shipping/list?${params.toString()}`);
      setShipments(resp.data);
    } catch (e: any) {
      setError(e.response?.data?.error || 'Не удалось загрузить доставки');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchShipments();
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

  const handleFilterApply = () => {
    fetchShipments();
  };

  const startEdit = (shipment: Shipment) => {
    setEditingId(shipment.id);
    setEditStatus(shipment.status);
    setEditCourierId(shipment.courier_id?.toString() ?? '');
  };

  const cancelEdit = () => {
    setEditingId(null);
    setEditStatus('');
    setEditCourierId('');
  };

  const handleUpdate = async (trackingCode: string) => {
    setLoading(true);
    setError(null);
    try {
      const payload: any = { status: editStatus };
      if (editCourierId) payload.courier_id = Number(editCourierId);
      await api.patch(`/shipping/${trackingCode}/status`, payload);
      cancelEdit();
      fetchShipments();
    } catch (e: any) {
      setError(e.response?.data?.error || 'Не удалось обновить статус');
    } finally {
      setLoading(false);
    }
  };

  const statusOptions = ['created', 'assigned', 'in_transit', 'delivered', 'cancelled'];

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold mb-4">Доставка</h1>

      {/* Фильтр */}
      <div className="bg-white p-4 rounded shadow space-y-3">
        <div className="flex gap-2 items-center flex-wrap">
          <select
            value={filterStatus}
            onChange={(e) => setFilterStatus(e.target.value)}
            className="border p-2 rounded"
          >
            <option value="">Все статусы</option>
            {statusOptions.map((s) => (
              <option key={s} value={s}>{s}</option>
            ))}
          </select>
          <button
            onClick={handleFilterApply}
            className="bg-amber-600 text-white px-4 py-2 rounded hover:bg-amber-700"
          >
            Применить фильтр
          </button>
          <button
            onClick={() => { setFilterStatus(''); fetchShipments(); }}
            className="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400"
          >
            Сбросить
          </button>
        </div>
        {error && <p className="text-red-600 text-sm">{error}</p>}
      </div>

      {loading && <p>Загрузка...</p>}

      {/* Список доставок */}
      {!loading && shipments.length === 0 && (
        <p className="text-gray-500">Доставок не найдено</p>
      )}

      <div className="space-y-4">
        {shipments.map((shipment) => (
          <div key={shipment.id} className="bg-white p-4 rounded shadow space-y-3 text-sm">
            <div className="flex justify-between flex-wrap gap-2">
              <div>
                <div className="font-semibold">
                  Отправление #{shipment.id} (order #{shipment.order_id})
                </div>
                <div>Пользователь: {shipment.user_id}</div>
                <div>Адрес: {shipment.address}</div>
                <div>Tracking: <span className="font-mono bg-gray-100 px-1 rounded">{shipment.tracking_code}</span></div>
              </div>
              <div className="text-right">
                <div>
                  Статус: <span className={`font-semibold ${
                    shipment.status === 'delivered' ? 'text-green-600' :
                    shipment.status === 'cancelled' ? 'text-red-600' :
                    shipment.status === 'in_transit' ? 'text-blue-600' :
                    'text-gray-700'
                  }`}>{shipment.status}</span>
                </div>
                <div>Курьер: {
                  shipment.courier_id
                    ? couriers.find(c => c.id === shipment.courier_id)
                        ? `${couriers.find(c => c.id === shipment.courier_id)?.first_name} ${couriers.find(c => c.id === shipment.courier_id)?.last_name}`
                        : `#${shipment.courier_id}`
                    : '—'
                }</div>
                <div className="text-gray-500 text-xs">Создано: {new Date(shipment.created_at).toLocaleString()}</div>
              </div>
            </div>

            {/* Товары */}
            <div className="border-t pt-3">
              <h3 className="font-semibold mb-1">Товары</h3>
              <ul className="list-disc pl-5">
                {shipment.items.map((i) => (
                  <li key={i.id}>
                    product #{i.product_id} — {i.quantity} шт.
                  </li>
                ))}
              </ul>
            </div>

            {/* Редактирование */}
            {editingId === shipment.id ? (
              <div className="border-t pt-3 space-y-2">
                <h3 className="font-semibold">Изменить статус / курьера</h3>
                <div className="flex flex-col md:flex-row gap-2">
                  <select
                    value={editStatus}
                    onChange={(e) => setEditStatus(e.target.value)}
                    className="border p-2 rounded flex-1"
                  >
                    {statusOptions.map((s) => (
                      <option key={s} value={s}>{s}</option>
                    ))}
                  </select>
                  <select
                    value={editCourierId}
                    onChange={(e) => setEditCourierId(e.target.value)}
                    className="border p-2 rounded flex-1"
                  >
                    <option value="">Без курьера</option>
                    {couriers.map((c) => (
                      <option key={c.id} value={c.id}>
                        {c.first_name} {c.last_name}
                      </option>
                    ))}
                  </select>
                  <button
                    onClick={() => handleUpdate(shipment.tracking_code)}
                    className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700"
                  >
                    Сохранить
                  </button>
                  <button
                    onClick={cancelEdit}
                    className="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400"
                  >
                    Отмена
                  </button>
                </div>
              </div>
            ) : (
              <div className="border-t pt-3">
                <button
                  onClick={() => startEdit(shipment)}
                  className="bg-amber-600 text-white px-4 py-2 rounded hover:bg-amber-700"
                >
                  Редактировать
                </button>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}

