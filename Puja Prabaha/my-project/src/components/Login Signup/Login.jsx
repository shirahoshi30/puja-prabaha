import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./Login.css";

const LoginSignup = () => {
  const [isLogin, setIsLogin] = useState(true);

  const toggleForm = () => {
    setIsLogin(!isLogin);
  };

  return (
    <div
      className="login-signup-container"
      style={{
        backgroundSize: "cover",
        backgroundPosition: "center",
        minHeight: "100vh",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        padding: "20px",
      }}
    >
      <div
        className="form-container"
        style={{
          background: "white",
          padding: "40px",
          borderRadius: "10px",
          boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)",
          width: "100%",
          maxWidth: "400px",
        }}
      >
        <h2
          style={{
            fontFamily: "'Pirata One', cursive",
            color: "#C0353C",
            marginBottom: "20px",
            textAlign: "center",
          }}
        >
          {isLogin ? "Login to Puja Prabaha" : "Create an Account"}
        </h2>
        <form>
          {!isLogin && (
            <>
              <label className="form-label">Full Name</label>
              <input
                type="text"
                className="form-input"
                placeholder="Enter your full name"
              />
            </>
          )}
          <label className="form-label">Email Address</label>
          <input
            type="email"
            className="form-input"
            placeholder="Enter your email"
          />
          <label className="form-label">Password</label>
          <input
            type="password"
            className="form-input"
            placeholder="Enter your password"
          />
          {!isLogin && (
            <>
              <label className="form-label">Confirm Password</label>
              <input
                type="password"
                className="form-input"
                placeholder="Re-enter your password"
              />
            </>
          )}
          <button
            type="submit"
            className="form-button"
            style={{
              background: "#C0353C",
              color: "white",
              padding: "10px 20px",
              border: "none",
              borderRadius: "5px",
              width: "100%",
              marginTop: "20px",
              cursor: "pointer",
            }}
          >
            {isLogin ? "Login" : "Sign Up"}
          </button>
        </form>
        <div
          className="toggle-link"
          style={{
            textAlign: "center",
            marginTop: "20px",
            color: "#C0353C",
            cursor: "pointer",
          }}
          onClick={toggleForm}
        >
          {isLogin
            ? "Don't have an account? Sign up"
            : "Already have an account? Login"}
        </div>
        <Link
          to="/home"
          style={{
            display: "block",
            marginTop: "10px",
            textAlign: "center",
            textDecoration: "none",
            color: "#C0353C",
          }}
        >
          Back to Home
        </Link>
      </div>
    </div>
  );
};

export default LoginSignup;
