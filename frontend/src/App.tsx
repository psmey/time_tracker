import { CssVarsProvider, CssBaseline } from '@mui/joy'
import { Router } from './router'

function App() {
  return (
    <CssVarsProvider>
      <CssBaseline />
      <Router></Router>
    </CssVarsProvider>
  )
}

export default App
