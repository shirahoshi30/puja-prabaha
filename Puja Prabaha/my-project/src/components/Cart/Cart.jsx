import React, { useState } from "react";
import { FaTrash } from "react-icons/fa";
import "./Cart.css";
import Navbar from "../Navbar/Navbar";
import Footnote from "../FootNote/FootNote";
import { Link } from "react-router-dom"; 
import ProductImage from "../../assets/products/product.png"; 

const CartPage = () => {
  const [cartItems, setCartItems] = useState([
    { id: 1, name: "Lorem ipsum dolor sit amet", price: 145, quantity: 1 },
    { id: 2, name: "Lorem ipsum dolor sit amet", price: 180, quantity: 1 },
    { id: 3, name: "Lorem ipsum dolor sit amet", price: 180, quantity: 1 },
  ]);

  // Calculate totals
  const subtotal = cartItems.reduce((total, item) => total + item.price * item.quantity, 0);
  const discount = subtotal * 0.2;
  const deliveryFee = 15;
  const total = subtotal - discount + deliveryFee;

  // Handle incrementing quantity
  const incrementQuantity = (id) => {
    setCartItems((prevItems) =>
      prevItems.map((item) =>
        item.id === id ? { ...item, quantity: item.quantity + 1 } : item
      )
    );
  };

  // Handle decrementing quantity
  const decrementQuantity = (id) => {
    setCartItems((prevItems) =>
      prevItems.map((item) =>
        item.id === id && item.quantity > 1 ? { ...item, quantity: item.quantity - 1 } : item
      )
    );
  };

  // Handle removing an item
  const removeItem = (id) => {
    setCartItems((prevItems) => prevItems.filter((item) => item.id !== id));
  };

  return (
    <>
      <Navbar />
      <div className="cart-page">
        <h1 className="cart-title">Your Cart</h1>
        <div className="cart-container">
          {/* Cart Items */}
          <div className="cart-items">
            {cartItems.map((item) => (
              <div key={item.id} className="cart-item">
                <div className="cart-item-info">
                  <img src={ProductImage} alt="Product" className="cart-item-image" />
                  <div>
                    <h3 className="cart-item-name">{item.name}</h3>
                    <p className="cart-item-price">${item.price}</p>
                  </div>
                </div>
                <div className="cart-item-controls">
                  <button className="quantity-btn" onClick={() => decrementQuantity(item.id)}>
                    -
                  </button>
                  <span className="quantity">{item.quantity}</span>
                  <button className="quantity-btn" onClick={() => incrementQuantity(item.id)}>
                    +
                  </button>
                  <FaTrash
                    className="delete-icon"
                    onClick={() => removeItem(item.id)}
                    title="Remove item"
                  />
                </div>
              </div>
            ))}
          </div>

          {/* Order Summary */}
          <div className="order-summary">
            <h2 className="summary-title">Order Summary</h2>
            <div className="summary-details">
              <p>
                <span>Subtotal</span>
                <span>${subtotal}</span>
              </p>
              <p>
                <span>Discount (-20%)</span>
                <span className="discount">-${discount.toFixed(2)}</span>
              </p>
              <p>
                <span>Delivery Fee</span>
                <span>${deliveryFee}</span>
              </p>
              <hr />
              <p className="total">
                <span>Total</span>
                <span>${total.toFixed(2)}</span>
              </p>
            </div>
            <div className="promo-code">
              <input type="text" placeholder="Add promo code" className="promo-input" />
              <button className="apply-btn">Apply</button>
            </div>
            <Link to = '/checkout'>
            <button className="checkout-btn">Go to Checkout →</button>
            </Link>
          </div>
        </div>
      </div>
      <Footnote />
    </>
  );
};

export default CartPage;

