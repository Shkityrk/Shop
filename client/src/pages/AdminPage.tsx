import { Link } from 'react-router-dom';

interface AdminLink {
  to: string;
  title: string;
  description: string;
  icon: string;
}

const adminLinks: AdminLink[] = [
  {
    to: '/admin/products',
    title: 'Товары',
    description: 'Управление каталогом товаров: добавление, редактирование, удаление',
    icon: '📦',
  },
  {
    to: '/admin/warehouses',
    title: 'Склады',
    description: 'Управление складами, ячейками, правилами хранения и инвентарём',
    icon: '🏭',
  },
  {
    to: '/admin/shipments',
    title: 'Доставки',
    description: 'Просмотр и управление доставками, назначение курьеров',
    icon: '🚚',
  },
  {
    to: '/admin/staff',
    title: 'Сотрудники',
    description: 'Регистрация новых сотрудников: работники, склад, курьеры',
    icon: '👥',
  },
];

export function AdminPage() {
  return (
    <div className="space-y-6">
      <div className="text-center mb-8">
        <h1 className="text-3xl font-bold text-gray-800">Админ-панель</h1>
        <p className="text-gray-600 mt-2">Выберите раздел для управления</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {adminLinks.map((link) => (
          <Link
            key={link.to}
            to={link.to}
            className="bg-white p-6 rounded-lg shadow-md hover:shadow-lg transition-shadow border-l-4 border-amber-500 group"
          >
            <div className="flex items-start gap-4">
              <span className="text-4xl">{link.icon}</span>
              <div>
                <h2 className="text-xl font-semibold text-gray-800 group-hover:text-amber-600 transition-colors">
                  {link.title}
                </h2>
                <p className="text-gray-600 text-sm mt-1">{link.description}</p>
              </div>
            </div>
          </Link>
        ))}
      </div>

      <div className="mt-8 p-4 bg-amber-100 rounded-lg text-center">
        <p className="text-amber-800 text-sm">
          💡 Совет: Используйте навигацию выше для быстрого доступа к нужному разделу
        </p>
      </div>
    </div>
  );
}

