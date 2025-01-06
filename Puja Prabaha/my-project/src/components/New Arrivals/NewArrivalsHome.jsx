import React from "react";
import "./NewArrivalsHome.css";
import ProductImage from "../../assets/products/product.png";
import { FaStar, FaRegStar } from "react-icons/fa";  // Import FontAwesome icons for stars
import { useNavigate } from "react-router-dom";

const NewArrivalsHome = () => {

  const navigate = useNavigate(); // Initialize the useNavigate hook

  // Function to navigate to the Best Seller page
  const handleButtonClick = () => {
    navigate("/newarrivals"); // Change "/bestseller" to match your route path
  };

  return (
    <div className="new-arrivals-container">
      {/* Heading */}
      <h2 className="new-arrivals-heading">New Arrivals</h2>

      {/* Section Content */}
      <div className="new-arrivals-content">
        {/* Products Grid */}
        <div className="products-grid">
          {/* Product 1 */}
          <div className="product-card">
            <img src={ProductImage} alt="Product 1" className="product-image" />
            <p className="product-name">Product 1</p>

            {/* Stars and Price/Discount */}
            <div className="product-rating">
              <FaStar className="star" />
              <FaStar className="star" />
              <FaStar className="star" />
              <FaStar className="star" />
              <FaRegStar className="star" />
            </div>

            <div className="product-price-discount">
              <p className="product-price">$20.00</p>
              <p className="product-discount">$25.00</p>
              <div className="discount-percent">20% OFF</div>
            </div>
          </div>

          {/* Product 2 */}
          <div className="product-card">
            <img src={ProductImage} alt="Product 2" className="product-image" />
            <p className="product-name">Product 2</p>

            {/* Stars and Price/Discount */}
            <div className="product-rating">
              <FaStar className="star" />
              <FaStar className="star" />
              <FaStar className="star" />
              <FaRegStar className="star" />
              <FaRegStar className="star" />
            </div>

            <div className="product-price-discount">
              <p className="product-price">$15.00</p>
              <p className="product-discount">$18.00</p>
              <div className="discount-percent">16% OFF</div>
            </div>
          </div>

          {/* Product 3 */}
          <div className="product-card">
            <img src={ProductImage} alt="Product 3" className="product-image" />
            <p className="product-name">Product 3</p>

            {/* Stars and Price/Discount */}
            <div className="product-rating">
              <FaStar className="star" />
              <FaStar className="star" />
              <FaStar className="star" />
              <FaStar className="star" />
              <FaStar className="star" />
            </div>

            <div className="product-price-discount">
              <p className="product-price">$18.00</p>
              <p className="product-discount">$22.00</p>
              <div className="discount-percent">18% OFF</div>
            </div>
          </div>

          {/* Product 4 - New Product */}
          <div className="product-card">
            <img src={ProductImage} alt="Product 4" className="product-image" />
            <p className="product-name">Product 4</p>

            {/* Stars and Price/Discount */}
            <div className="product-rating">
              <FaStar className="star" />
              <FaStar className="star" />
              <FaStar className="star" />
              <FaRegStar className="star" />
              <FaRegStar className="star" />
            </div>

            <div className="product-price-discount">
              <p className="product-price">$25.00</p>
              <p className="product-discount">$30.00</p>
              <div className="discount-percent">17% OFF</div>
            </div>
          </div>
        </div>
      </div>

      {/* View All Products Button */}
      <button className="view-all-button" onClick={handleButtonClick}>
        View All Products
      </button>
    </div>
  );
};

export default NewArrivalsHome;


