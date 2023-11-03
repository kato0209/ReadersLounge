import './App.css'
import { AppRoutes } from './routes'
import { BrowserRouter } from 'react-router-dom'
import { ErrorBoundary } from 'react-error-boundary'
import ErrorFallback from './components/Error/ErrorFallback'
import { AuthProvider } from './lib/auth/auth'

function App() {
  return (
    <ErrorBoundary FallbackComponent={ErrorFallback}>
      <AuthProvider>
        <BrowserRouter>
          <AppRoutes />
        </BrowserRouter>
      </AuthProvider>
    </ErrorBoundary>
  )
}

export default App
