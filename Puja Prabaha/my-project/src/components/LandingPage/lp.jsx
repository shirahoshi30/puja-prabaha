// src/components/LandingPage/lp.jsx
import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import videoBg from '../../assets/LandingPage/background.mp4'; 
import logo from '../../assets/logo/logo.png';
import './LandingPage.css';

const LandingPage = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const timer = setTimeout(() => {
      navigate('/home'); // Redirect to Home after 5 seconds
    }, 7000);

    // Cleanup timer on component unmount
    return () => clearTimeout(timer);
  }, [navigate]);

  // Handle click anywhere on the landing page to redirect immediately
  const handleClick = () => {
    navigate('/home');
  };

  return (
    <div className="landing-page" onClick={handleClick}>
      {/* Video Background */}
      <video className="video-background" autoPlay loop muted>
        <source src={videoBg} type="video/mp4" />
        Your browser does not support the video tag.
      </video>

      {/* Content Overlay */}
      <div className="content-overlay">
        <img src={logo} alt="Logo" className="logo" />
        <h1 className="slogan">Dharma is Karma</h1>
      </div>
    </div>
  );
};

export default LandingPage;
