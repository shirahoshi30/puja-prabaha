import React, { useState } from "react";
import "./BestSeller.css";
import Navbar from "../Navbar/Navbar"; 
import ProductImage from "../../assets/products/product.png";
import { Link } from "react-router-dom"; 
import Footnote from "../FootNote/FootNote";

const BestSeller = () => {
  // Set the initial state for pagination
  const [currentPage, setCurrentPage] = useState(1);
  const productsPerPage = 10;

  // Generate a list of 40 products (for example)
  const allProducts = [...Array(40)].map((_, index) => ({
    id: index + 1,
    name: `Product ${index + 1}`,
    price: 50,
  }));

  const totalPages = Math.ceil(allProducts.length / productsPerPage); // Calculate total pages

  // Calculate which products to show based on the current page
  const indexOfLastProduct = currentPage * productsPerPage;
  const indexOfFirstProduct = indexOfLastProduct - productsPerPage;
  const currentProducts = allProducts.slice(indexOfFirstProduct, indexOfLastProduct);

  // Handle next and previous page navigation
  const handleNextPage = () => {
    if (currentPage < totalPages) {
      setCurrentPage(currentPage + 1);
    }
  };

  const handlePrevPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  const renderPagination = () => {
    const paginationButtons = [];

    // Always show the first page
    paginationButtons.push(
      <button
        key={1}
        className={`page-btn ${currentPage === 1 ? "active" : ""}`}
        onClick={() => setCurrentPage(1)}
      >
        1
      </button>
    );

    // If the current page is greater than 3, show the dots
    if (currentPage > 3) {
      paginationButtons.push(<span key="dots-left" className="dots">...</span>);
    }

    // Show the current page, one before and one after it
    for (let i = Math.max(2, currentPage - 1); i <= Math.min(totalPages - 1, currentPage + 1); i++) {
      paginationButtons.push(
        <button
          key={i}
          className={`page-btn ${currentPage === i ? "active" : ""}`}
          onClick={() => setCurrentPage(i)}
        >
          {i}
        </button>
      );
    }

    // If the current page is less than totalPages - 2, show the dots
    if (currentPage < totalPages - 2) {
      paginationButtons.push(<span key="dots-right" className="dots">...</span>);
    }

    // Always show the last page
    if (totalPages > 1) {
      paginationButtons.push(
        <button
          key={totalPages}
          className={`page-btn ${currentPage === totalPages ? "active" : ""}`}
          onClick={() => setCurrentPage(totalPages)}
        >
          {totalPages}
        </button>
      );
    }

    return paginationButtons;
  };

  return (
    <>
      <Navbar /> {/* Navbar outside of shop-container */}
      <div className="shop-container">
        <div className="breadcrumb">
          <Link to="/" className="breadcrumb-link">
            Home
          </Link>{" "}
          <Link to="/Shop" className="breadcrumb-link">
          {`>`} Shop
          </Link>
          {`>`} Best Seller
        </div>
        <div className="shop-content">
          {/* Sidebar */}
          <aside className="shop-sidebar">
            <h3>Lorem Ipsum Dolor</h3>
            <ul>
              <li>Lorem ipsum</li>
              <li>Lorem ipsum</li>
              <li>Lorem ipsum</li>
              <li>Lorem ipsum</li>
            </ul>

            <div className="price-filter">
              <h4>Price</h4>
              <input type="range" min="50" max="200" className="price-slider" />
              <button className="apply-filter">Apply Filter</button>
            </div>
          </aside>

          {/* Main Shop Section */}
          <main className="shop-main">
            <div className="shop-header">
              <h2>Best Seller</h2>
              <div className="shop-header-controls">
                <span>
                  Showing {indexOfFirstProduct + 1}-{Math.min(indexOfLastProduct, allProducts.length)} of{" "}
                  {allProducts.length} Products
                </span>
                <select>
                  <option>Most Popular</option>
                  <option>Price Low to High</option>
                  <option>Price High to Low</option>
                </select>
              </div>
            </div>

            {/* Product Grid */}
            <div className="product-grid">
              {currentProducts.map((product) => (
                <div key={product.id} className="product-card">
                  <div
                    className="product-image"
                    style={{ backgroundImage: `url(${ProductImage})` }}
                  ></div>
                  <h3 className="product-name">{product.name}</h3>
                  <div className="product-price">
                    <span className="current-price">${product.price}</span>
                  </div>
                  <Link to={`/product/${product.id}`} className="product-link">
                    <div className="view-details">View Details</div>
                  </Link>
                </div>
              ))}
            </div>

            {/* Pagination */}
            <div className="pagination">
              <button
                className="arrow-btn"
                onClick={handlePrevPage}
                disabled={currentPage === 1}
              >
                ←
              </button>
              {renderPagination()}
              <button
                className="arrow-btn"
                onClick={handleNextPage}
                disabled={currentPage === totalPages}
              >
                →
              </button>
            </div>
          </main>
        </div>
      </div>
      <Footnote />
    </>
  );
};

export default BestSeller;