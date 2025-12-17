export interface Product {
  id: number;
  name: string;
  short_description: string;
  full_description: string;
  price: number;
  weight: number;
  composition: string;
  photo: string;
  storage_rule_id?: number | null;
}

export interface StorageRule {
  id: number;
  name: string;
  description?: string | null;
  is_hazardous: boolean;
  is_oversized: boolean;
  temp_min?: number | null;
  temp_max?: number | null;
}

export interface User {
  id: number;
  username: string;
  email: string;
  first_name: string;
  last_name: string;
  user_role?: 'client' | 'staff' | 'warehouse' | 'courier' | 'admin';
}

export interface CartItem {
  id?: number;
  user_id?: number;
  product_id: number;
  quantity: number;
  product?: Product | null;
}