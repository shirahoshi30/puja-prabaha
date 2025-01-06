import React, { useState } from "react";
import { useNavigate } from "react-router-dom"; // Import useNavigate hook
import "./Hero.css";
import HeroImage2 from "../../assets/Hero/Hero 1.jpg";  // Import image
import HeroImage3 from "../../assets/Hero/Hero.jpg";   // Import image
;
const Hero = () => {
  const [backgroundImage, setBackgroundImage] = useState(HeroImage2);  // Default background

  const backgroundImages = [HeroImage2, HeroImage3];  // Array of imported images

  const navigate = useNavigate(); // Initialize the navigate function

  const changeBackground = () => {
    const currentIndex = backgroundImages.indexOf(backgroundImage);
    const nextIndex = (currentIndex + 1) % backgroundImages.length;
    setBackgroundImage(backgroundImages[nextIndex]);
  };

  const handleButtonClick = () => {
    navigate("/Shop"); // Redirect to the shop page
  };

  return (
    <div className="hero" style={{ backgroundImage: `url(${backgroundImage})` }} onClick={changeBackground}>
      <div className="hero-overlay"></div>
      <div className="hero-text">
        <h1>
          Bring the Essence of <br />
          Tradition Home
        </h1>
        <p>
          At Puja Prabaha, we bring spirituality closer to you. <br />
          Guided by <strong>'Dharma is Karma,'</strong> we offer authentic puja <br />
          items to celebrate traditions and preserve South Asian <br /> heritage.
        </p>
        <button onClick={handleButtonClick}>Discover Puja Essentials</button>
      </div>
    </div>
  );
};

export default Hero;
