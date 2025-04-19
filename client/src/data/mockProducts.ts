import { Product } from '../types';

export const mockProducts: Product[] = [
  {
    id: '1',
    name: 'Classic Croissant',
    description: 'Buttery, flaky croissant made with premium French butter. Each layer is carefully folded to create a perfect, airy texture.',
    price: 3.99,
    imageUrl: 'https://images.unsplash.com/photo-1555507036-ab1f4038808a?w=800',
    category: 'Pastries'
  },
  {
    id: '2',
    name: 'Sourdough Bread',
    description: 'Traditional sourdough bread made with our 100-year-old starter. Crusty exterior with a soft, tangy interior.',
    price: 6.99,
    imageUrl: 'https://images.unsplash.com/photo-1585478259715-4d3f8f7334ed?w=800',
    category: 'Bread'
  },
  {
    id: '3',
    name: 'Chocolate Danish',
    description: 'Flaky Danish pastry filled with rich chocolate and topped with a chocolate drizzle.',
    price: 4.49,
    imageUrl: 'https://images.unsplash.com/photo-1509440159596-0249088772ff?w=800',
    category: 'Pastries'
  },
  {
    id: '4',
    name: 'Baguette',
    description: 'Traditional French baguette with a crispy crust and soft interior. Baked fresh daily.',
    price: 3.49,
    imageUrl: 'https://images.unsplash.com/photo-1523471826770-c437b4636fe6?w=800',
    category: 'Bread'
  },
  {
    id: '5',
    name: 'Blueberry Muffin',
    description: 'Moist muffin packed with fresh blueberries and topped with a crumb streusel.',
    price: 3.99,
    imageUrl: 'https://images.unsplash.com/photo-1607958996333-41aef7caefaa?w=800',
    category: 'Muffins'
  },
  {
    id: '6',
    name: 'Cinnamon Roll',
    description: 'Soft, swirled roll filled with cinnamon-sugar and topped with cream cheese frosting.',
    price: 4.99,
    imageUrl: 'https://images.unsplash.com/photo-1509365465985-25d11c17e812?w=800',
    category: 'Pastries'
  }
];