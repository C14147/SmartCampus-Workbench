import React, { useState } from 'react'
import { ApiClient } from '../services/apiClient'

export const RegisterPage: React.FC<{ onRegistered: () => void; onBack: () => void }> = ({ onRegistered, onBack }) => {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState<string | null>(null)

  const submit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await ApiClient.post('/auth/register', { username, email, password })
      onRegistered()
    } catch (err: any) {
      setError(err?.response?.data?.error || err.message)
    }
  }

  return (
    <div style={{maxWidth:400, margin:'60px auto'}}>
      <h2>Register</h2>
      <form onSubmit={submit}>
        <div>
          <label>Username</label>
          <input value={username} onChange={e => setUsername(e.target.value)} />
        </div>
        <div>
          <label>Email</label>
          <input value={email} onChange={e => setEmail(e.target.value)} />
        </div>
        <div>
          <label>Password</label>
          <input type="password" value={password} onChange={e => setPassword(e.target.value)} />
        </div>
        <div style={{marginTop:12}}>
          <button type="submit">Register</button>
          <button type="button" onClick={onBack} style={{marginLeft:8}}>Back</button>
        </div>
        {error && <div style={{color:'red', marginTop:8}}>{error}</div>}
      </form>
    </div>
  )
}
