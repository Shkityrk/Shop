import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { api } from '../api/client';
import { useAuthStore } from '../store/useAuthStore';

interface Warehouse {
  id: number;
  name: string;
  address: string;
  working_hours?: string | null;
}

interface BinLocation {
  id: number;
  warehouse_id: number;
  zone: string;
  aisle: string;
  rack: string;
  bin_code: string;
  product_id?: number | null;
  quantity?: number | null;
}

interface Product {
  id: number;
  name: string;
  price: number;
}

export function AdminWarehousesPage() {
  const navigate = useNavigate();
  const { user, fetchUser, isLoading: authLoading } = useAuthStore();

  const [warehouses, setWarehouses] = useState<Warehouse[]>([]);
  const [bins, setBins] = useState<BinLocation[]>([]);
  const [products, setProducts] = useState<Product[]>([]);
  const [selectedWarehouseId, setSelectedWarehouseId] = useState<number | null>(null);
  const [whForm, setWhForm] = useState<Omit<Warehouse, 'id'>>({
    name: '',
    address: '',
    working_hours: '',
  });
  const [binForm, setBinForm] = useState({
    warehouse_id: 0,
    zone: '',
    aisle: '',
    rack: '',
    bin_code: '',
    product_id: '' as string | number,
    quantity: '' as string | number,
  });

  // Состояния для редактирования
  const [editingWarehouse, setEditingWarehouse] = useState<Warehouse | null>(null);
  const [editingBin, setEditingBin] = useState<BinLocation | null>(null);
  const [editWhForm, setEditWhForm] = useState<Omit<Warehouse, 'id'>>({
    name: '',
    address: '',
    working_hours: '',
  });
  const [editBinForm, setEditBinForm] = useState({
    zone: '',
    aisle: '',
    rack: '',
    bin_code: '',
    product_id: '' as string | number,
    quantity: '' as string | number,
  });

  // Загружаем данные пользователя при монтировании
  useEffect(() => {
    if (!user) {
      fetchUser();
    }
  }, [user, fetchUser]);

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

  const loadWarehouses = async () => {
    try {
      const resp = await api.get('/warehouse/warehouses');
      const data = resp.data;
      setWarehouses(Array.isArray(data) ? data : []);
    } catch (e) {
      console.error('Failed to load warehouses', e);
      setWarehouses([]);
    }
  };

  const loadBins = async (warehouseId?: number) => {
    try {
      const params = warehouseId ? { warehouse_id: warehouseId } : {};
      const resp = await api.get('/warehouse/locations/bins', { params });
      const data = resp.data;
      setBins(Array.isArray(data) ? data : []);
    } catch (e) {
      console.error('Failed to load bins', e);
      setBins([]);
    }
  };

  const loadProducts = async () => {
    try {
      const resp = await api.get('/product/list');
      const data = resp.data;
      setProducts(Array.isArray(data) ? data : []);
    } catch (e) {
      console.error('Failed to load products', e);
      setProducts([]);
    }
  };

  useEffect(() => {
    loadWarehouses();
    loadBins();
    loadProducts();
  }, []);

  const handleWhSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await api.post('/warehouse/warehouses', whForm);
    setWhForm({ name: '', address: '', working_hours: '' });
    await loadWarehouses();
  };

  const handleBinSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedWarehouseId) return;

    const payload: any = {
      warehouse_id: selectedWarehouseId,
      zone: binForm.zone,
      aisle: binForm.aisle,
      rack: binForm.rack,
      bin_code: binForm.bin_code,
    };

    // Добавляем product_id и quantity только если они заполнены
    if (binForm.product_id && binForm.quantity) {
      payload.product_id = Number(binForm.product_id);
      payload.quantity = Number(binForm.quantity);
    }

    await api.post('/warehouse/locations/bins', payload);
    setBinForm({
      warehouse_id: selectedWarehouseId,
      zone: '',
      aisle: '',
      rack: '',
      bin_code: '',
      product_id: '',
      quantity: '',
    });
    await loadBins(selectedWarehouseId);
  };

  const handleSelectWarehouse = async (id: number) => {
    setSelectedWarehouseId(id);
    await loadBins(id);
  };

  // === Удаление и редактирование складов ===
  const handleDeleteWarehouse = async (id: number) => {
    if (!confirm('Удалить склад? Все связанные полки тоже будут удалены.')) return;
    try {
      await api.delete(`/warehouse/warehouses/${id}`);
      if (selectedWarehouseId === id) {
        setSelectedWarehouseId(null);
        setBins([]);
      }
      await loadWarehouses();
    } catch (e) {
      console.error('Failed to delete warehouse', e);
      alert('Не удалось удалить склад');
    }
  };

  const startEditWarehouse = (wh: Warehouse) => {
    setEditingWarehouse(wh);
    setEditWhForm({
      name: wh.name,
      address: wh.address,
      working_hours: wh.working_hours ?? '',
    });
  };

  const handleUpdateWarehouse = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!editingWarehouse) return;
    try {
      await api.put(`/warehouse/warehouses/${editingWarehouse.id}`, editWhForm);
      setEditingWarehouse(null);
      await loadWarehouses();
    } catch (e) {
      console.error('Failed to update warehouse', e);
      alert('Не удалось обновить склад');
    }
  };

  // === Удаление и редактирование ячеек ===
  const handleDeleteBin = async (id: number) => {
    if (!confirm('Удалить ячейку?')) return;
    try {
      await api.delete(`/warehouse/locations/bins/${id}`);
      await loadBins(selectedWarehouseId ?? undefined);
    } catch (e) {
      console.error('Failed to delete bin', e);
      alert('Не удалось удалить ячейку');
    }
  };

  const startEditBin = (bin: BinLocation) => {
    setEditingBin(bin);
    setEditBinForm({
      zone: bin.zone,
      aisle: bin.aisle,
      rack: bin.rack,
      bin_code: bin.bin_code,
      product_id: bin.product_id ?? '',
      quantity: bin.quantity ?? '',
    });
  };

  const handleUpdateBin = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!editingBin) return;
    try {
      const payload: any = {
        zone: editBinForm.zone,
        aisle: editBinForm.aisle,
        rack: editBinForm.rack,
        bin_code: editBinForm.bin_code,
      };
      if (editBinForm.product_id && editBinForm.quantity) {
        payload.product_id = Number(editBinForm.product_id);
        payload.quantity = Number(editBinForm.quantity);
      } else {
        payload.product_id = null;
        payload.quantity = null;
      }
      await api.put(`/warehouse/locations/bins/${editingBin.id}`, payload);
      setEditingBin(null);
      await loadBins(selectedWarehouseId ?? undefined);
    } catch (e) {
      console.error('Failed to update bin', e);
      alert('Не удалось обновить ячейку');
    }
  };

  const getProductName = (productId: number | null | undefined) => {
    if (!productId) return '—';
    const product = products.find(p => p.id === productId);
    return product ? product.name : `#${productId}`;
  };

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold mb-4">Управление складами и полками</h1>

      <div className="grid md:grid-cols-2 gap-6">
        <form onSubmit={handleWhSubmit} className="bg-white p-4 rounded shadow space-y-2">
          <h2 className="font-semibold mb-2">Создать склад</h2>
          <input
            name="name"
            value={whForm.name}
            onChange={(e) => setWhForm((p) => ({ ...p, name: e.target.value }))}
            placeholder="Название"
            className="border p-2 rounded w-full"
            required
          />
          <input
            name="address"
            value={whForm.address}
            onChange={(e) => setWhForm((p) => ({ ...p, address: e.target.value }))}
            placeholder="Адрес"
            className="border p-2 rounded w-full"
            required
          />
          <input
            name="working_hours"
            value={whForm.working_hours ?? ''}
            onChange={(e) => setWhForm((p) => ({ ...p, working_hours: e.target.value }))}
            placeholder="Часы работы"
            className="border p-2 rounded w-full"
          />
          <button
            type="submit"
            className="bg-amber-600 text-white py-2 px-4 rounded hover:bg-amber-700"
          >
            Добавить склад
          </button>
        </form>

        <form onSubmit={handleBinSubmit} className="bg-white p-4 rounded shadow space-y-2">
          <h2 className="font-semibold mb-2">Создать полку с товаром</h2>
          <select
            value={selectedWarehouseId ?? ''}
            onChange={(e) => handleSelectWarehouse(Number(e.target.value))}
            className="border p-2 rounded w-full"
          >
            <option value="">Выберите склад</option>
            {warehouses.map((w) => (
              <option key={w.id} value={w.id}>
                {w.name}
              </option>
            ))}
          </select>
          <div className="grid grid-cols-2 gap-2">
            <input
              placeholder="Зона"
              value={binForm.zone}
              onChange={(e) => setBinForm((p) => ({ ...p, zone: e.target.value }))}
              className="border p-2 rounded"
              required
            />
            <input
              placeholder="Проход"
              value={binForm.aisle}
              onChange={(e) => setBinForm((p) => ({ ...p, aisle: e.target.value }))}
              className="border p-2 rounded"
              required
            />
            <input
              placeholder="Стеллаж"
              value={binForm.rack}
              onChange={(e) => setBinForm((p) => ({ ...p, rack: e.target.value }))}
              className="border p-2 rounded"
              required
            />
            <input
              placeholder="Код ячейки"
              value={binForm.bin_code}
              onChange={(e) => setBinForm((p) => ({ ...p, bin_code: e.target.value }))}
              className="border p-2 rounded"
              required
            />
          </div>

          <div className="border-t pt-3 mt-3">
            <h3 className="text-sm font-medium text-gray-700 mb-2">Товар на полке (опционально)</h3>
            <select
              value={binForm.product_id}
              onChange={(e) => setBinForm((p) => ({ ...p, product_id: e.target.value }))}
              className="border p-2 rounded w-full mb-2"
            >
              <option value="">Без товара</option>
              {products.map((p) => (
                <option key={p.id} value={p.id}>
                  {p.name} (#{p.id})
                </option>
              ))}
            </select>
            <input
              type="number"
              placeholder="Количество"
              value={binForm.quantity}
              onChange={(e) => setBinForm((p) => ({ ...p, quantity: e.target.value }))}
              className="border p-2 rounded w-full"
              min="1"
              disabled={!binForm.product_id}
            />
          </div>

          <button
            type="submit"
            className="bg-amber-600 text-white py-2 px-4 rounded hover:bg-amber-700 w-full"
            disabled={!selectedWarehouseId}
          >
            Добавить полку
          </button>
        </form>
      </div>

      <div className="grid md:grid-cols-2 gap-6">
        <div className="bg-white rounded shadow p-4">
          <h2 className="font-semibold mb-2">Склады</h2>
          <ul className="space-y-1 text-sm">
            {warehouses.map((w) => (
              <li
                key={w.id}
                className={`p-2 rounded ${selectedWarehouseId === w.id ? 'bg-amber-100' : 'hover:bg-gray-50'}`}
              >
                <div
                  className="cursor-pointer"
                  onClick={() => handleSelectWarehouse(w.id)}
                >
                  <div className="font-semibold">
                    #{w.id} {w.name}
                  </div>
                  <div className="text-gray-600">{w.address}</div>
                  {w.working_hours && (
                    <div className="text-gray-500 text-xs">Часы работы: {w.working_hours}</div>
                  )}
                </div>
                <div className="flex gap-2 mt-2">
                  <button
                    onClick={(e) => { e.stopPropagation(); startEditWarehouse(w); }}
                    className="text-xs bg-blue-500 text-white px-2 py-1 rounded hover:bg-blue-600"
                  >
                    ✏️ Изменить
                  </button>
                  <button
                    onClick={(e) => { e.stopPropagation(); handleDeleteWarehouse(w.id); }}
                    className="text-xs bg-red-500 text-white px-2 py-1 rounded hover:bg-red-600"
                  >
                    🗑️ Удалить
                  </button>
                </div>
              </li>
            ))}
          </ul>
        </div>

        <div className="bg-white rounded shadow p-4">
          <h2 className="font-semibold mb-2">
            Полки {selectedWarehouseId ? `(склад #${selectedWarehouseId})` : ''}
          </h2>
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b">
                <th className="text-left p-2">ID</th>
                <th className="text-left p-2">Локация</th>
                <th className="text-left p-2">Товар</th>
                <th className="text-left p-2">Кол-во</th>
                <th className="text-left p-2">Действия</th>
              </tr>
            </thead>
            <tbody>
              {bins.map((b) => (
                <tr key={b.id} className="border-b hover:bg-gray-50">
                  <td className="p-2">{b.id}</td>
                  <td className="p-2">
                    {b.zone} / {b.aisle} / {b.rack} / {b.bin_code}
                  </td>
                  <td className="p-2">
                    {b.product_id ? (
                      <span className="text-green-700 font-medium">
                        {getProductName(b.product_id)}
                      </span>
                    ) : (
                      <span className="text-gray-400">Пусто</span>
                    )}
                  </td>
                  <td className="p-2">
                    {b.quantity ? (
                      <span className="bg-amber-100 px-2 py-1 rounded text-amber-800">
                        {b.quantity} шт.
                      </span>
                    ) : (
                      '—'
                    )}
                  </td>
                  <td className="p-2">
                    <div className="flex gap-1">
                      <button
                        onClick={() => startEditBin(b)}
                        className="text-xs bg-blue-500 text-white px-2 py-1 rounded hover:bg-blue-600"
                      >
                        ✏️
                      </button>
                      <button
                        onClick={() => handleDeleteBin(b.id)}
                        className="text-xs bg-red-500 text-white px-2 py-1 rounded hover:bg-red-600"
                      >
                        🗑️
                      </button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>

      {/* Модальное окно редактирования склада */}
      {editingWarehouse && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white p-6 rounded-lg shadow-xl w-full max-w-md">
            <h2 className="text-xl font-bold mb-4">Редактировать склад #{editingWarehouse.id}</h2>
            <form onSubmit={handleUpdateWarehouse} className="space-y-3">
              <input
                value={editWhForm.name}
                onChange={(e) => setEditWhForm((p) => ({ ...p, name: e.target.value }))}
                placeholder="Название"
                className="border p-2 rounded w-full"
                required
              />
              <input
                value={editWhForm.address}
                onChange={(e) => setEditWhForm((p) => ({ ...p, address: e.target.value }))}
                placeholder="Адрес"
                className="border p-2 rounded w-full"
                required
              />
              <input
                value={editWhForm.working_hours ?? ''}
                onChange={(e) => setEditWhForm((p) => ({ ...p, working_hours: e.target.value }))}
                placeholder="Часы работы"
                className="border p-2 rounded w-full"
              />
              <div className="flex gap-2 justify-end">
                <button
                  type="button"
                  onClick={() => setEditingWarehouse(null)}
                  className="px-4 py-2 border rounded hover:bg-gray-100"
                >
                  Отмена
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-amber-600 text-white rounded hover:bg-amber-700"
                >
                  Сохранить
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Модальное окно редактирования ячейки */}
      {editingBin && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white p-6 rounded-lg shadow-xl w-full max-w-md">
            <h2 className="text-xl font-bold mb-4">Редактировать ячейку #{editingBin.id}</h2>
            <form onSubmit={handleUpdateBin} className="space-y-3">
              <div className="grid grid-cols-2 gap-2">
                <input
                  value={editBinForm.zone}
                  onChange={(e) => setEditBinForm((p) => ({ ...p, zone: e.target.value }))}
                  placeholder="Зона"
                  className="border p-2 rounded"
                  required
                />
                <input
                  value={editBinForm.aisle}
                  onChange={(e) => setEditBinForm((p) => ({ ...p, aisle: e.target.value }))}
                  placeholder="Проход"
                  className="border p-2 rounded"
                  required
                />
                <input
                  value={editBinForm.rack}
                  onChange={(e) => setEditBinForm((p) => ({ ...p, rack: e.target.value }))}
                  placeholder="Стеллаж"
                  className="border p-2 rounded"
                  required
                />
                <input
                  value={editBinForm.bin_code}
                  onChange={(e) => setEditBinForm((p) => ({ ...p, bin_code: e.target.value }))}
                  placeholder="Код ячейки"
                  className="border p-2 rounded"
                  required
                />
              </div>
              <div className="border-t pt-3">
                <h3 className="text-sm font-medium text-gray-700 mb-2">Товар на полке</h3>
                <select
                  value={editBinForm.product_id}
                  onChange={(e) => setEditBinForm((p) => ({ ...p, product_id: e.target.value }))}
                  className="border p-2 rounded w-full mb-2"
                >
                  <option value="">Без товара</option>
                  {products.map((p) => (
                    <option key={p.id} value={p.id}>
                      {p.name} (#{p.id})
                    </option>
                  ))}
                </select>
                <input
                  type="number"
                  placeholder="Количество"
                  value={editBinForm.quantity}
                  onChange={(e) => setEditBinForm((p) => ({ ...p, quantity: e.target.value }))}
                  className="border p-2 rounded w-full"
                  min="1"
                  disabled={!editBinForm.product_id}
                />
              </div>
              <div className="flex gap-2 justify-end">
                <button
                  type="button"
                  onClick={() => setEditingBin(null)}
                  className="px-4 py-2 border rounded hover:bg-gray-100"
                >
                  Отмена
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-amber-600 text-white rounded hover:bg-amber-700"
                >
                  Сохранить
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

