import React, { Suspense } from 'react';
import ReactDOM from 'react-dom';
import './i18n';
import './index.css';
import App from './App';

ReactDOM.render(
  <Suspense fallback="loading">
    <App />
  </Suspense>,
  document.getElementById('root')
);
