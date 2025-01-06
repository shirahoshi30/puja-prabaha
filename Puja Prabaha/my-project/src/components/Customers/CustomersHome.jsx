import React, { useState } from "react";
import "./CustomersHome.css";
import { FaStar, FaArrowLeft, FaArrowRight } from "react-icons/fa";

const reviews = [
  { name: "John Doe", rating: 5, feedback: "Great experience! Highly recommended." },
  { name: "Jane Smith", rating: 4, feedback: "Good service, but delivery was slightly delayed." },
  { name: "Sam Wilson", rating: 5, feedback: "Amazing products and excellent customer support!" },
  { name: "Emily Davis", rating: 3, feedback: "Products were fine but packaging needs improvement." },
  { name: "Michael Brown", rating: 5, feedback: "Exceptional quality! Will buy again." },
  { name: "Sarah Johnson", rating: 4, feedback: "Good products overall, but a bit expensive." },
];

const CustomersHome = () => {
  const [currentIndex, setCurrentIndex] = useState(0);

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
    <div className="customers-home-container">
      <h2 className="customers-heading">What Our Customers Say</h2>

      <div className="reviews-container">
        {/* Displaying 4 reviews at a time */}
        {reviews.slice(currentIndex, currentIndex + 4).map((review, index) => (
          <div className="review-card" key={index}>
            <div className="review-rating">
              {[...Array(review.rating)].map((_, i) => (
                <FaStar className="star" key={i} />
              ))}
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
  );
};

export default CustomersHome;
