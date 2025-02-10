import React from 'react';
import './HeartsBackground.css';

const heartsCount = 20; // number of hearts to display

function HeartsBackground() {
  // Create an array with a length of heartsCount and map over it to create heart divs.
  return (
    <div className="hearts-background">
      {Array.from({ length: heartsCount }).map((_, index) => {
        // Randomize left position, animation duration, and delay.
        const left = Math.random() * 100; // percentage of viewport width
        const duration = 3 + Math.random() * 5; // seconds
        const delay = Math.random() * 2; // seconds

        return (
          <div
            key={index}
            className="heart"
            style={{
              left: `${left}vw`,
              animationDuration: `${duration}s`,
              animationDelay: `${delay}s`,
            }}
          ></div>
        );
      })}
    </div>
  );
}

export default HeartsBackground;
