import React from "react";
import "./CelebrateTogetherHome.css";
import DashainImage from "../../assets/CelebrateTogether/blank.jpg";
import DipawaliImage from "../../assets/CelebrateTogether/blank.jpg";
import PurnimaImage from "../../assets/CelebrateTogether/blank.jpg";
import ShivaratriImage from "../../assets/CelebrateTogether/blank.jpg";

const CelebrateTogetherHome = () => {
  return (
    <div className="celebrate-together-section">
      <h2 className="celebrate-together-heading">Celebrate Together</h2>
      <p className="celebrate-together-description">
        Join our global South Asian family in cherishing timeless rituals and
        ceremonies. Let's keep the spirit of our traditions alive.
      </p>

      <div className="celebrate-cards">
        {/* Dashain Card */}
        <div className="celebrate-card">
          <img src={DashainImage} alt="Dashain" className="celebrate-image" />
          <div className="celebrate-card-content">
            <h3>Dashain</h3>
            <p>Celebrate victory and blessings with family and friends.</p>
            <button className="celebrate-learn-more">Learn More</button>
          </div>
        </div>

        {/* Dipawali Card */}
        <div className="celebrate-card">
          <img src={DipawaliImage} alt="Dipawali" className="celebrate-image" />
          <div className="celebrate-card-content">
            <h3>Dipawali</h3>
            <p>Light up your home with joy and prosperity.</p>
            <button className="celebrate-learn-more">Learn More</button>
          </div>
        </div>

        {/* Purnima Card */}
        <div className="celebrate-card">
          <img src={PurnimaImage} alt="Purnima" className="celebrate-image" />
          <div className="celebrate-card-content">
            <h3>Purnima</h3>
            <p>Auspicious full moon celebrations of peace and reflection.</p>
            <button className="celebrate-learn-more">Learn More</button>
          </div>
        </div>

        {/* Shivaratri Card */}
        <div className="celebrate-card">
          <img src={ShivaratriImage} alt="Shivaratri" className="celebrate-image" />
          <div className="celebrate-card-content">
            <h3>Shivaratri</h3>
            <p>Honor Lord Shiva with devotion and rituals.</p>
            <button className="celebrate-learn-more">Learn More</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CelebrateTogetherHome;