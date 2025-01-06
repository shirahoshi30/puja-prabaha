import React from "react";
import "./StayConnected.css";
import { FaEnvelope } from "react-icons/fa";

const StayConnected = () => {
  return (
    <div className="stay-connected-wrapper">
      <div className="stay-connected-box">
        <div className="stay-connected-content">
          <div className="stay-connected-left">
            <h2 className="stay-connected-heading">Stay Connected</h2>
            <p className="stay-connected-text">
              Subscribe to get updates on our latest products, offers, and tips for your next puja.
            </p>
          </div>
          <div className="stay-connected-right">
            <div className="subscription-form">
              <div className="email-input-container">
                <FaEnvelope className="email-icon" />
                <input
                  type="email"
                  placeholder="Enter your email address"
                  className="email-input"
                />
              </div>
              <button className="subscribe-button">Subscribe Now</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default StayConnected;
