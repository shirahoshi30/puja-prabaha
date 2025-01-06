// src/App.jsx
import React from 'react';
import { Routes, Route, BrowserRouter } from 'react-router-dom';
import LandingPage from './components/LandingPage/lp'; // Import Landing Page
import Home from './components/Home/Home'; // Import Home Page
import Shop from './components/Shop/Shop'; // Import Shop Component
import BestSeller from './components/Best Seller/BestSeller';
import NewArrivals from './components/New Arrivals/NewArrivals';
import ProductDetail from './components/Product Details/ProductDetail';
import Login from './components/Login Signup/Login';
import Cart from './components/Cart/Cart';
import Checkout from './components/Checkout/Checkout';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<LandingPage />} /> {/* Landing Page */}
        <Route path="/home" element={<Home />} /> {/* Home Page with Navbar */}
        <Route path="/shop" element={<Shop />} /> {/* Shop Page */}
        <Route path="/bestseller" element={<BestSeller />} /> {/* BestSeller route */}
        <Route path="/newarrivals" element={<NewArrivals />} /> 
        <Route path="/product/:id" element={<ProductDetail />} />
        <Route path="/login-signup" element={<Login />} /> 
        <Route path="/cart" element={<Cart />} /> 
        <Route path="/checkout" element={<Checkout />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;



