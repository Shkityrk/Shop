import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { LogOut, User, Package, Settings, Edit2, ShoppingBag, Clock, MapPin } from 'lucide-react';
import { useAuthStore } from '../store/useAuthStore';
import { api } from '../api/client';

interface Order {
  id: number;
  created_at: string;
  status: string;
  total: number;
  items: Array<{
    product_name: string;
    quantity: number;
    price: number;
  }>;
}

interface UserProfile {
  username: string;
  email: string;
  first_name: string;
  last_name: string;
  address?: string;
  phone?: string;
  created_at: string;
}

export function ProfilePage() {
  const navigate = useNavigate();
  const { user, logout, isAuthenticated } = useAuthStore();
  const [activeTab, setActiveTab] = useState<'overview' | 'orders' | 'settings'>('overview');
  const [isEditing, setIsEditing] = useState(false);
  const [orders, setOrders] = useState<Order[]>([]);
  const [profile, setProfile] = useState<UserProfile | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [editForm, setEditForm] = useState({
    first_name: '',
    last_name: '',
    email: '',
    phone: '',
    address: ''
  });

  useEffect(() => {
    const fetchUserData = async () => {
      if (!isAuthenticated) {
        navigate('/login');
        return;
      }

      try {
        // Fetch user profile
        const profileResponse = await api.get('/auth/info');
        const profileData = profileResponse.data;
        setProfile(profileData);
        setEditForm({
          first_name: profileData.first_name || '',
          last_name: profileData.last_name || '',
          email: profileData.email || '',
          phone: profileData.phone || '',
          address: profileData.address || ''
        });

        // Fetch order history
        const ordersResponse = await api.get('/orders');
        const ordersData = Array.isArray(ordersResponse.data) ? ordersResponse.data : [];
        setOrders(ordersData);
      } catch (error) {
        console.error('Failed to fetch user data:', error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchUserData();
  }, [isAuthenticated, navigate]);

  const handleLogout = async () => {
    try {
      await logout();
      navigate('/');
    } catch (error) {
      console.error('Logout failed:', error);
    }
  };

  const handleSaveProfile = async () => {
    try {
      await api.put('/user/profile', editForm);
      setProfile(prev => ({ ...prev!, ...editForm }));
      setIsEditing(false);
    } catch (error) {
      console.error('Failed to update profile:', error);
    }
  };

  if (!isAuthenticated || !profile) {
    return null;
  }

  if (isLoading) {
    return (
      <div className="flex justify-center items-center min-h-[400px]">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-amber-600"></div>
      </div>
    );
  }

  return (
    <div className="max-w-6xl mx-auto">
      <div className="bg-white rounded-lg shadow-md overflow-hidden">
        {/* Profile Header */}
        <div className="bg-amber-50 p-6 border-b">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <div className="bg-amber-100 p-4 rounded-full">
                <User className="h-8 w-8 text-amber-600" />
              </div>
              <div>
                <h1 className="text-2xl font-bold text-amber-900">
                  {profile.first_name} {profile.last_name}
                </h1>
                <p className="text-gray-600">@{profile.username}</p>
              </div>
            </div>
            <button
              onClick={handleLogout}
              className="flex items-center px-4 py-2 text-red-600 hover:text-red-700 hover:bg-red-50 rounded-md transition-colors"
            >
              <LogOut className="h-5 w-5 mr-2" />
              Logout
            </button>
          </div>
        </div>

        {/* Navigation Tabs */}
        <div className="border-b">
          <nav className="flex">
            <button
              onClick={() => setActiveTab('overview')}
              className={`px-6 py-3 text-sm font-medium ${
                activeTab === 'overview'
                  ? 'border-b-2 border-amber-600 text-amber-600'
                  : 'text-gray-500 hover:text-amber-600'
              }`}
            >
              <User className="h-4 w-4 inline mr-2" />
              Overview
            </button>
            <button
              onClick={() => setActiveTab('orders')}
              className={`px-6 py-3 text-sm font-medium ${
                activeTab === 'orders'
                  ? 'border-b-2 border-amber-600 text-amber-600'
                  : 'text-gray-500 hover:text-amber-600'
              }`}
            >
              <Package className="h-4 w-4 inline mr-2" />
              Orders
            </button>
            <button
              onClick={() => setActiveTab('settings')}
              className={`px-6 py-3 text-sm font-medium ${
                activeTab === 'settings'
                  ? 'border-b-2 border-amber-600 text-amber-600'
                  : 'text-gray-500 hover:text-amber-600'
              }`}
            >
              <Settings className="h-4 w-4 inline mr-2" />
              Settings
            </button>
          </nav>
        </div>

        {/* Content */}
        <div className="p-6">
          {activeTab === 'overview' && (
            <div className="space-y-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div className="bg-amber-50 p-4 rounded-lg">
                  <div className="flex items-center justify-between mb-4">
                    <h3 className="text-lg font-semibold text-amber-900">Personal Information</h3>
                    <button
                      onClick={() => setIsEditing(!isEditing)}
                      className="text-amber-600 hover:text-amber-700"
                    >
                      <Edit2 className="h-4 w-4" />
                    </button>
                  </div>
                  {isEditing ? (
                    <div className="space-y-4">
                      <div>
                        <label className="block text-sm font-medium text-gray-700">First Name</label>
                        <input
                          type="text"
                          value={editForm.first_name}
                          onChange={(e) => setEditForm({ ...editForm, first_name: e.target.value })}
                          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-amber-500 focus:ring focus:ring-amber-200"
                        />
                      </div>
                      <div>
                        <label className="block text-sm font-medium text-gray-700">Last Name</label>
                        <input
                          type="text"
                          value={editForm.last_name}
                          onChange={(e) => setEditForm({ ...editForm, last_name: e.target.value })}
                          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-amber-500 focus:ring focus:ring-amber-200"
                        />
                      </div>
                      <div>
                        <label className="block text-sm font-medium text-gray-700">Email</label>
                        <input
                          type="email"
                          value={editForm.email}
                          onChange={(e) => setEditForm({ ...editForm, email: e.target.value })}
                          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-amber-500 focus:ring focus:ring-amber-200"
                        />
                      </div>
                      <div>
                        <label className="block text-sm font-medium text-gray-700">Phone</label>
                        <input
                          type="tel"
                          value={editForm.phone}
                          onChange={(e) => setEditForm({ ...editForm, phone: e.target.value })}
                          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-amber-500 focus:ring focus:ring-amber-200"
                        />
                      </div>
                      <div>
                        <label className="block text-sm font-medium text-gray-700">Address</label>
                        <textarea
                          value={editForm.address}
                          onChange={(e) => setEditForm({ ...editForm, address: e.target.value })}
                          rows={3}
                          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-amber-500 focus:ring focus:ring-amber-200"
                        />
                      </div>
                      <div className="flex justify-end space-x-2">
                        <button
                          onClick={() => setIsEditing(false)}
                          className="px-4 py-2 text-gray-600 hover:text-gray-700"
                        >
                          Cancel
                        </button>
                        <button
                          onClick={handleSaveProfile}
                          className="px-4 py-2 bg-amber-600 text-white rounded-md hover:bg-amber-700"
                        >
                          Save Changes
                        </button>
                      </div>
                    </div>
                  ) : (
                    <div className="space-y-3">
                      <p className="text-gray-600">
                        <span className="font-medium">Email:</span> {profile.email}
                      </p>
                      <p className="text-gray-600">
                        <span className="font-medium">Phone:</span> {profile.phone || 'Not provided'}
                      </p>
                      <p className="text-gray-600">
                        <span className="font-medium">Address:</span> {profile.address || 'Not provided'}
                      </p>
                      <p className="text-gray-600">
                        <span className="font-medium">Member since:</span>{' '}
                        {new Date(profile.created_at).toLocaleDateString()}
                      </p>
                    </div>
                  )}
                </div>

                <div className="bg-amber-50 p-4 rounded-lg">
                  <h3 className="text-lg font-semibold text-amber-900 mb-4">Recent Activity</h3>
                  <div className="space-y-4">
                    {orders.slice(0, 3).map((order) => (
                      <div key={order.id} className="flex items-center space-x-3">
                        <div className="bg-amber-100 p-2 rounded-full">
                          <ShoppingBag className="h-4 w-4 text-amber-600" />
                        </div>
                        <div>
                          <p className="text-sm font-medium text-amber-900">
                            Order #{order.id} - ${order.total.toFixed(2)}
                          </p>
                          <p className="text-xs text-gray-500">
                            {new Date(order.created_at).toLocaleDateString()}
                          </p>
                        </div>
                      </div>
                    ))}
                    {orders.length === 0 && (
                      <p className="text-gray-500">No recent orders</p>
                    )}
                  </div>
                </div>
              </div>
            </div>
          )}

          {activeTab === 'orders' && (
            <div className="space-y-6">
              <h3 className="text-xl font-semibold text-amber-900">Order History</h3>
              <div className="space-y-4">
                {orders.map((order) => (
                  <div key={order.id} className="bg-white border rounded-lg p-4">
                    <div className="flex items-center justify-between mb-4">
                      <div>
                        <p className="font-medium text-amber-900">Order #{order.id}</p>
                        <p className="text-sm text-gray-500">
                          {new Date(order.created_at).toLocaleDateString()}
                        </p>
                      </div>
                      <div className="text-right">
                        <p className="font-medium text-amber-600">${order.total.toFixed(2)}</p>
                        <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                          {order.status}
                        </span>
                      </div>
                    </div>
                    <div className="border-t pt-4">
                      {order.items && order.items.map((item, index) => (
                        <div key={index} className="flex justify-between text-sm">
                          <span>{item.product_name} x{item.quantity}</span>
                          <span>${(item.price * item.quantity).toFixed(2)}</span>
                        </div>
                      ))}
                    </div>
                  </div>
                ))}
                {orders.length === 0 && (
                  <p className="text-gray-500">No orders found</p>
                )}
              </div>
            </div>
          )}

          {activeTab === 'settings' && (
            <div className="max-w-2xl space-y-6">
              <h3 className="text-xl font-semibold text-amber-900">Account Settings</h3>
              <div className="space-y-4">
                <div className="bg-amber-50 p-4 rounded-lg">
                  <h4 className="font-medium text-amber-900 mb-2">Email Preferences</h4>
                  <div className="space-y-2">
                    <label className="flex items-center">
                      <input type="checkbox" className="rounded text-amber-600" />
                      <span className="ml-2 text-sm text-gray-700">Order confirmations</span>
                    </label>
                    <label className="flex items-center">
                      <input type="checkbox" className="rounded text-amber-600" />
                      <span className="ml-2 text-sm text-gray-700">Special offers and promotions</span>
                    </label>
                    <label className="flex items-center">
                      <input type="checkbox" className="rounded text-amber-600" />
                      <span className="ml-2 text-sm text-gray-700">Newsletter</span>
                    </label>
                  </div>
                </div>

                <div className="bg-amber-50 p-4 rounded-lg">
                  <h4 className="font-medium text-amber-900 mb-2">Password</h4>
                  <button className="text-amber-600 hover:text-amber-700 text-sm font-medium">
                    Change Password
                  </button>
                </div>

                <div className="bg-amber-50 p-4 rounded-lg">
                  <h4 className="font-medium text-amber-900 mb-2">Delete Account</h4>
                  <p className="text-sm text-gray-600 mb-2">
                    Once you delete your account, there is no going back. Please be certain.
                  </p>
                  <button className="text-red-600 hover:text-red-700 text-sm font-medium">
                    Delete Account
                  </button>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}