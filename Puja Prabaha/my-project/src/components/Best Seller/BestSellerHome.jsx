import React from "react";
import "./BestSellerHome.css";
import BestSellerBar from "../../assets/BestSeller/bestseller.jpg";
import DividerImage from "../../assets/BestSeller/divider.png";
import EndDividerImage from "../../assets/BestSeller/dividerend.png";
import SectionBackground from "../../assets/BestSeller/background.jpg";
import ProductImage from "../../assets/products/product.png";
import { useNavigate } from "react-router-dom";

const BestSellerHome = () => {

  const navigate = useNavigate(); // Initialize the useNavigate hook

  // Function to navigate to the Best Seller page
  const handleButtonClick = () => {
    navigate("/bestseller"); // Change "/bestseller" to match your route path
  };
  
  return (
    <div>
      {/* Best Seller Bar */}
      <div className="best-seller-bar">
        <h2 className="best-seller-heading">Our Best Sellers</h2>
      </div>

      {/* Section Content with Background */}
      <div
        className="section-background"
        style={{ backgroundImage: `url(${SectionBackground})` }}
      >
        {/* Positioned Divider */}
        <img src={DividerImage} alt="Divider" className="divider-on-background" />

        <div className="section-content">
          <p className="section-description">
            Explore premium puja essentials and connect with your spiritual roots. Trusted by South Asian families worldwide.
          </p>

          {/* Products Grid - Modified to 4x Grid */}
          <div className="products-grid">
            <div className="product-card">
              <img src={ProductImage} alt="Handcrafted Idols" className="product-image" />
              <p className="product-label">Handcrafted Idols</p>
            </div>
            <div className="product-card">
              <img src={ProductImage} alt="Authentic Puja Kits" className="product-image" />
              <p className="product-label">Authentic Puja Kits</p>
            </div>
            <div className="product-card">
              <img src={ProductImage} alt="Organic Incense and Oils" className="product-image" />
              <p className="product-label">Organic Incense and Oils</p>
            </div>
            <div className="product-card">
              <img src={ProductImage} alt="Sacred Scriptures and Ritual Guides" className="product-image" />
              <p className="product-label">Sacred Scriptures and Ritual Guides</p>
            </div>
          </div>

          {/* Centered Button */}
          <button className="center-button" onClick={handleButtonClick}>
            View All Products
          </button>
        </div>
      </div>

      {/* End Divider */}
      <img src={EndDividerImage} alt="End Divider" className="divider-end-image" />
    </div>
  );
};

export default BestSellerHome;


