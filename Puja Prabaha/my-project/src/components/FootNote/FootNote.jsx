import React from "react";
import "./FootNote.css"; // Ensure this CSS file contains updated code
import { FaEnvelope } from "react-icons/fa";
import StayDividerImage from "../../assets/BestSeller/dividerend.png"; // Divider image
import Logo from "../../assets/Logo/logo.png"; // Replace with actual logo path
import { SocialIcon } from 'react-social-icons'; // Import SocialIcon component
import Mastercard from 'react-payment-logos/dist/flat/Mastercard';
import Visa from 'react-payment-logos/dist/flat/Visa';
import Paypal from 'react-payment-logos/dist/flat/Paypal';

const Footnote = () => {
  return (
    <div>
      {/* Space before Stay Connected Section */}
      <div className="pre-stay-connected-space"></div>

      {/* Stay Connected Box */}
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

      {/* Divider Below Stay Connected */}
      <div className="divider-image">
        <img src={StayDividerImage} alt="End Divider" />
      </div>

      {/* Footnote Section */}
      <div className="footnote-container">
        <div className="footer-section">
          <div className="about">
            {/* Logo Above Explore Section */}
            <img src={Logo} alt="Logo" className="about-logo" />
            <p className="about-text">
              Explore premium puja essentials and connect with your spiritual roots. Trusted by South Asian families worldwide.
            </p>

            <p className="follow-us-text">
              Follow us on social media for daily spiritual tips and updates!
            </p>
            {/* Social Media Icons */}
            <div className="social-icons">
              <SocialIcon url="https://www.facebook.com" target="_blank" style={{ marginRight: '15px' }} />
              <SocialIcon url="https://www.instagram.com" target="_blank" style={{ marginRight: '15px' }} />
              <SocialIcon url="https://www.twitter.com" target="_blank" style={{ marginRight: '15px' }} />
            </div>
          </div>

          <div className="links">
            <h4>Company</h4>
            <ul>
              <li>About Us</li>
              <li>Products</li>
              <li>Our Mission</li>
              <li>Happy Customers</li>
            </ul>
          </div>
          <div className="help">
            <h4>Help</h4>
            <ul>
              <li>Customer Support</li>
              <li>Delivery Details</li>
              <li>Terms & Conditions</li>
              <li>Privacy Policy</li>
            </ul>
          </div>
          <div className="faq">
            <h4>FAQ</h4>
            <ul>
              <li>Account</li>
              <li>Manage Deliveries</li>
              <li>Orders</li>
              <li>Payments</li>
            </ul>
          </div>
        </div>

         {/* Payment Methods Section */}
         <div className="payment-methods">
            <div className="payment-icons">
              <Mastercard style={{ marginRight: '15px', width: '40px', height: '40px' }} />
              <Visa style={{ marginRight: '15px', width: '40px', height: '40px' }} />
              <Paypal style={{ width: '40px', height: '40px' }} />
            </div>
          </div>

        <div className="footer-bottom">
          <p>&copy; 2023 Puja Pravah. All Rights Reserved.</p>
        </div>
      </div>
    </div>
  );
};

export default Footnote;
