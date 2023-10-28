import './App.css'
import { AppRoutes } from './routes'
import { BrowserRouter } from 'react-router-dom'
import { ErrorBoundary } from 'react-error-boundary'
import ErrorFallback from './components/Error/ErrorFallback'
import { CookiesProvider } from 'react-cookie';

function App() {
  return (
    <ErrorBoundary FallbackComponent={ErrorFallback}>
      <CookiesProvider>
        <BrowserRouter>
          <AppRoutes />
        </BrowserRouter>
      </CookiesProvider>
    </ErrorBoundary>
  )
}

export default App
