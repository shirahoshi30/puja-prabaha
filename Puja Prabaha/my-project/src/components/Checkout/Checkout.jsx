import React, { useState } from "react";
import Navbar from "../Navbar/Navbar"; // Optional
import Footnote from "../FootNote/FootNote";
import "./Checkout.css";

const CheckoutPage = () => {
  const [formData, setFormData] = useState({
    fullName: "",
    email: "",
    address: "",
    city: "",
    state: "",
    zip: "",
    cardNumber: "",
    expiryDate: "",
    cvv: "",
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    alert("Order Placed Successfully!");
  };

  return (
    <>
      <Navbar />
      <div className="checkout-page">
        <h1 className="checkout-title">Checkout</h1>
        <div className="checkout-container">
          {/* Billing Details */}
          <form className="checkout-form" onSubmit={handleSubmit}>
            <h2 className="form-section-title">Billing Details</h2>
            <div className="form-group">
              <label>Full Name</label>
              <input
                type="text"
                name="fullName"
                value={formData.fullName}
                onChange={handleChange}
                required
              />
            </div>
            <div className="form-group">
              <label>Email</label>
              <input
                type="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                required
              />
            </div>
            <div className="form-group">
              <label>Address</label>
              <input
                type="text"
                name="address"
                value={formData.address}
                onChange={handleChange}
                required
              />
            </div>
            <div className="form-row">
              <div className="form-group">
                <label>City</label>
                <input
                  type="text"
                  name="city"
                  value={formData.city}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="form-group">
                <label>State</label>
                <input
                  type="text"
                  name="state"
                  value={formData.state}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="form-group">
                <label>ZIP</label>
                <input
                  type="text"
                  name="zip"
                  value={formData.zip}
                  onChange={handleChange}
                  required
                />
              </div>
            </div>

            {/* Payment Details */}
            <h2 className="form-section-title">Payment Details</h2>
            <div className="form-group">
              <label>Card Number</label>
              <input
                type="text"
                name="cardNumber"
                value={formData.cardNumber}
                onChange={handleChange}
                required
              />
            </div>
            <div className="form-row">
              <div className="form-group">
                <label>Expiry Date</label>
                <input
                  type="text"
                  name="expiryDate"
                  value={formData.expiryDate}
                  onChange={handleChange}
                  placeholder="MM/YY"
                  required
                />
              </div>
              <div className="form-group">
                <label>CVV</label>
                <input
                  type="text"
                  name="cvv"
                  value={formData.cvv}
                  onChange={handleChange}
                  required
                />
              </div>
            </div>

            <button type="submit" className="place-order-btn">
              Place Order
            </button>
          </form>

          {/* Order Summary */}
          <div className="order-summary">
            <h2 className="summary-title">Order Summary</h2>
            <div className="summary-details">
              <p>
                <span>Subtotal</span>
                <span>$467</span>
              </p>
              <p>
                <span>Discount (-20%)</span>
                <span className="discount">-$93</span>
              </p>
              <p>
                <span>Delivery Fee</span>
                <span>$15</span>
              </p>
              <hr />
              <p className="total">
                <span>Total</span>
                <span>$389</span>
              </p>
            </div>
          </div>
        </div>
      </div>
      <Footnote />
    </>
  );
};

export default CheckoutPage;

