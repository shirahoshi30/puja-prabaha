import React from "react";
import { Link } from "react-router-dom";
import Logo from "../../assets/logo/logo.png"; 
import NavbarBg from "../../assets/navbar/navbar.png"; 
import { IoCartOutline } from "react-icons/io5";
import { IoMdSearch } from "react-icons/io";
import { FaUserCircle } from "react-icons/fa";
import './Navbar.css';

const Navbar = () => {
  return (
    <header
      style={{
        backgroundImage: `url(${NavbarBg})`, // Adding the background image
        backgroundSize: "cover",
        backgroundPosition: "center",
        backgroundRepeat: "no-repeat",
        width: "100%",
        height: "100px", // Adjust height to match design
        borderBottom: "1px solid #d3c5b8", // Subtle border at the bottom
      }}
    >
      <nav
        className="container mx-auto flex items-center justify-between h-full px-8"
        style={{ display: "flex", alignItems: "center" }}
      >
        {/* Left Section: Logo */}
        <div className="flex items-center space-x-4">
          <Link to="/" className="w-auto h-auto">
            <img
              src={Logo}
              alt="Puja Prabaha Logo"
              className="w-auto h-auto"
              style={{ width: "150px", height: "auto" }} // Set logo to 200px width (reverting back to original size)
            />
          </Link>
        </div>

        {/* Center Section: Navigation Links */}
        <div className="hidden lg:flex space-x-10 text-lg font-semibold text-gray-800">
          <Link
            to="/home"
            className="hover:text-red-500 transition-colors pirata-one-regular"
            style={{ color: "#C0353C" }} // Red text color
          >
            Home
          </Link>
          <Link
            to="/aboutus"
            className="hover:text-red-500 transition-colors pirata-one-regular"
            style={{ color: "#C0353C" }} // Red text color
          >
            About Us
          </Link>
          <div className="relative group">
            <Link
              to="/shop"
              className="hover:text-red-500 transition-colors pirata-one-regular"
              style={{ color: "#C0353C" }} // Red text color
            >
              <span>Shop</span>
              <span>▼</span>
            </Link>
            {/* Dropdown menu */}
            <div
              className="absolute hidden group-hover:block top-8 right-0 bg-white shadow-md rounded-md p-2"
              style={{
                width: "200px", // Adjusted width to fit both items
              }}
            >
              <Link
                to="/newarrivals"
                className="block px-4 py-2 hover:bg-gray-100 pirata-one-regular"
                style={{ color: "#C0353C" }}
              >
                New Arrivals
              </Link>
              <Link
                to="/bestseller"
                className="block px-4 py-2 hover:bg-gray-100 pirata-one-regular"
                style={{ color: "#C0353C" }}
              >
                Best Sellers
              </Link>
            </div>
          </div>
        </div>

        {/* Center Section: Search Bar */}
        <div
          className="relative flex-grow max-w-md mx-10 hidden lg:flex"
          style={{ display: "flex", alignItems: "center", flexGrow: 1 }}
        >
          <IoMdSearch
            className="absolute left-4 text-xl text-gray-400"
            style={{ position: "absolute", left: "10px" }}
          />
          <input
            type="text"
            placeholder="Search for products..."
            className="rounded-full border border-gray-300 px-12 py-2"
            style={{
              width: "100%",
              height: "40px",
              paddingLeft: "40px",
              paddingRight: "10px",
            }}
          />
        </div>

        {/* Right Section: Icons */}
        <div className="flex items-center space-x-6 text-xl text-gray-700">

        <Link to="/cart">
          <IoCartOutline className="text-[#C0353C] hover:text-black cursor-pointer transition-colors" />
        </Link>

          <Link to="/login-signup">
           <FaUserCircle className="text-[#C0353C] hover:text-black cursor-pointer transition-colors" />
          </Link>

          <div className="relative group">
            <button className="flex items-center text-[#C0353C] hover:text-black transition-colors pirata-one-regular">
              Language
              <span className="ml-2">▼</span>
            </button>
            {/* Dropdown menu */}
            <div className="absolute hidden group-hover:block top-8 right-0 bg-white shadow-md rounded-md p-2">
              <button className="block px-4 py-2 hover:bg-gray-100 pirata-one-regular" style={{ color: "#C0353C" }}>English</button>
              <button className="block px-4 py-2 hover:bg-gray-100 pirata-one-regular" style={{ color: "#C0353C" }}>Nepali</button>
              <button className="block px-4 py-2 hover:bg-gray-100 pirata-one-regular" style={{ color: "#C0353C" }}>Hindi</button>
            </div>
          </div>
        </div>
      </nav>
    </header>
  );
};

export default Navbar;

