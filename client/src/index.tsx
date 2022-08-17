import { MantineProvider } from '@mantine/core';
import { NotificationsProvider } from '@mantine/notifications';
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { AuthProvider } from './hooks/useAuth';
import { PostProvider } from './hooks/usePost';
import './styles/index.css'

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <AuthProvider>
      <PostProvider>
      <MantineProvider>
        <NotificationsProvider>
          <App />
        </NotificationsProvider>
      </MantineProvider>
      </PostProvider>
    </AuthProvider>
  </React.StrictMode>
);
