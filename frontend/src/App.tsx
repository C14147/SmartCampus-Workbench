import React, { useState, useEffect } from 'react'
import { LoginPage } from './pages/LoginPage'
import { RegisterPage } from './pages/RegisterPage'
import { DashboardPage } from './pages/DashboardPage'
import { SchoolsPage } from './pages/SchoolsPage'
import { AssignmentsPage } from './pages/AssignmentsPage'
import { ApiClient } from './services/apiClient'

type View = 'login' | 'register' | 'dashboard' | 'schools' | 'assignments'

export const App: React.FC = () => {
  const [token, setToken] = useState<string | null>(localStorage.getItem('token'))
  const [view, setView] = useState<View>(token ? 'dashboard' : 'login')

  useEffect(() => {
    if (token) {
      localStorage.setItem('token', token)
      ApiClient.setToken(token)
    } else {
      localStorage.removeItem('token')
      ApiClient.setToken('')
    }
  }, [token])

  if (!token) {
    if (view === 'register') return <RegisterPage onRegistered={() => setView('login')} onBack={() => setView('login')} />
    return <LoginPage onLogin={(t) => { setToken(t); setView('dashboard') }} onRegister={() => setView('register')} />
  }

  return (
    <div style={{padding:20}}>
      <header style={{display:'flex', gap:12, marginBottom:16}}>
        <button onClick={() => setView('dashboard')}>Dashboard</button>
        <button onClick={() => setView('schools')}>Schools</button>
        <button onClick={() => setView('assignments')}>Assignments</button>
        <button onClick={() => { setToken(null); setView('login') }}>Sign out</button>
      </header>

      <main>
        {view === 'dashboard' && <DashboardPage />}
        {view === 'schools' && <SchoolsPage />}
        {view === 'assignments' && <AssignmentsPage />}
      </main>
    </div>
  )
}

