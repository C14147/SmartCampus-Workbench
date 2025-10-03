import React, { useState } from 'react'
import axios from 'axios'

export const LoginPage: React.FC<{ onLogin: (token: string) => void }> = ({ onLogin }) => {
  const [username, setUsername] = useState('admin')
  const [password, setPassword] = useState('password')
  const [error, setError] = useState<string | null>(null)

  const submit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      const res = await axios.post('/api/v1/auth/login', { username, password })
      onLogin(res.data.token)
    } catch (err: any) {
      setError(err?.response?.data?.error || err.message)
    }
  }

  return (
    <div style={{maxWidth:400, margin:'60px auto'}}>
      <h2>SmartCampus Login (demo)</h2>
      <form onSubmit={submit}>
        <div>
          <label>Username</label>
          <input value={username} onChange={e => setUsername(e.target.value)} />
        </div>
        <div style={{marginTop:8}}>
          <label>Password</label>
          <input type="password" value={password} onChange={e => setPassword(e.target.value)} />
        </div>
        <div style={{marginTop:12}}>
          <button type="submit">Sign in</button>
        </div>
        {error && <div style={{color:'red', marginTop:8}}>{error}</div>}
      </form>
    </div>
  )
}
