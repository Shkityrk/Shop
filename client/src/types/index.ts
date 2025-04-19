export interface Product {
  id: number;
  name: string;
  short_description: string;
  full_description: string;
  price: number;
  weight: number;
  composition: string;
  photo: string;
}

export interface User {
  username: string;
  email: string;
  first_name: string;
  last_name: string;
}

export interface CartItem {
  productId: string;
  quantity: number;
}