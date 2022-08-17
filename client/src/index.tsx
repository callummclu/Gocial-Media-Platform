import { MantineProvider } from '@mantine/core';
import { NotificationsProvider } from '@mantine/notifications';
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { AuthProvider } from './hooks/useAuth';
import './styles/index.css'

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <AuthProvider>
      <MantineProvider>
        <NotificationsProvider>
          <App />
        </NotificationsProvider>
      </MantineProvider>
    </AuthProvider>
  </React.StrictMode>
);
