import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import headerImage from './assets/StareDown.png'; 
import './App.css';

function Home() {
  const [noButtonPosition, setNoButtonPosition] = useState({
    left: 'calc(50% - 50px)',
    top: '200px'
  });
  const navigate = useNavigate();

  const handleYesClick = () => {
    navigate('/suggestions');
  };

  const handleNoMouseOver = () => {
    const buttonWidth = 100;  // approximate width of the "No" button
    const buttonHeight = 50;   // approximate height of the "No" button

    // Define the region that is 80% of the viewport dimensions.
    const regionWidth = window.innerWidth * 0.8;
    const regionHeight = window.innerHeight * 0.8;

    // Calculate the offset so that the 80% region is centered in the viewport.
    const leftOffset = (window.innerWidth - regionWidth) / 2;
    const topOffset = (window.innerHeight - regionHeight) / 2;

    // Generate random positions within the 80% region, ensuring the button doesn't overflow.
    const randomLeftWithinRegion = Math.floor(Math.random() * (regionWidth - buttonWidth));
    const randomTopWithinRegion = Math.floor(Math.random() * (regionHeight - buttonHeight));

    // Add the offset so the region is centered.
    const randomLeft = leftOffset + randomLeftWithinRegion;
    const randomTop = topOffset + randomTopWithinRegion;

    setNoButtonPosition({ left: `${randomLeft}px`, top: `${randomTop}px` });
  };

  return (
    <div className="full-screen-container">
      {/* Use the imported image */}
      <img src={headerImage} alt="Header" className="header-image" />
      
      <h1>Will you be my Valentine?</h1>
      <div className="button-container" style={{ position: 'relative', width: '100%', height: '100%' }}>
        <button className="yes-button" onClick={handleYesClick}>
          Yes
        </button>
        <button
          className="no-button"
          onMouseOver={handleNoMouseOver}
          style={{
            position: 'absolute',
            left: noButtonPosition.left,
            top: noButtonPosition.top,
          }}
          onClick={(e) => e.preventDefault()}
        >
          No
        </button>
      </div>
    </div>
  );
}

export default Home;
