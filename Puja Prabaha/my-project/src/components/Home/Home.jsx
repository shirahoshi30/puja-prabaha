// src/components/Home/Home.jsx
import React from 'react';
import Navbar from '../Navbar/Navbar'; // Import Navbar
import Hero from '../Hero/Hero';
import CelebrateTogetherHome from '../CelebrateTogether/CelebrateTogetherHome';
import BestSellerHome from '../Best Seller/BestSellerHome';
import NewArrivalsHome from '../New Arrivals/NewArrivalsHome';
import CustomersHome from '../Customers/CustomersHome';
import StayConnected from '../StayConnected/StayConnected';
import Footnote from '../FootNote/FootNote';

const Home = () => {
  return (
    <div>
      <Navbar /> {/* Navbar will render here */}
      <Hero /> 
      <BestSellerHome /> 
      <NewArrivalsHome />
      <CelebrateTogetherHome />
      <CustomersHome />
      
      <Footnote />
    </div>
  );
};

export default Home;
