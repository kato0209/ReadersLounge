import './App.css'
import { AppRoutes } from './routes'
import { BrowserRouter } from 'react-router-dom'
import { ErrorBoundary } from 'react-error-boundary'
import ErrorFallback from './components/Error/ErrorFallback'
import { AuthProvider } from './lib/auth/auth'
import { CookiesProvider } from 'react-cookie';

function App() {
  return (
    <ErrorBoundary FallbackComponent={ErrorFallback}>
      <CookiesProvider>
        <AuthProvider>
          <BrowserRouter>
            <AppRoutes />
          </BrowserRouter>
        </AuthProvider>
      </CookiesProvider>
    </ErrorBoundary>
  )
}

export default App
