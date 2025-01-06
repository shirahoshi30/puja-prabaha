import React, { useState } from "react";
import { FaStar, FaArrowLeft, FaArrowRight } from "react-icons/fa";
import "./ProductDetail.css";
import Navbar from "../Navbar/Navbar";
import Footnote from "../FootNote/FootNote";
import ProductImage from "../../assets/products/product.png";

const ProductDetail = () => {
  const [activeTab, setActiveTab] = useState("reviews");
  const [quantity, setQuantity] = useState(1);
  const [currentIndex, setCurrentIndex] = useState(0);

  const reviews = [
    { name: "John Doe", rating: 5, feedback: "Great product, high quality and timely delivery!" },
    { name: "Jane Smith", rating: 4, feedback: "Good value for money but could improve packaging." },
    { name: "Emily Davis", rating: 5, feedback: "Absolutely loved it, works as advertised!" },
    { name: "Michael Brown", rating: 4, feedback: "Would recommend to friends and family!" },
  ];

  const renderStars = (rating) => {
    return [...Array(5)].map((_, index) => (
      <FaStar key={index} className={index < rating ? "star filled" : "star"} />
    ));
  };

  const handleQuantity = (type) => {
    if (type === "inc") {
      setQuantity(quantity + 1);
    } else if (type === "dec" && quantity > 1) {
      setQuantity(quantity - 1);
    }
  };

  const handleNext = () => {
    setCurrentIndex((prevIndex) =>
      prevIndex + 4 >= reviews.length ? 0 : prevIndex + 4
    );
  };

  const handlePrevious = () => {
    setCurrentIndex((prevIndex) =>
      prevIndex - 4 < 0 ? reviews.length - 4 : prevIndex - 4
    );
  };

  return (
    <>
      <Navbar />
      <div className="product-detail-page">
        {/* Breadcrumb */}
        <div className="breadcrumb">
          <a href="/" className="breadcrumb-link">Home</a> &gt; Shop &gt; Product
        </div>

        {/* Main Product Section */}
        <div className="product-section">
          <div className="product-image">
            <div className="product-images">
              <div className="thumbnail-images">
                <div className="thumbnail" />
                <div className="thumbnail" />
                <div className="thumbnail" />
              </div>
              <div className="main-image" />
            </div>
          </div>
          <div className="product-info">
            <h1 className="product-title">Lorem Ipsum dolor sit amet</h1>
            <div className="rating">
              <span className="stars">⭐⭐⭐⭐⭐</span>
              <span className="rating-score">4.5/5</span>
            </div>
            <div className="price">
              <span className="current-price">$260</span>
              <span className="old-price">$300</span>
              <span className="discount">-40%</span>
            </div>
            <p className="product-description">
              Lorem ipsum dolor sit amet consectetur. Gravida risus adipiscing mauris
              quam faucibus tortor nisl eget sodales. Sed congue sed lorem vitae enim
              tellus eget pulvinar.
            </p>

            {/* Quantity and Add to Cart */}
            <div className="cart-section">
              <div className="quantity">
                <button className="qty-btn" onClick={() => handleQuantity("dec")}>
                  -
                </button>
                <span className="qty-value">{quantity}</span>
                <button className="qty-btn" onClick={() => handleQuantity("inc")}>
                  +
                </button>
              </div>
              <button className="add-to-cart-btn">Add to Cart</button>
            </div>
          </div>
        </div>

        {/* Tabs Section */}
        <div className="tabs-section">
          <div className="tabs">
            <button onClick={() => setActiveTab("details")} className={activeTab === "details" ? "active" : ""}>Product Details</button>
            <button onClick={() => setActiveTab("reviews")} className={activeTab === "reviews" ? "active" : ""}>Ratings & Reviews</button>
            <button onClick={() => setActiveTab("faq")} className={activeTab === "faq" ? "active" : ""}>FAQ</button>
          </div>

          <div className="tab-content">
            {activeTab === "reviews" && (
              <div className="reviews-section">
                <h2>Customer Reviews</h2>
                <div className="reviews-container">
                  {/* Displaying 4 reviews at a time */}
                  {reviews.slice(currentIndex, currentIndex + 4).map((review, index) => (
                    <div className="review-card" key={index}>
                      <div className="review-rating">
                        {renderStars(review.rating)}
                      </div>
                      <p className="review-name">{review.name}</p>
                      <p className="review-feedback">{review.feedback}</p>
                    </div>
                  ))}
                </div>

                {/* Navigation Arrows */}
                <div className="navigation-buttons">
                  <button className="arrow-button" onClick={handlePrevious}>
                    <FaArrowLeft />
                  </button>
                  <button className="arrow-button" onClick={handleNext}>
                    <FaArrowRight />
                  </button>
                </div>
              </div>
            )}
            {activeTab === "details" && <p className="details-content">Lorem ipsum details content goes here.</p>}
            {activeTab === "faq" && <p className="faq-content">FAQ content goes here.</p>}
          </div>
        </div>

        {/* You Might Also Like Section */}
        <div className="suggestions-section">
          <h2 className="suggestions-heading">You May Also Like</h2>
          <div className="suggestions-container">
            {[1, 2, 3, 4].map((item) => (
              <div className="suggestion-card" key={item}>
                <img src={ProductImage} alt={`Product ${item}`} className="suggestion-img" />
                <div className="suggestion-info">
                  <h3 className="suggestion-title">Lorem Ipsum dolor sit amet</h3>
                  <div className="suggestion-rating">
                    {[...Array(5)].map((_, i) => (
                      <FaStar className="star" key={i} />
                    ))}
                  </div>
                  <p className="suggestion-price">$100.00</p>
                </div>
              </div>
            ))}
          </div>

          {/* Navigation Arrows */}
          <div className="suggestions-navigation">
            <button className="arrow-button" onClick={handlePrevious}>
              <FaArrowLeft />
            </button>
            <button className="arrow-button" onClick={handleNext}>
              <FaArrowRight />
            </button>
          </div>
        </div>
      </div>
      <Footnote />
    </>
  );
};

export default ProductDetail;

