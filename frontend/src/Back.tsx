import React from 'react';
import root from './index';

function Back(Component : JSX.Element) {
    root.render(
        <React.StrictMode>
          {Component}
        </React.StrictMode>
      );
}

export default Back;
